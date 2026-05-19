/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package cron_job_service

import (
	"errors"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/models/vo"
	"yixiang.co/go-mall/app/service/cron_job_service/task"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/util/cron"
)

type SysCronJob struct {
	Id      int64
	Name    string
	Enabled int

	PageNum  int
	PageSize int

	M *models.SysCronJob

	Ids []int64
}

func (d *SysCronJob) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Enabled >= 0 {
		maps["enabled"] = d.Enabled
	}
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllSysCronJob(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *SysCronJob) Insert() error {
	return models.AddSysCronJob(d.M)
}

func (d *SysCronJob) Save() error {
	return models.UpdateBySysCronJob(d.M)
}

func (d *SysCronJob) Del() error {
	return models.DelBySysCronJob(d.Ids)
}

func (d *SysCronJob) Exec() error {
	var job models.SysCronJob
	res := global.YSHOP_DB.Where("id = ?", d.Id).First(&job).RowsAffected
	err := global.YSHOP_DB.Model(&models.SysCronJob{}).Where("id = ?", d.Id).Update("status", 0).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
	}

	if res == 0 {
		return errors.New("data not found")
	}
	if !task.IsExistFunc(job.InvokeTarget) {
		return errors.New("add the target function under fun")
	}
	f := task.GetByName(job.InvokeTarget)

	// start job
	err = cron.Start(f, job.Id, job.CronExpression)
	return err

}

func (d *SysCronJob) Stop() error {
	var job models.SysCronJob
	res := global.YSHOP_DB.Where("id = ?", d.Id).First(&job).RowsAffected
	global.YSHOP_DB.Model(&models.SysCronJob{}).Where("id = ?", d.Id).Update("status", 1)
	global.YSHOP_LOG.Error(res)
	if res == 0 {
		return errors.New("data not found")
	}

	if cron.IsExistCron(d.Id) {
		cron.Stop(d.Id)
	}

	return nil

}
