// MemStore - Volatile in-memory storage.
// Simple storage solution for tasks.
package memstore

import (
	// Package Imports
	"fmt"
	"time"
	// Internal Imports
)

// Initialize - Initialize the memstore.
func Initialize() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]Task),
		nextId: 0,
	}
}

// StoreTask - Store a task in memstore.
func (ts *TaskStore) StoreTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Due:  due,
	}

	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextId] = task
	ts.nextId++

	return task.Id
}

// RetrieveById - Retrieve task from memstore by id.
func (ts *TaskStore) RetrieveById(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]

	// CHECK - Task exists
	if ok {
		return task, nil
	}

	// ERROR - Task not found
	return Task{}, fmt.Errorf("ERROR: ID - %d | Task not found", id)
}

// RemoveTask - Remove task from memstore by id.
func (ts *TaskStore) RemoveTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	_, ok := ts.tasks[id]

	// CHECK - Task exists
	if !ok {
		return fmt.Errorf("ERROR: ID - %d | Task not found", id)
	}

	delete(ts.tasks, id)
	return nil
}

// RemoveAllTasks - Remove all tasks from memstore.
func (ts *TaskStore) RemoveAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]Task)
	return nil
}

// RetrieveAllTasks - Retrieve all tasks from memstore.
func (ts *TaskStore) RetrieveAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	allTasks := make([]Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks
}

// RetrieveByTag - Retrieve all tasks from memstore by tag.
func (ts *TaskStore) RetrieveByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				tasks = append(tasks, task)
				continue
			}
		}
	}

	return tasks
}

// RetrieveByDueDate - Retrieve all tasks from memstore by due date.
func (ts *TaskStore) RetrieveByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
