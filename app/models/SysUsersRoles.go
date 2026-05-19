/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models

type SysUsersRoles struct {
	Id int64
	UserId int64 `gorm:"column:sys_user_id;"`
	RoleId int64 `gorm:"column:sys_role_id;"`
}

