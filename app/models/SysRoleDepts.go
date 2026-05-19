/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models

type SysRolesDepts struct {
	Id int64
	RoleId *SysRole `gorm:"column:sys_role_id;association_autoupdate:false;association_autocreate:false"`
	DeptId *SysDept `gorm:"column:sys_dept_id;association_autoupdate:false;association_autocreate:false"`
}




