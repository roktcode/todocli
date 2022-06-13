package todo

import (
	"os"
	"testing"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	// create empty to-do list
	l := List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, l[0].Task)
	}

}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New Task shouldn't be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, val := range tasks {
		l.Add(val)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected list length %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	taskName := "New task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q",
			taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to a file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match task %q", l1[0].Task, l2[0].Task)
	}

}
