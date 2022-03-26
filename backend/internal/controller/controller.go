package controller

import "exam.com/webchecker/internal/core"

type Controller struct {
	WebCheckerController
}

func NewController(core *core.Core) *Controller {
	return &Controller{
		NewWebCheckerController(core),
	}
}
