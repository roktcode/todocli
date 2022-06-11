package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

// default file name
var todoFilename = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the to-do list app\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usga information:")
		flag.PrintDefaults()
	}

	// Parsing flags for the cli
	add := flag.Bool("add", false, "Add task to ToDo list")
	list := flag.Bool("list", false, "List all the tasks")
	complete := flag.Int("complete", 0, "Mark the task as complete")
	// add delete
	del := flag.Int("del", 0, "Delete a task from the list")
	// add verbose flag
	verbose := flag.Bool("verbose", false, "Show information about tasks")
	// add an only-completed flag
	notCompleted := flag.Bool("not-completed", false, "Show only uncompleted tasks")
	flag.Parse()

	// Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFilename = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	// Use the Get method to read to do list items from file
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// print all the tasks
		if *notCompleted {
			// print all not completed items
			if *verbose {
				fmt.Print(l.Verbose(l.GetUncompleted()))
			} else {
				fmt.Println(l.GetUncompleted())
			}
		} else {
			// print all items
			if *verbose {
				fmt.Print(l.Verbose(*l))
			} else {
				fmt.Print(l)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// save the fie after complete task
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// read the value from stdin (using piping)
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(t)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// save the fie after complete task
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// create a scanner to read from io.Reader provided to getTask
	s := bufio.NewScanner(r)
	s.Scan()

	// if an error occurred, return it
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		// fmt.Errorf returns an error
		return "", fmt.Errorf("Task can't be blank")
	}

	return s.Text(), nil
}
