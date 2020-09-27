package commonds

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	testTask := toolbox.NewTask("test", "0 12 * * * *", Test)
	toolbox.AddTask(testTask.Taskname, testTask)
}

func Test() error {
	fmt.Println("task crontab")
	return nil
}
