package internal2

import (
	"log"
	"os"
	s "strings"
)

type InternalParseType int

const (
	InternalParseContent InternalParseType = iota
	InternalParseEndTag
	InternalParseStartTag
	InternalParseNoContentTag
)

type InternalParseNode struct {
	Type    InternalParseType
	Content string
}

type InternalParseNodes []InternalParseNode


type InternalHtmlParsingEngine struct {
	StartOpenTagToken  string
	StartCloseTagToken string
	EndTagToken        string
	EndNoTagToken      string
	loggingLevel       loggingLevel
	logger             *log.Logger
}

func (parseType InternalParseType) String() string {
	switch parseType {
	case InternalParseContent:
		return "InternalParseType"
	case InternalParseEndTag:
		return "InternalParseEndTag"
	case InternalParseStartTag:
		return "InternalParseStartTag"
	case InternalParseNoContentTag:
		return "InternalParseNoContentTag"
	}
	return ""
}

type loggingLevel int

const (
	loggingOff loggingLevel = iota
	debugLevel
	traceLevel
)

func NewInternalHtmlParsingEngine() *InternalHtmlParsingEngine {
	engine := new(InternalHtmlParsingEngine)
	engine.StartOpenTagToken = "<"
	engine.StartCloseTagToken = "</"
	engine.EndNoTagToken = "/>"
	engine.EndTagToken = ">"
	engine.loggingLevel = loggingOff
	engine.logger = buildDefaultLogger()

	return engine
}
func buildDefaultLogger() *log.Logger {
	return log.New(os.Stdout, "Internal Html Parsing >> ", log.Lmicroseconds)
}

func (engine *InternalHtmlParsingEngine) TurnUpLogging() *InternalHtmlParsingEngine {
	engine.loggingLevel++
	return engine
}
func (engine *InternalHtmlParsingEngine) TurnDownLogging() *InternalHtmlParsingEngine {
	engine.loggingLevel--
	return engine
}
func (engine *InternalHtmlParsingEngine) TurnOffLogging() *InternalHtmlParsingEngine {
	engine.loggingLevel = loggingOff
	return engine
}
func (engine *InternalHtmlParsingEngine) IsDebugging() bool {
	return engine.loggingLevel >= debugLevel
}
func (engine *InternalHtmlParsingEngine) IsTracing() bool {
	return engine.loggingLevel >= traceLevel
}
func (engine *InternalHtmlParsingEngine) debug(msg ...interface{}) *InternalHtmlParsingEngine {
	if engine.loggingLevel >= debugLevel {
		engine.logger.Println(msg)
	}

	return engine
}
func (engine *InternalHtmlParsingEngine) trace(msg ...interface{}) *InternalHtmlParsingEngine {
	if engine.loggingLevel >= traceLevel {
		engine.logger.Println(msg)
	}

	return engine
}

func (engine *InternalHtmlParsingEngine) TokenizeHtmlStringIntoParseNodes(html string) InternalParseNodes {
	engine.debug("ANY HTML [", html, "]")
	token := engine.EndNoTagToken
	splitAfterEnds := s.SplitAfter(html, token)
	if len(splitAfterEnds) == 1 {
		engine.debug("Did not find [", token, "] for splitting main html string")
		return engine.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(html)
	} else if len(splitAfterEnds) == 2 {
		if (splitAfterEnds[1]) // not empty... then recurse more
		engine.debug("Found 2 parts. No Content Part [", splitAfterEnds[0], "]; Normal Part [", splitAfterEnds[1], "]")
		return engine.buildParseNodesFromNoContentTagSplit(splitAfterEnds, s.Count(html, token))
	} else {
		nodeList := make(InternalParseNodes, 0)
		for _, part := range splitAfterEnds {
			newNodeList := engine.TokenizeHtmlStringIntoParseNodes(part)
			nodeList.AppendSlice(newNodeList)
		}
		return nodeList
		//engine.debug("Found ["+token+"] for splitting main html string into [", len(splitAfterEnds), "] parts")
		//return engine.buildParseNodesFromNoContentTagSplit(splitAfterEnds, s.Count(html, token))
	}
}

