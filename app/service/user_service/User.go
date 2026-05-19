/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package user_service

import (
	"errors"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/models/vo"
	userDto "yixiang.co/go-mall/app/service/user_service/dto"
	"yixiang.co/go-mall/pkg/util"
)

type User struct {
	Id int64
	Username string

	DeptId int64
	Enabled int

	PageNum int
	PageSize int

	M *models.SysUser

	Ids []int64

	UserPost userDto.UserPost
	UserPass userDto.UserPass

	ImageUrl string

}

func (u *User) UpdateImage() error {
	user ,err := models.GetUserById(u.Id)
	if err != nil {
		return  errors.New("user does not exist")
	}

	user.Avatar = u.ImageUrl
	return models.UpdateCurrentUser(&user)
}

func (u *User) UpdatePass() error {
	user ,err := models.GetUserById(u.Id)
	if err != nil {
		return  errors.New("user does not exist")
	}
	if !util.ComparePwd(user.Password,[]byte(u.UserPass.OldPass)) {
		return  errors.New("incorrect old password")
	}
	user.Password = util.HashAndSalt([]byte(u.UserPass.NewPass))
	return models.UpdateCurrentUser(&user)
}


func (u *User) UpdateProfile() error {
	user ,err := models.GetUserById(u.Id)
	if err != nil {
		return  err
	}

	user.Phone = u.UserPost.Phone
	user.Sex = u.UserPost.Sex
	user.NickName = u.UserPost.NickName

	return models.UpdateCurrentUser(&user)

}

func (u *User) GetUserOneByName() (*models.SysUser, error) {
	return models.GetUserByUsername(u.Username)
}

func (u *User) GetUserAll() vo.ResultList {
	maps := make(map[string]interface{})
	if u.DeptId >= 0 {
		maps["dept_id"] = u.DeptId
	}
	if u.Enabled >= 0 {
		maps["enabled"] = u.Enabled
	}
	if u.Username != "" {
		maps["username"] = u.Username
	}

	total,list := models.GetAllUser(u.PageNum,u.PageSize,maps)
	return vo.ResultList{Content: list,TotalElements: total}
}

func (u *User) Insert() error {
	return models.AddUser(u.M)
}

func (u *User) Save() error {
	return models.UpdateByUser(u.M)
}

func (u *User) Del() error {
	return models.DelByUser(u.Ids)
}

