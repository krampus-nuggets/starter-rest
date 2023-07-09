// MemStore Types - Types for the MemStore module.
package memstore

import (
	// Package Imports
	"sync"
	"time"
	// Internal Imports
)

// Structure | Task
// ================
// Id - Unique identifier for the task | Type: int
// Text - Description of the task | Type: string
// Tags - List of tags associated with the task | Type: []string
// Due - Due date of the task | Type: time.Time

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// TaskStore - In-memory DB structure for tasks.
type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}
