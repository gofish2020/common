package cronjob

import (
	"github.com/robfig/cron/v3"
)

type CronFunc func()

type CronTask struct {
	Spec string
	f    CronFunc
}

// NewCronTask 创建定时任务
func NewCronTask(spec string, f CronFunc) CronTask {
	return CronTask{
		Spec: spec,
		f:    f,
	}
}

// NewCronjob 创建cron任务
func NewCronjob() *CronJob {
	return &CronJob{
		c: cron.New(),
	}

}

type CronJob struct {
	c *cron.Cron
}

func (t *CronJob) AddFunc(tasks ...CronTask) {
	for _, task := range tasks {
		t.c.AddFunc(task.Spec, task.f)
	}
}

func (t *CronJob) Start() {
	t.c.Start()
}
