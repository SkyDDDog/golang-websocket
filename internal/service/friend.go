package service

import (
	"demo04/internal/repository"
	"demo04/pkg/api"
	"demo04/pkg/errno"
	"log"
)

type FriendService struct {
}

func NewFriendService() *FriendService {
	return &FriendService{}
}

func (*FriendService) AddFriend(req *api.FriendActionRequest) (resp *api.FriendResponse, err error) {
	resp = new(api.FriendResponse)
	friend := &repository.Friend{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	}
	if friend.CheckIsFriend() {
		resp.ErrNo = errno.ErrorAlreadyFriendError
		return resp, err
	}
	err = friend.AddFriend()
	if err != nil {
		resp.ErrNo = errno.UnexpectedTypeError
		return resp, err
	}
	log.Println(resp)
	resp.ErrNo = errno.Success
	return resp, err
}

func (*FriendService) DeleteFriend(req *api.FriendActionRequest) (resp *api.FriendResponse, err error) {
	resp = new(api.FriendResponse)
	friend := &repository.Friend{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	}
	if !friend.CheckIsFriend() {
		resp.ErrNo = errno.ErrorNotFriendError
		return resp, err
	}
	err = friend.DeleteFriend()
	if err != nil {
		resp.ErrNo = errno.UnexpectedTypeError
		return resp, err
	}
	resp.ErrNo = errno.Success
	return resp, err

}
