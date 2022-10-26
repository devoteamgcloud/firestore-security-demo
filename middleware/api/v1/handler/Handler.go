package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler[QueryString any] struct {
	path    string
	methods []string
	view    func(qs *QueryString) (int, interface{})
}

func (handler Handler[QueryString]) Init(methods []string, path string, view func(qs *QueryString) (int, interface{})) *Handler[QueryString] {
	handler.path = path
	handler.methods = methods
	handler.view = view
	return &handler
}

func (handler *Handler[QueryString]) Handle(ctx *gin.Context) {
	withRequest(ctx, func(queryString *QueryString) {
		status, response := handler.view(queryString)
		ctx.JSON(status, response)
	})
}

func (handler *Handler[QueryString]) Register(router gin.IRouter) {
	if handler.methods == nil {
		panic("handler was not initiated")
	}
	for _, method := range handler.methods {
		router.Handle(method, handler.path, handler.Handle)
	}
}

func (handler *Handler[QueryString]) String() string {
	return fmt.Sprintf("Handler Handler for %s", handler.path)
}

func withRequest[QueryString any](ctx *gin.Context, call func(*QueryString)) {
	var qs QueryString
	if isNil(qs) {
		call(nil)
	} else if err := ctx.ShouldBindQuery(&qs); err != nil {
		UnprocessableEntityResponse(err, ctx)
	} else {
		call(&qs)
	}
}

func isNil[QueryString any](qs QueryString) bool {
	var test interface{} = qs
	_, ok := test.(Nil)
	return ok
}
