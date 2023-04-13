package repository

import "gorm.io/gorm"

type Friend struct {
	UserId   string
	FriendId string
	Model
}

func (friend *Friend) CheckIsFriend() bool {
	err1 := DB.Where("user_id = ?", friend.UserId).Where("friend_id = ?", friend.FriendId).First(&friend).Error
	err2 := DB.Where("user_id = ?", friend.FriendId).Where("friend_id = ?", friend.UserId).First(&friend).Error
	if err1 == gorm.ErrRecordNotFound && err2 == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (friend *Friend) AddFriend() error {
	friend.Model.Id = SF.NextVal()
	err := DB.Create(&friend).Error
	if err != nil {
		return err
	}
	friend.Model.Id = SF.NextVal()
	temp := friend.UserId
	friend.FriendId = temp
	friend.UserId = friend.FriendId
	err = DB.Create(&friend).Error
	return err
}

func (friend *Friend) DeleteFriend() error {
	err := DB.Delete(&friend).Error
	if err != nil {
		return err
	}
	temp := friend.UserId
	friend.UserId = friend.FriendId
	friend.FriendId = temp
	err = DB.Delete(&friend).Error
	return err
}
