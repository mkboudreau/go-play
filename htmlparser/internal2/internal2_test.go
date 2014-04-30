package internal2_test

import (
	. "github.com/mkboudreau/reservations/htmlparser/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"reflect"
	"testing"
)

func TestInternalHtmlParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html Parsing INTERNALS Test Suite")
}

var _ = Describe("Html Parsing Internal Tests", func() {

	Context("INTERNAL Regular Nodes ONLY: Valid HTML Text", func() {
		var parser *InternalHtmlParsingEngine

		BeforeEach(func() {
			parser = NewInternalHtmlParsingEngine()
		})
		AfterEach(func() {
			parser.TurnOffLogging()
		})
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleRegularNodeWithContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleRegularNodeWithIdAndContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 1 node", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleNoContentNodeWithId)
			Expect(nodes).To(HaveLen(1))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
		})
		It("Should contain 1 node", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleNoContentNode)
			Expect(nodes).To(HaveLen(1))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
		})
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlMultipleNoContentNode)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[1].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))

			if parser.IsDebugging() {
				fmt.Println(nodes)
			}
		})
		It("Should correctly process the full html set", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlAllVariationTest)
			Expect(nodes).ToNot(BeNil())
			Expect(nodes).ToNot(HaveLen(0))
		})
	})
	Context("INTERNAL Regular Nodes ONLY: Invalid HTML Text", func() {
		var parser *InternalHtmlParsingEngine

		BeforeEach(func() {
			parser = NewInternalHtmlParsingEngine()
		})
		AfterEach(func() {
			parser.TurnOffLogging()
		})
		It("Should panic on no opening start tag", func(done Done) {
			expectFunctionToPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningStartTag) })
		})
		It("Should panic on no opening end tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningEndTag) })
		})
		It("Should panic on no closing start tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoClosingStartTag) })
		})
		It("Should not panic on no closing end tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoClosingEndTag) })
		})
		It("Should panic on no opening slash inside end tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningEndTagSlash) })
		})
	})
	Context("INTERNAL Ensure Capacity Utility Function", func() {
		var parser *InternalHtmlParsingEngine

		BeforeEach(func() {
			parser = NewInternalHtmlParsingEngine()
		})
		AfterEach(func() {
			parser.TurnOffLogging()
		})
		// func (slice *InternalParseNodes) ensureCapacity(neededLength int) {

		It("Should Add One To Slice", func() {
			slice := make(InternalParseNodes, 1)
			Expect(slice).To(HaveLen(1))
			slice.EnsureCapacity(2)
			Expect(slice).To(HaveLen(2))
		})
		It("Should Add Two To Slice", func() {
			slice := make(InternalParseNodes, 1)
			Expect(slice).To(HaveLen(1))
			slice.EnsureCapacity(3)
			Expect(slice).To(HaveLen(3))
		})
		It("Should Add NONE To Slice", func() {
			slice := make(InternalParseNodes, 1)
			Expect(slice).To(HaveLen(1))
			slice.EnsureCapacity(1)
			Expect(slice).To(HaveLen(1))
		})
		It("Should Add NONE To 10 item Slice", func() {
			slice := make(InternalParseNodes, 10)
			Expect(slice).To(HaveLen(10))
			slice.EnsureCapacity(1)
			Expect(slice).To(HaveLen(10))
		})
		It("Should Add 0 To 10 item Slice", func() {
			slice := make(InternalParseNodes, 10)
			Expect(slice).To(HaveLen(10))
			slice.EnsureCapacity(0)
			Expect(slice).To(HaveLen(10))
		})
	})

	Context("INTERNAL Adding Slice At Index Utility Function", func() {
		var parser *InternalHtmlParsingEngine

		BeforeEach(func() {
			parser = NewInternalHtmlParsingEngine()
		})
		AfterEach(func() {
			parser.TurnOffLogging()
		})
		// func (slice *InternalParseNodes) AddSliceAtIndex(newSlice InternalParseNodes, index int) (nextIndex int) {

		It("Should Simply Append (by expanding)", func() {
			slice := make(InternalParseNodes, 10)
			sliceToAdd := make(InternalParseNodes, 10)
			addAtIndex := 10
			expectedNextIndex := 20
			expectedNewLength := 20
			actualNextIndex := slice.AddSliceAtIndex(sliceToAdd, addAtIndex)
			actualNewLength := len(slice)

			Expect(actualNewLength).To(BeNumerically(">", expectedNewLength))
			Expect(actualNextIndex).To(Equal(expectedNextIndex))
		})
		It("Should Simply Fill Existing (no expanding)", func() {
			slice := make(InternalParseNodes, 10)
			sliceToAdd := make(InternalParseNodes, 10)
			addAtIndex := 0
			expectedNextIndex := 10
			expectedNewLength := 10
			actualNextIndex := slice.AddSliceAtIndex(sliceToAdd, addAtIndex)
			actualNewLength := len(slice)

			Expect(actualNewLength).To(BeNumerically(">", expectedNewLength))
			Expect(actualNextIndex).To(Equal(expectedNextIndex))
		})
		It("Should Add In Middle and Properly Expand", func() {
			slice := make(InternalParseNodes, 10)
			sliceToAdd := make(InternalParseNodes, 10)
			addAtIndex := 3
			expectedNextIndex := 13
			expectedNewLength := 13
			actualNextIndex := slice.AddSliceAtIndex(sliceToAdd, addAtIndex)
			actualNewLength := len(slice)

			Expect(actualNewLength).To(BeNumerically(">", expectedNewLength))
			Expect(actualNextIndex).To(Equal(expectedNextIndex))
		})
		It("Should Add Real Data", func() {
			slice := make(InternalParseNodes, 1)
			sliceToAdd := make(InternalParseNodes, 2)
			slice[0] = InternalParseNode{Type: InternalParseContent, Content: "I'm Alone"}
			sliceToAdd[0] = InternalParseNode{Type: InternalParseContent, Content: "Hello"}
			sliceToAdd[1] = InternalParseNode{Type: InternalParseContent, Content: "World"}

			addAtIndex := 1
			expectedNextIndex := 3
			expectedNewLength := 3
			actualNextIndex := slice.AddSliceAtIndex(sliceToAdd, addAtIndex)
			actualNewLength := len(slice)

			Expect(actualNewLength).To(BeNumerically(">", expectedNewLength))
			Expect(actualNextIndex).To(Equal(expectedNextIndex))
			Expect(slice[actualNextIndex]).ToNot(BeNil())
			Expect(slice[actualNextIndex].Content).To(BeEmpty())
			Expect(slice[actualNextIndex-1].Content).To(Equal("World"))
		})
	})

	Context("INTERNAL Full Node Processing: Valid HTML Text", func() {
		var parser *InternalHtmlParsingEngine

		BeforeEach(func() {
			parser = NewInternalHtmlParsingEngine()
		})
		AfterEach(func() {
			parser.TurnOffLogging()
		})

		It("Should contain 3 nodes after parsing a single regular node with content ", func() {
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlSingleRegularNodeWithContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 3 nodes after parsing a single regular node with content and an ID", func() {
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlSingleRegularNodeWithIdAndContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 1 node after parsing a single no content node with an ID", func() {
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlSingleNoContentNodeWithId)
			Expect(nodes).To(HaveLen(1))
			Expect(nodes[0].Type).To(Equal(InternalParseNoContentTag))
		})
		It("Should contain 1 node after parsing a single no content node", func() {
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlSingleNoContentNode)
			Expect(nodes).To(HaveLen(1))
			Expect(nodes[0].Type).To(Equal(InternalParseNoContentTag))
		})
		It("Should contain 3 nodes after parsing a multiple no content node", func() {
			parser.TurnUpLogging()
			parser.TurnUpLogging()
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlMultipleNoContentNode)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseNoContentTag))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should correctly process the full html set", func() {
			nodes := parser.TokenizeHtmlStringIntoParseNodes(HtmlAllVariationTest)
			Expect(nodes).ToNot(BeNil())
			Expect(nodes).ToNot(HaveLen(0))
		})
	})

})

