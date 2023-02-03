package behavioral

import (
	"fmt"
	"testing"
)

var editor =&Editor{text:"i am editor"}
var app = &Application{activeEditor: editor,clipboard: "command test start"}

func TestCommand(t *testing.T) {
	printCommand()
	// 命令
	copyC := NewCopyCommand(editor,app)
	copyC.execute()
	app.history = append(app.history, copyC)
	printCommand()

	app.clipboard = "app new content"
	past := NewPasteCommand(editor,app)
	past.execute()
	app.history = append(app.history, past)
	printCommand()

	cut := NewCutCommand(editor,app)
	cut.execute()
	app.history = append(app.history, cut)
	printCommand()

	fmt.Println(">> Undo Command")
	for i := len(app.history) - 1; i >= 0; i--  {
		app.history[i].undo()
		printCommand()
	}
}

func printCommand()  {
	fmt.Printf("app.clipboard: %s \n",app.clipboard)
	fmt.Printf("editor.selection: %s \n",editor.getSelection())
}
