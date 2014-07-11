package taskrunner

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTaskRunner(t *testing.T) {

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

	Convey("When a custom string filter is used", t, func() {
		Convey("A string of HELLO should never make it through", func() {
			taskData, in, out, err := buildBasicTestTaskData()
			RunTask(taskData, &FilterString{Filter: "HELLO"})

			defer close(err)

			go func() {
				in <- "TESTING A"
				in <- "123"
				in <- "HELLO"
				in <- "456"
				in <- "TESTING B"
				close(in)
			}()

		outer:
			for {
				select {
				case okVal, okOut := <-out:
					if okOut {
						So(okVal, ShouldNotEqual, "HELLO")
					} else {
						So(okVal, ShouldBeNil)
						break outer
					}
				case errVal, okErr := <-err:
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
			data.Error <- fmt.Errorf(" --> filtering %v", filter.Filter)
		} else {
			data.Out <- newString
		}
	}
}
