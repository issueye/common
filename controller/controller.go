package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/common/model"
)

type Controller[T any] struct {
	*gin.Context
}

func New[T any](ctx *gin.Context) *Controller[T] {
	return &Controller[T]{
		Context: ctx,
	}
}

func NewA(ctx *gin.Context) *Controller[any] {
	return New[any](ctx)
}

func (c *Controller[T]) Success() {
	Success(c.Context)
}

func (c *Controller[T]) SuccessByMsg(msg string) {
	SuccessByMsg(c.Context, msg)
}

func (c *Controller[T]) SuccessByMsgf(fmtStr string, args ...any) {
	SuccessByMsgf(c.Context, fmtStr, args...)
}

func (c *Controller[T]) SuccessData(data interface{}) {
	SuccessData(c.Context, data)
}

func (c *Controller[T]) SuccessPage(req *model.Page[T], data interface{}) {
	SuccessPage(c.Context, req, data)
}

func (c *Controller[T]) Fail() {
	Fail(c.Context, BAD_REQUEST)
}

func (c *Controller[T]) FailByMsg(msg string) {
	FailByMsg(c.Context, msg)
}

func (c *Controller[T]) FailByMsgf(fmtStr string, args ...any) {
	FailByMsgf(c.Context, fmtStr, args...)
}

func (c *Controller[T]) FailBind(err error) {
	FailBind(c.Context, err)
}
