package core

import "exam.com/webchecker/internal/gateway"

type Core struct {
	WebCheckerCore
}

func NewCore(g *gateway.Gateway) *Core {
	return &Core{
		NewWebCheckerCore(g),
	}
}
