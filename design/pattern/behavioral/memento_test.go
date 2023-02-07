package behavioral

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &Caretaker{
		mementoArray: make([]Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento())
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento())
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento())
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
