package main

import (
	"fmt"

	"github.com/gofish2020/common/compress"
	"github.com/gofish2020/common/cronjob"
)

//var localCache = cache.NewLocalCache(context.Background(), nil)

func main() {
	//zap.AddCaller()
	//zap.AddCallerSkip(1)
	//zap.AddStacktrace(zapcore.InfoLevel)

	//fmt.Printf("%s", debug.Stack())
	//CompressData()
	//select {}
	//RedisTest()
	awesomeCache := single.NewAwesomeCache{}
	// for i := 0; i < 500; i++ {
	// 	go func() {
	// 		GroupGo()
	// 	}()
	// }
	select {}
}

// var (
// 	cacheKey = "key:cache:user:info"
// )

// var gp singleflight.Group

// func GroupGo() string {
// 	var result *interface{} = new(interface{})

// 	if localCache.Get(cacheKey, &result) == cache.Nil {

// 		result, err, _ := gp.Do(cacheKey, func() (interface{}, error) {
// 			fmt.Println("get data from db")

// 			localCache.Set(cacheKey, "happy mother day")

// 			return "happy mother day", nil
// 		})

// 		if err != nil {
// 			return ""
// 		}

// 		return result.(string)
// 	}
// 	fmt.Println(" get data from cache")
// 	return (*result).(string)
// }

func RedisTest() {
	// fmt.Println(gredis.NewClient().Set(context.Background(), "nash", "中国", time.Minute))

	// fmt.Println(gredis.NewClient().Get(context.Background(), "nash"))

	// fmt.Println(gredis.NewClient().MSet(context.Background(), "lala", "22", "bb", "111"))
	// fmt.Println(gredis.NewClient().MGet(context.Background(), "lala", "bb"))

	// go func() {

	// 	//gredis.NewClient().Subscripe(context.Background(), "xiaoyu", "xiaoyu1").Receive(context.Background())
	// 	for val := range gredis.NewClient().Subscripe(context.Background(), "xiaoyu", "xiaoyu1").Channel() {
	// 		fmt.Println(val)
	// 	}

	// 	fmt.Println("end")
	// }()

	// time.Sleep(time.Second)
	// gredis.NewClient().Publish(context.Background(), "xiaoyu", "love")
	// time.Sleep(time.Second)
	// gredis.NewClient().Publish(context.Background(), "xiaoyu1", "love1")

	//select {}
}

func CompressData() {
	z := compress.NewCompressZlib()
	result, _ := z.Marshal([]byte("5545kjdsj中国"))

	fmt.Println(len(result))
	data, _ := z.Unmarshal(result)
	fmt.Println(len(data))
	fmt.Println(string(data))

}

type MyStruct struct {
	Val int `json:"val"`
}

func Tsss() {
	// _, file, line, ok := runtime.Caller(1)
	// fmt.Println(file)
	// fmt.Println(line)
	// fmt.Println(ok)

	var obj MyStruct
	obj.Val = 0
	go func() {
		c := cronjob.NewCronjob()
		tasks := []cronjob.CronTask{}
		tasks = append(tasks, cronjob.NewCronTask("@every 2s", func() {
			obj.Val++
			localCache.Set("love", obj)
		}))
		c.AddFunc(tasks...)
		c.Start()
	}()

	go func() {

		c := cronjob.NewCronjob()
		tasks := []cronjob.CronTask{}
		tasks = append(tasks, cronjob.NewCronTask("@every 4s", func() {
			res := &MyStruct{}
			localCache.Get("love", res)
			fmt.Printf("res:%+v", res)
		}))
		c.AddFunc(tasks...)
		c.Start()

	}()
}
