package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

// 如果给定一系列task，要求在规定的timeout里跑完，不然就报错
// 如果操作系统给了中断信号也报错
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time // 用来计时
	tasks     []func(int)      // task的列表
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t), // 接受一个Duration
		tasks:     make([]func(int), 0),
	}
}
func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	// run the task
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}

}
