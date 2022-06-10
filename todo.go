package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

//String prints out a formatted list
//Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, v := range *l {

		prefix := "  "

		if v.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers starting from 1 instead of
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, v.Task)
	}

	return formatted
}

// a to-do item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	completedAt time.Time
}

// List holds all to-do items
type List []item

// Add method - adds an item to the list
func (l *List) Add(task string) {
	// construct a type
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		completedAt: time.Time{},
	}

	// add to the list
	*l = append(*l, t)
}

// Complete method - marks a todo item as completed
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].completedAt = time.Now()

	return nil
}

// i am learning new things i didn't know before, so be coool

// Delete method deletes an item from the to-do list
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	// what does ... do here?
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method, saves the to-do lsit to a json file
func (l *List) Save(filename string) error {
	// encode the pointer to the
	// 	... list into a json format, or produces an error
	js, err := json.Marshal(l)

	// return an error if encoding into json fails
	if err != nil {
		return err
	}

	// write the js to a filename on the disk, and return the result
	// err(fail) or nil (success)
	return os.WriteFile(filename, js, 0644)
}

// Get method, reads the list data from a file and decodes the json
// into a list again
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	// parses the json and stores the data into the pointer to the source
	// or return err != nil instead
	return json.Unmarshal(file, l)
}
