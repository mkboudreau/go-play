package taskrunner

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestChainingMultipleTypesOfTaskRunner(t *testing.T) {

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
