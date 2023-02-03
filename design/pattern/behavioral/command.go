package behavioral

import "fmt"

/**
行为型模式
命令模式，特点是可以把任何操作转换为对象，并可序列化，可存储，可传递，可撤销和回滚。
需要提供
1 command interface
2 concrete commands:持有入参和receiver对象
3 receiver
4 sender
 */

//编辑器作为receiver
type Editor struct {
	text string
}
func (p *Editor) getSelection() string { return p.text }
func (p *Editor) deleteSelection(){p.text = ""}
func (p *Editor) replaceSelection(text string) { p.text = text}

//Application作为发送着
type Application struct {
	clipboard string
	activeEditor *Editor
	history []Command
}

//定义Command Interface
type Command interface {
	saveBackup()
	undo()
	execute()
}
type BaseCommand struct {
	editor *Editor
	app *Application
	backup string
}
func (p *BaseCommand) saveBackup() {
	p.backup = p.editor.text
}
func (p *BaseCommand) undo() {
	p.editor.text = p.backup
}

type CopyCommand struct {
	BaseCommand
}
func (p *CopyCommand) execute() {
	fmt.Println(">> CopyCommand Execute")
	p.backup = p.editor.text
	p.app.clipboard = p.editor.getSelection()
}
func NewCopyCommand(editor *Editor,app *Application) Command{
	return &CopyCommand{BaseCommand{
		editor: editor,
		app:    app,
		backup: "",
	}}
}

type CutCommand struct {
	BaseCommand
}
func (p *CutCommand) execute()  {
	fmt.Println(">> CutCommand Execute")
	p.saveBackup()
	p.app.clipboard = p.editor.getSelection()
	p.editor.deleteSelection()
}
func NewCutCommand(editor *Editor,app *Application) Command{
	return &CutCommand{BaseCommand{
		editor: editor,
		app:    app,
		backup: "",
	}}
}

type PasteCommand struct {
	BaseCommand
}
func (p *PasteCommand) execute()  {
	fmt.Println(">> PasteCommand Execute")
	p.saveBackup()
	p.editor.replaceSelection(p.app.clipboard)
}
func NewPasteCommand(editor *Editor,app *Application) Command{
	return &PasteCommand{BaseCommand{
		editor: editor,
		app:    app,
		backup: "",
	}}
}
