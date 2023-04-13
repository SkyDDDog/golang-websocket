package handler

import (
	"demo04/internal/service"
	"demo04/pkg/api"
	"demo04/pkg/res"
	"github.com/gin-gonic/gin"
	"log"
)

func UserLogin(ginCtx *gin.Context) {
	var req api.UserAuthRequest
	var userService service.UserService
	PanicIfError(ginCtx.ShouldBind(&req))
	log.Println(req.Username)
	log.Println(req.Password)
	resp, err := userService.UserLogin(&req)
	res.SendAutoResponse(ginCtx, err, resp)
}

func UserRegister(ginCtx *gin.Context) {
	var req api.UserAuthRequest
	var userService service.UserService
	PanicIfError(ginCtx.ShouldBind(&req))
	resp, err := userService.UserRegister(&req)
	res.SendAutoResponse(ginCtx, err, resp)
}
