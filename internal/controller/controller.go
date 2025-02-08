package controller

import (
	"awesomeProject/Project/WMS/internal/service"
	"github.com/gin-gonic/gin"
	oresponse "github.com/omniful/go_commons/response"
)

type Controller struct {
	Hello string
}

func EchoController(ctx *gin.Context) {
	var empty Controller
	empty.Hello = "Hello World"
	internal.Test_service()
	oresponse.NewSuccessResponse(ctx, empty)
}
