package job

import (
	"context"
	"log"
	"sync"
)

type JobManager struct {
	mu   sync.Mutex
	jobs map[string]context.CancelFunc
}

var managerInstance *JobManager
var managerOnce sync.Once

func GetJobManager() *JobManager {
	managerOnce.Do(func() {
		managerInstance = &JobManager{
			jobs: make(map[string]context.CancelFunc),
		}
	})
	return managerInstance
}

func (m *JobManager) Start(name string, job func(ctx context.Context) error, onError func(error)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if cancel, ok := m.jobs[name]; ok {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	m.jobs[name] = cancel

	go func() {
		err := job(ctx)
		if err != nil && onError != nil {
			onError(err)
		}
		m.mu.Lock()
		delete(m.jobs, name)
		m.mu.Unlock()
	}()
}

func (m *JobManager) IsRunning(name string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.jobs[name]
	return exists
}

func (m *JobManager) Cancel(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if cancel, ok := m.jobs[name]; ok {
		cancel()
		delete(m.jobs, name)
		log.Println("Cancelled job:", name)
	}
}

func (m *JobManager) CancelAll() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for name, cancel := range m.jobs {
		cancel()
		log.Println("Cancelled job:", name)
	}
	m.jobs = make(map[string]context.CancelFunc)
}
