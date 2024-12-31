package util

import (
	"context"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/pkg/constant"
)

type Context struct {
	mu     sync.Mutex
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewContext() *Context {
	ctxInstance := &Context{
		Ctx: context.Background(),
	}

	return ctxInstance
}

func (c *Context) WithTimeout(dur int) *Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctx, cancel := context.WithTimeout(c.Ctx, time.Duration(dur)*time.Second)
	c.Ctx = ctx
	c.Cancel = cancel
	return c
}

func (c *Context) WithClaims(fc *fiber.Ctx) *Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctx := context.WithValue(c.Ctx, constant.ClaimsContextKey, fc.Locals(constant.ClaimsContextKey))
	c.Ctx = ctx

	return c
}
