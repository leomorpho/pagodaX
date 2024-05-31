package types

import "github.com/mikestefanello/pagoda/pkg/controller"

type (
	RegisterForm struct {
		Name       string `form:"name" validate:"required"`
		Email      string `form:"email" validate:"required,email"`
		Password   string `form:"password" validate:"required"`
		Submission controller.FormSubmission
	}
)
