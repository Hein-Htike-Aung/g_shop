/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models

type SysTables struct {
	TBName          string       `gorm:"column:table_name" json:"tableName"` // table name
	MLTBName        string       `gorm:"-" json:"-"`
	TableComment    string       `json:"tableComment"`    // table comment
	ClassName       string       `json:"className"`       // class name
	TplCategory     string       `json:"tplCategory"`     //
	PackageName     string       `json:"packageName"`     // package name
	ModuleName      string       `json:"moduleName"`      // Go module file name
	ModuleFrontName string       `json:"moduleFrontName"` // frontend file name
	BusinessName    string       `json:"businessName"`    //
	FunctionName    string       `json:"functionName"`
	FunctionAuthor  string       `json:"functionAuthor"`
	Crud            bool         `json:"crud"`
	Remark          string       `json:"remark"`
	IsDataScope     int          `json:"isDataScope"`
	IsActions       int          `json:"isActions"`
	IsAuth          int          `json:"isAuth"`
	Columns         []SysColumns `gorm:"-" json:"columns"`
	BaseModel
}

func GetAllTables(pageNUm int, pageSize int, maps interface{}) (int64, []SysTables) {
	var (
		total int64
		lists []SysTables
	)
	db.Model(&SysTables{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists

	return total, lists
}

func UpdateByTables(m *SysTables) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}
