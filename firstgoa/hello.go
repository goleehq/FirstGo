package main

import (
	"github.com/phxuecan/FirstGo/firstgoa/app"
	"github.com/goadesign/goa"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Answer runs the answer action.
func (c *HelloController) Answer(ctx *app.AnswerHelloContext) error {
	// HelloController_Answer: start_implement

	// Put your logic here
	answer := app.Message{
		User:ctx.Payload.User,
		Info:ctx.Payload.Info,
	}
	// HelloController_Answer: end_implement
	return ctx.OK([]byte(*answer.Info.Content))
}

// Say runs the say action.
func (c *HelloController) Say(ctx *app.SayHelloContext) error {
	// HelloController_Say: start_implement

	// Put your logic here

	// HelloController_Say: end_implement
	return ctx.OK([]byte(ctx.Words))
}
