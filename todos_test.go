package todo

import "testing"

func TestTodoComplete(t *testing.T) {
	td := &Todo{}
	td.Complete()

	if td.CompletedAt == nil {
		t.Errorf("Expected CompletedAt to be non nil.")
	}
}

func TestTodoUncomplete(t *testing.T) {
	td := &Todo{}
	td.Complete()
	td.Uncomplete()

	if td.CompletedAt != nil {
		t.Errorf("Expected CompletedAt to be nil.")
	}
}
