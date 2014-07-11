package taskrunner

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringTaskRunner(t *testing.T) {

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
}

func FilterOutHelloWorlds(in <-chan string, out chan<- string, err chan<- error) {
	for newString := range in {
		if newString == "Hello World" {
			err <- errors.New(" (string func) I am really sick of Hello World")
		} else {
			out <- newString
		}
	}
}
