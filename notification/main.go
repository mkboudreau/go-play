package main

import (
	"log"
	"os/exec"
)

var icondir string = "/home/mkboudreau/Development/ionicons/png/512"

type CommandBuilder interface {
	BuildCommand() *exec.Cmd
}

type CommandData struct {
	Command string
	Args    []string
}

func NewGnomeNotifier(iconPath string, title string, text string) CommandBuilder {
	return &CommandData{
		Command: "notify-send",
		Args:    []string{"-i", iconPath, title, text},
	}
}

func (command *CommandData) BuildCommand() *exec.Cmd {
	return exec.Command(command.Command, command.Args...)
}

func main() {
	icon := icondir + "/social-apple-outline.png" //"/usr/share/icons/gnome/scalable/emotes/face-monkey.svg"
	title := "Go Notification"
	text := "oooh ooooh aah aaah"

	cmdBuilder := NewGnomeNotifier(icon, title, text)

	cmd := cmdBuilder.BuildCommand() //exec.Command("notify-send", "-i", icon, title, text)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
