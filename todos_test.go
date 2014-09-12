package todo

import "testing"

func TestTodo_Complete(t *testing.T) {
	td := &Todo{}
	td.Complete()

	if td.CompletedAt == nil {
		t.Errorf("Expected CompletedAt to be non nil.")
	}
}

func TestTodo_Uncomplete(t *testing.T) {
	td := &Todo{}
	td.Complete()
	td.Uncomplete()

	if td.CompletedAt != nil {
		t.Errorf("Expected CompletedAt to be nil.")
	}
}
