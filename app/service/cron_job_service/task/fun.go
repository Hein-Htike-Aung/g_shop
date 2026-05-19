package task

import (
	"encoding/json"
	"fmt"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/redis"
)

// add custom cron tasks here
func init() {
	AddTask("TestCronFun", TestCronFun)
	AddTask("TestCronFun2", TestCronFun2)

}

// test with no parameters
func TestCronFun() {
	fmt.Println("test with no parameters")
}

type taskDemo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// test with parameters
func TestCronFun2() {
	cache := GetCacheParam("TestCronFun2")
	if cache == "" {
		fmt.Println("parameters not provided")
	} else {
		t := new(taskDemo)
		fmt.Println(cache)
		err := json.Unmarshal([]byte(cache), t)
		if err != nil {
			fmt.Println("invalid parameters")
		}
		fmt.Print("test with parameters: ")
		fmt.Println(t)
	}

}

// cache
func GetCacheParam(str string) string {
	var job models.SysCronJob
	key := constant.CRON_KEY + str
	res := redis.GetString(key)
	if res == "" {
		global.YSHOP_DB.Where("invoke_target= ?", str).First(&job)
		redis.SetString(key, job.JobParams, 0)

		res = job.JobParams
	}

	return res

}
