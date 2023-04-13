package handler

import (
	"demo04/internal/service"
	"demo04/pkg/api"
	"demo04/pkg/res"
	"github.com/gin-gonic/gin"
	"log"
)

func AddFriend(ginCtx *gin.Context) {
	var req api.FriendActionRequest
	var friendService service.FriendService
	PanicIfError(ginCtx.ShouldBind(&req))
	log.Println(req)
	resp, err := friendService.AddFriend(&req)
	res.SendAutoResponse(ginCtx, err, resp)
}

func DeleteFriend(ginCtx *gin.Context) {
	var req api.FriendActionRequest
	var friendService service.FriendService
	PanicIfError(ginCtx.ShouldBind(&req))
	log.Println(req)
	resp, err := friendService.DeleteFriend(&req)
	res.SendAutoResponse(ginCtx, err, resp)
}
