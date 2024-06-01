package routenames

// Routes live in their own package so that they can be imported
// into any other package without import cycles. They can thus
// be imported into templates seamlessy.
const (
	RouteNameForgotPassword       = "forgot_password"
	RouteNameForgotPasswordSubmit = "forgot_password.submit"
	RouteNameLogin                = "login"
	RouteNameLoginSubmit          = "login.submit"
	RouteNameLogout               = "logout"
	RouteNameRegister             = "register"
	RouteNameRegisterSubmit       = "register.submit"
	RouteNameResetPassword        = "reset_password"
	RouteNameResetPasswordSubmit  = "reset_password.submit"
	RouteNameVerifyEmail          = "verify_email"
	RouteNameContact              = "contact"
	RouteNameContactSubmit        = "contact.submit"
	RouteNameAbout                = "about"
	RouteNameHome                 = "home"
	RouteNameSearch               = "search"
	RouteNameDashboard            = "dashboard"
	RouteNameRealtime             = "realtime"
)
