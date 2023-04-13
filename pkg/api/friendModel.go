package api

import "demo04/pkg/errno"

type FriendActionRequest struct {
	UserId   string `json:"userId" form:"userId"`
	FriendId string `json:"friendId" form:"friendId"`
}

type FriendResponse struct {
	errno.ErrNo
}