const HtmlSingleRegularNodeWithContent = "<div>Hello</div>"
const HtmlSingleRegularNodeWithIdAndContent = `<div id="world">Hello</div>`
const HtmlSingleNoContentNodeWithId = `<div id="world" />`
const HtmlSingleNoContentNode = `<div />`
const HtmlMultipleNoContentNode = `<div id="hello"><div /></div>`

const InvalidHtmlNoOpeningStartTag = `div id="hello">Hello World</div>`
const InvalidHtmlNoOpeningEndTag = `<div id="hello"Hello World</div>`
const InvalidHtmlNoClosingStartTag = `<div id="hello">Hello World/div>`
const InvalidHtmlNoClosingEndTag = `<div id="hello">Hello World</div`
const InvalidHtmlNoOpeningEndTagSlash = `<div id="hello">Hello World<div>`
const InvalidHtmlPlainString = `html`

const HtmlWithDocTypeTest = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html>
<head><title>Hello</title></head>
<body>
<h1>Heading</h1>
</body>
</html>
`

const HtmlWithCommentsTest = `
<html>
<head><title>Hello</title></head>
<body>
<!-- Watch Out! -->
<h1>Heading</h1>

<!-- <h2>This should not be a node</h2> -->
</body>
</html>

`

const HtmlAllVariationTest = `

<html>
<head><title>Hello</title>

</head>

<body>

<div id="id1" class="someclass"> <h1>Test H1</h1></div>

<div id="outer" >
	<div id="middle1" abc="aaa">Test</div>
	<div id="middle2" abc="bbb">Test 2</div>
	<div id="middle3" abc="ccc ddd">Test 3</div>
	<div id="middle4" ><div id="nocontent" /></div>
</div>
<div></div>

<div>       </div>
</body>
</html>

`

/* private helper functions */

func expectFunctionToPanic(done Done, fn interface{}) {
	expectFunctionToPanicAsSpecified(done, fn, true)
}
func expectFunctionToNotPanic(done Done, fn interface{}) {
	expectFunctionToPanicAsSpecified(done, fn, false)
}
func expectFunctionToPanicAsSpecified(done Done, fn interface{}, shouldPanic bool) {
	go func() {
		defer GinkgoRecover()

		panicFunc := func() {
			val := reflect.ValueOf(fn)
			if val.Kind() == reflect.Func {
				in := make([]reflect.Value, 0)
				val.Call(in)
			}
		}

		if shouldPanic {
			Expect(panicFunc).To(Panic())
		} else {
			Expect(panicFunc).ToNot(Panic())
		}

		close(done)
	}()
}
