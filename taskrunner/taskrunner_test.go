package taskrunner

import (
	"errors"
	"fmt"
	"github.com/mkboudreau/loggo"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"reflect"
	"testing"
)

func TestTaskRunner(t *testing.T) {
	testLogger := loggo.DefaultLevelLogger()
	testLogger.SetInfo()

	Convey("When a typical task runner implementation is used", t, func() {
		Convey("A string should pass right through", func() {
			taskData, in, out, _ := buildBasicTestTaskData()
			RunTask(taskData, new(TestRunnerImpl))
			in <- "TESTING"
			close(in)

			val, ok := <-out
			So(val, ShouldEqual, "TESTING")
			So(ok, ShouldBeTrue)
		})
		Convey("A struct should pass right through", func() {
			taskData, in, out, _ := buildBasicTestTaskData()
			RunTask(taskData, new(TestRunnerImpl))
			in <- struct{}{}
			close(in)

			val, ok := <-out
			So(val, ShouldHaveSameTypeAs, struct{}{})
			So(val, ShouldNotHaveSameTypeAs, "TESTING")
			So(ok, ShouldBeTrue)
		})
	})

	Convey("When a string task runner implementation is used", t, func() {
		Convey("A string should pass right through", func() {
			taskData, in, out, _ := buildBasicTestTaskData()
			RunTask(taskData, StringTaskRunner(FilterOutHelloWorlds))
			in <- "TESTING"
			close(in)

			val, ok := <-out
			So(val, ShouldEqual, "TESTING")
			So(ok, ShouldBeTrue)
		})
		Convey("A string of Hello World, should end up on the error channel", func() {
			taskData, in, out, err := buildBasicTestTaskData()
			RunTask(taskData, StringTaskRunner(FilterOutHelloWorlds))
			in <- "Hello World"
			close(in)
			defer close(err)

			var okVal, errVal interface{}
			var okOut, okErr bool

			select {
			case okVal, okOut = <-out:
			case errVal, okErr = <-err:
			}
			So(okVal, ShouldBeNil)
			So(errVal, ShouldNotBeNil)
			So(okOut, ShouldBeFalse)
			So(okErr, ShouldBeTrue)
		})
		Convey("A struct should end up as an error", func() {
			taskData, in, out, err := buildBasicTestTaskData()
			RunTask(taskData, StringTaskRunner(FilterOutHelloWorlds))
			in <- struct{}{}
			close(in)
			defer close(err)

			var okVal, errVal interface{}
			var okOut, okErr bool

			select {
			case okVal, okOut = <-out:
			case errVal, okErr = <-err:
			}
			So(okVal, ShouldBeNil)
			So(errVal, ShouldNotBeNil)
			So(okOut, ShouldBeFalse)
			So(okErr, ShouldBeTrue)
		})
	})

	Convey("When a custom string filter is used", t, func() {
			Convey("A string of HELLO should never make it through", func() {
					taskData, in, out, err := buildBasicTestTaskData()
					RunTask(taskData, &FilterString{ Filter: "HELLO" })

					defer close(err)

					go func() {
						in <- "TESTING A"
						in <- "123"
						in <- "HELLO"
						in <- "456"
						in <- "TESTING B"
						close(in)
					}()

					outer: for {
						select {
						case okVal, okOut := <-out:
							testLogger.Info("value",okVal,"ok value",okOut)
							if okOut {
								So(okVal, ShouldNotEqual, "HELLO")
							} else {
								So(okVal, ShouldBeNil)
								break outer
							}
						case errVal, okErr := <-err:
							testLogger.Info("error",errVal,"ok value",okErr)
							if okErr {
								So(fmt.Sprint(errVal), ShouldContainSubstring, "HELLO")
							} else {
								So(errVal, ShouldBeNil)
								break outer
							}
						}
					}
				})
		})


	Convey("When chaining a regular interface task runner to a string task runner", t, func() {
			Convey("A string should pass right through", func() {
					taskData, in, out, _ := buildBasicTestTaskData()
					ChainTasks(taskData, new(TestRunnerImpl), StringTaskRunner(FilterOutHelloWorlds))
					in <- "TESTING"
					close(in)

					val, ok := <-out
					So(val, ShouldEqual, "TESTING")
					So(ok, ShouldBeTrue)
				})
			Convey("A string of Hello World, should end up on the error channel", func() {
					taskData, in, out, err := buildBasicTestTaskData()
					ChainTasks(taskData, new(TestRunnerImpl), StringTaskRunner(FilterOutHelloWorlds))
					in <- "Hello World"
					close(in)
					defer close(err)

					var okVal, errVal interface{}
					var okOut, okErr bool

					select {
					case okVal, okOut = <-out:
					case errVal, okErr = <-err:
					}
					So(okVal, ShouldBeNil)
					So(errVal, ShouldNotBeNil)
					So(okOut, ShouldBeFalse)
					So(okErr, ShouldBeTrue)
				})
			Convey("A struct should end up as an error", func() {
					taskData, in, out, err := buildBasicTestTaskData()
					ChainTasks(taskData, new(TestRunnerImpl), StringTaskRunner(FilterOutHelloWorlds))
					in <- struct{}{}
					close(in)
					defer close(err)

					var okVal, errVal interface{}
					var okOut, okErr bool

					select {
					case okVal, okOut = <-out:
					case errVal, okErr = <-err:
					}
					So(okVal, ShouldBeNil)
					So(errVal, ShouldNotBeNil)
					So(okOut, ShouldBeFalse)
					So(okErr, ShouldBeTrue)
				})
		})

	Convey("When chaining a string task runner to a regular interface task runner", t, func() {
			Convey("A string should pass right through", func() {
					taskData, in, out, _ := buildBasicTestTaskData()
					ChainTasks(taskData, StringTaskRunner(FilterOutHelloWorlds), new(TestRunnerImpl))
					in <- "TESTING"
					close(in)

					val, ok := <-out
					So(val, ShouldEqual, "TESTING")
					So(ok, ShouldBeTrue)
				})
			Convey("A string of Hello World, should end up on the error channel", func() {
					taskData, in, out, err := buildBasicTestTaskData()
					ChainTasks(taskData, StringTaskRunner(FilterOutHelloWorlds), new(TestRunnerImpl))
					in <- "Hello World"
					close(in)
					defer close(err)

					var okVal, errVal interface{}
					var okOut, okErr bool

					select {
					case okVal, okOut = <-out:
					case errVal, okErr = <-err:
					}
					So(okVal, ShouldBeNil)
					So(errVal, ShouldNotBeNil)
					So(okOut, ShouldBeFalse)
					So(okErr, ShouldBeTrue)
				})
			Convey("A struct should end up as an error", func() {
					taskData, in, out, err := buildBasicTestTaskData()
					ChainTasks(taskData, StringTaskRunner(FilterOutHelloWorlds), new(TestRunnerImpl))
					in <- struct{}{}
					close(in)
					defer close(err)

					var okVal, errVal interface{}
					var okOut, okErr bool

					select {
					case okVal, okOut = <-out:
					case errVal, okErr = <-err:
					}
					So(okVal, ShouldBeNil)
					So(errVal, ShouldNotBeNil)
					So(okOut, ShouldBeFalse)
					So(okErr, ShouldBeTrue)
				})
		})

}

