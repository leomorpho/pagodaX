package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"goweb/config"
	"goweb/ent"
	"goweb/ent/passwordtoken"
	"goweb/ent/user"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	// sessionName stores the name of the session which contains authentication data
	sessionName = "ua"

	// sessionKeyUserID stores the key used to store the user ID in the session
	sessionKeyUserID = "user_id"

	// sessionKeyAuthenticated stores the key used to store the authentication status in the session
	sessionKeyAuthenticated = "authenticated"
)

// Client is the client that handles authentication requests
type Client struct {
	config *config.Config
	orm    *ent.Client
}

// NewClient creates a new authentication client
func NewClient(cfg *config.Config, orm *ent.Client) *Client {
	return &Client{
		config: cfg,
		orm:    orm,
	}
}

// Login logs in a user of a given ID
func (c *Client) Login(ctx echo.Context, userID int) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	sess.Values[sessionKeyUserID] = userID
	sess.Values[sessionKeyAuthenticated] = true
	return sess.Save(ctx.Request(), ctx.Response())
}

// Logout logs the requesting user out
func (c *Client) Logout(ctx echo.Context) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	sess.Values[sessionKeyAuthenticated] = false
	return sess.Save(ctx.Request(), ctx.Response())
}

// GetAuthenticatedUserID returns the authenticated user's ID, if the user is logged in
func (c *Client) GetAuthenticatedUserID(ctx echo.Context) (int, error) {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return 0, err
	}

	if sess.Values[sessionKeyAuthenticated] == true {
		return sess.Values[sessionKeyUserID].(int), nil
	}

	return 0, NotAuthenticatedError{}
}

// GetAuthenticatedUser returns the authenticated user if the user is logged in
func (c *Client) GetAuthenticatedUser(ctx echo.Context) (*ent.User, error) {
	if userID, err := c.GetAuthenticatedUserID(ctx); err == nil {
		return c.orm.User.Query().
			Where(user.ID(userID)).
			Only(ctx.Request().Context())
	}

	return nil, NotAuthenticatedError{}
}

// HashPassword returns a hash of a given password
func (c *Client) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword check if a given password matches a given hash
func (c *Client) CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GeneratePasswordResetToken generates a password reset token for a given user.
// For security purposes, the token itself is not stored in the database but rather
// a hash of the token, exactly how passwords are handled. This method returns both
// the generated token as well as the token entity which only contains the hash.
func (c *Client) GeneratePasswordResetToken(ctx echo.Context, userID int) (string, *ent.PasswordToken, error) {
	// Generate the token, which is what will go in the URL, but not the database
	token, err := c.RandomToken(c.config.App.PasswordToken.Length)
	if err != nil {
		return "", nil, err
	}

	// Hash the token, which is what will be stored in the database
	hash, err := c.HashPassword(token)
	if err != nil {
		return "", nil, err
	}

	// Create and save the password reset token
	pt, err := c.orm.PasswordToken.
		Create().
		SetHash(hash).
		SetUserID(userID).
		Save(ctx.Request().Context())

	return token, pt, err
}

// GetValidPasswordToken returns a valid password token entity for a given user and a given token.
// Since the actual token is not stored in the database for security purposes, all non-expired token entities
// are fetched from the database belonging to the requesting user and a hash of the provided token is compared
// with the hash stored in the database.
func (c *Client) GetValidPasswordToken(ctx echo.Context, token string, userID int) (*ent.PasswordToken, error) {
	// Ensure expired tokens are never returned
	expiration := time.Now().Add(-c.config.App.PasswordToken.Expiration)

	// Query to find all tokens for te user that haven't expired
	// We need to get all of them in order to properly match the token to the hashes
	pts, err := c.orm.PasswordToken.
		Query().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Where(passwordtoken.CreatedAtGTE(expiration)).
		All(ctx.Request().Context())

	if err != nil {
		ctx.Logger().Error(err)
		return nil, err
	}

	// Check all tokens for a hash match
	for _, pt := range pts {
		if err := c.CheckPassword(token, pt.Hash); err == nil {
			return pt, nil
		}
	}

	return nil, InvalidTokenError{}
}

// DeletePasswordTokens deletes all password tokens in the database for a belonging to a given user.
// This should be called after a successful password reset.
func (c *Client) DeletePasswordTokens(ctx echo.Context, userID int) error {
	_, err := c.orm.PasswordToken.
		Delete().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Exec(ctx.Request().Context())

	return err
}

// RandomToken generates a random token string of a given length
func (c *Client) RandomToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