func (engine *InternalHtmlParsingEngine) buildParseNodesFromNoContentTagSplit(noContentEndSplit []string, estimatedTotal int) InternalParseNodes {
	parseNodes := make(InternalParseNodes, estimatedTotal)
	currentIndex := 0
	for _, stringWithStartChar := range noContentEndSplit {
		var tag, content string

		engine.debug("(no content) stringWithStartChar:", stringWithStartChar)
		if len(stringWithStartChar) == 0 || len(s.TrimSpace(stringWithStartChar)) == 0 {
			engine.trace("   * (no content) ignoring")
			continue
		}
		engine.trace("   - (no content) continuing")

		if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartOpenTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]

			tokenizedRegularNodes := engine.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(string(content))
			engine.trace("   - (no content) currentIndex (before adding slice):", currentIndex)
			engine.trace("   - (no content) parseNodes length (before adding slice):", len(parseNodes), "; capacity:", cap(parseNodes))
			engine.trace("   - (no content) tokenizedRegularNodes length (before adding slice):", len(tokenizedRegularNodes), "; capacity:", cap(tokenizedRegularNodes))
			currentIndex = parseNodes.AddSliceAtIndex(tokenizedRegularNodes, currentIndex)
			engine.trace("   - (no content) currentIndex (after adding slice):", currentIndex)
			engine.trace("   - (no content) parseNodes length (after adding slice):", len(parseNodes), "; capacity:", cap(parseNodes))
			engine.trace("   - (no content) tokenizedRegularNodes length (after adding slice):", len(tokenizedRegularNodes), "; capacity:", cap(tokenizedRegularNodes))
			engine.trace("   - (no content) tag:", tag)
			parseNodes.EnsureCapacity(currentIndex + 1)
			parseNodes[currentIndex] = InternalParseNode{Type: InternalParseNoContentTag, Content: tag}
			currentIndex++
		} else {
			engine.debug("PANIC!", "stringWithStartChar:", stringWithStartChar, "currentIndex:", currentIndex, "noContentEndSplit length:", len(noContentEndSplit))
			panic("no start tag found for end tag of contentless element")
		}
	}

	return parseNodes[:currentIndex]
}

func (slice *InternalParseNodes) AppendSlice(newSlice InternalParseNodes) {
	if len(newSlice) > 0 {
		newTargetSize := len(newSlice) + len(*slice)
		slice.EnsureCapacity(newTargetSize)
		tmpSlice := *slice
		copy(tmpSlice[:], newSlice)
	}
}
func (slice *InternalParseNodes) AddSliceAtIndex(newSlice InternalParseNodes, index int) (nextIndex int) {
	if len(newSlice) > 0 {
		newTargetSize := (len(newSlice) * 2) + index
		slice.EnsureCapacity(newTargetSize)
		tmpSlice := *slice
		copy(tmpSlice[index:], newSlice)
		nextIndex = index + len(newSlice)
	} else {
		nextIndex = index
	}

	return nextIndex
}

func (slice *InternalParseNodes) EnsureCapacity(neededLength int) {
	if neededLength > cap(*slice) {
		oldValues := *slice
		*slice = make(InternalParseNodes, neededLength, neededLength*2)
		copy(*slice, oldValues)
	} else if neededLength > len(*slice) {
		oldValues := *slice
		*slice = oldValues[:neededLength]
	}
}

func (engine *InternalHtmlParsingEngine) TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(html string) InternalParseNodes {
	engine.debug("REGULAR HTML [", html, "]")
	token := engine.EndTagToken
	splitAfterEnds := s.SplitAfter(html, token)

	if len(splitAfterEnds) == 1 {
		engine.debug("Did not find [", token, "] for splitting regular tags")
		return make(InternalParseNodes, 0)
	} else {
		engine.debug("Found [", token, "] for splitting regular tags into [", len(splitAfterEnds), "] parts")
		return engine.buildParseNodesFromContentTagSplit(splitAfterEnds)
	}
}

func (engine *InternalHtmlParsingEngine) buildParseNodesFromContentTagSplit(contentEndSplit []string) InternalParseNodes {
	parseNodes := make(InternalParseNodes, len(contentEndSplit)*2)
	currentIndex := 0
	for _, stringWithStartChar := range contentEndSplit {
		var tag, content string
		var tagType InternalParseType

		engine.debug("(content) stringWithStartChar:", stringWithStartChar)
		if len(stringWithStartChar) == 0 || len(s.TrimSpace(stringWithStartChar)) == 0 {
			engine.trace("   * (content) ignoring")
			continue
		}
		engine.trace("   - (content) continuing")

		if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartCloseTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]
			tagType = InternalParseEndTag
		} else if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartOpenTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]
			tagType = InternalParseStartTag
		} else {
			engine.debug("PANIC!", "stringWithStartChar:", stringWithStartChar, "currentIndex:", currentIndex, "contentEndSplit length:", len(contentEndSplit))
			panic("no start tag found for end tag for content element")
		}

		if len(content) != 0 {
			engine.trace("   - (content) content:", content)
			parseNodes.EnsureCapacity(currentIndex + 1)
			parseNodes[currentIndex] = InternalParseNode{Type: InternalParseContent, Content: content}
			currentIndex++
		}
		engine.trace("   - (content) tag:", tag)
		parseNodes.EnsureCapacity(currentIndex + 1)
		parseNodes[currentIndex] = InternalParseNode{Type: tagType, Content: tag}
		currentIndex++
	}

	return parseNodes[:currentIndex]
}

func (engine *InternalHtmlParsingEngine) buildTagNodeFromString(tag string, content string) *InternalParseNode {
	var tagType InternalParseType
	if openTagIndex := s.LastIndex(tag, engine.StartCloseTagToken); openTagIndex != -1 {
		tagType = InternalParseEndTag
	} else if openTagIndex := s.LastIndex(tag, engine.StartOpenTagToken); openTagIndex != -1 {
		tagType = InternalParseStartTag
	} else {
		engine.debug("PANIC!", "tag:", tag, "tag type:", tagType)
		panic("no start tag found for end tag for content element")
	}
	return &InternalParseNode{Type: tagType, Content: content}
}
