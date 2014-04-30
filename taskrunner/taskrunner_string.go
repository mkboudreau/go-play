package taskrunner

import (
	"errors"
	"fmt"
	"reflect"
)

type StringTaskRunner func(in <-chan string, out chan<- string, err chan<- error)

func (runner StringTaskRunner) Run(data *TaskData) {
	defer func() {
		if r := recover(); r != nil {
			data.Error <- errors.New(fmt.Sprintf(" (string task runner) Caught error trying to run in a string task runner: %v", r))
		}
	}()

	inString := make(chan string)
	outString := make(chan string)

	go func() {
		adaptInterfaceChannelToStringChannel(data.In, (chan<- string)(inString), data.Error)
	}()
	go func() {
		adaptStringChannelToInterfaceChannel((<-chan string)(outString), data.Out, data.Error)
	}()

	runner(inString, outString, data.Error)
}

func adaptInterfaceChannelToStringChannel(in <-chan interface{}, out chan<- string, err chan<- error) {
	defer close(out)

	for hopefullyString := range in {
		if reflect.TypeOf(hopefullyString).Kind() == reflect.String {
			out <- hopefullyString.(string)
		} else {
			err <- errors.New(fmt.Sprintf(" (string task runner) Could not adapt interface %v to string", hopefullyString))
		}
	}
}

func adaptStringChannelToInterfaceChannel(in <-chan string, out chan<- interface{}, err chan<- error) {
	defer close(out)

	for someString := range in {
		out <- someString
	}
}
