package workers

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TransactionID string
}

type Worker struct {
	ID        int
	TaskQueue chan Task
	DB        *gorm.DB
	WaitGroup *sync.WaitGroup
}

func NewWorker(id int, taskQueue chan Task, db *gorm.DB, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:        id,
		TaskQueue: taskQueue,
		DB:        db,
		WaitGroup: wg,
	}
}

func (w *Worker) Start() {
	go func() {
		defer w.WaitGroup.Done()
		for task := range w.TaskQueue {
			w.ProcessTask(&task)
		}
	}()
}

func (w *Worker) ProcessTask(task *Task) {
	time.Sleep(2 * time.Second) // Replace with actual check
	fmt.Printf("Worker %d processed transaction %s\n", w.ID, task.TransactionID)
}
