package taskrunner

type TaskData struct {
	In    <-chan interface{}
	Out   chan<- interface{}
	Error chan<- error
}
type TaskRunner interface {
	Run(data *TaskData)
}

type TaskRunnerFunc func(data *TaskData)

func (runner TaskRunnerFunc) Run(data *TaskData) {
	runner(data)
}

func RunTask(data *TaskData, runner TaskRunner) {
	go func() {
		defer close(data.Out)
		runner.Run(data)
	}()
}

func ChainTasks(data *TaskData, runners ...TaskRunner) {
	for i, runner := range runners {
		if i < len(runners)-1 {
			ch := make(chan interface{})
			RunTask(&TaskData{In: data.In, Out: ch, Error: data.Error}, runner)
			data.In = ch
		} else {
			RunTask(&TaskData{In: data.In, Out: data.Out, Error: data.Error}, runner)
		}
	}
}
