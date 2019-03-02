package dom

import (
	"syscall/js"

	"github.com/syumai/gq"
)

var document = gq.NewElement(jsValue(js.Global().Get("document")))

func CreateElement(tagName string, attrs map[string]interface{}) gq.Element {
	// TODO
	return nil
}

func CreateNode(tagName string, attrs map[string]interface{}) gq.Node {
	// TODO
	return nil
}

func Query(q string) gq.Element {
	return document.Query(q)
}

func QueryAll(q string) []gq.Element {
	return document.QueryAll(q)
}