func buildBasicTestTaskData() (task *TaskData, in chan interface{}, out chan interface{}, err chan error) {
	in = make(chan interface{})
	out = make(chan interface{})
	err = make(chan error)
	task = &TaskData{
		In:    in,
		Out:   out,
		Error: err,
	}
	return
}

/*
in <- "111111111111111"
in <- "Hello"
in <- "World"
in <- struct{}{}
in <- "Mike"
in <- "Hello World"
in <- "2222222222222222"

in <- struct{}{}

in <- "33333333333"
*/

func AsyncReader(in <-chan interface{}) {
	go func() {
		for obj := range in {
			log.Println(" (test) READING OFF CHANNEL: ", obj)
		}
	}()
}
func AsyncErrReader(in <-chan error) {
	go func() {
		for obj := range in {
			log.Println(" (test) READING OFF ERR CHANNEL: ", obj)
		}
	}()
}


func FilterOutHelloWorlds(in <-chan string, out chan<- string, err chan<- error) {
	for newString := range in {
		log.Println(" (string func) Doing something with string", newString)

		if newString == "Hello World" {
			err <- errors.New(" (string func) I am really sick of Hello World")
		} else {
			out <- newString
		}
	}
}

type TestRunnerImpl struct{}

func (abc *TestRunnerImpl) String() string {
	return "TestRunnerImpl"
}
func (abc *TestRunnerImpl) Run(data *TaskData) {
	defer func() {
		if r := recover(); r != nil {
			data.Error <- fmt.Errorf(" (test runner impl) Problem with run on TestRunnerImpl: %v", r)
		}
	}()
	for obj := range data.In {
		log.Println(" (test runner impl) Found object of type", reflect.TypeOf(obj), "with value", reflect.ValueOf(obj))
		data.Out <- obj
	}
}



type FilterString struct{ Filter string }

func (filter *FilterString) String() string {
	return fmt.Sprintf("FilterString [%v]", filter.Filter)
}
func (filter *FilterString) Run(data *TaskData) {
	defer func() {
		if r := recover(); r != nil {
			data.Error <- fmt.Errorf("%v", filter)
		}
	}()
	for newString := range data.In {
		if newString == filter.Filter {
			data.Error <- fmt.Errorf(" --> filtering %v", filter.Filter )
		} else {
			data.Out <- newString
		}
	}
}
