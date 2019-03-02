package gq

type Element interface {
	Node
	Querier
	Manipulator
}

type Node interface {
	Value
	GetValue() Value
}

type element struct {
	Value
}

func NewElement(v Value) Element {
	return &element{v}
}

type Querier interface {
	Query(q string) Element
	QueryAll(q string) []Element
}

func (e *element) Query(q string) Element {
	v := e.Value.Call("querySelector", q)
	return NewElement(v)
}

func (e *element) QueryAll(q string) []Element {
	v := e.Value.Call("querySelectorAll", q)
	elements := make([]Element, v.Length())
	for i := 0; i < v.Length(); i++ {
		elements[i] = NewElement(v.Index(i))
	}
	return elements
}

type Manipulator interface {
	Append(node Node)
	Prepend(node Node)
	InsertBefore(node Node, child Node)
}

func (e *element) Append(node Node) {
	e.Call("appendChild", node.GetValue())
}

func (e *element) Prepend(node Node) {
	child := e.Get("firstChild")
	e.InsertBefore(node, &element{child}) // FIXME: this can be textNode
}

func (e *element) InsertBefore(node Node, child Node) {
	e.Call("insertBefore", node.GetValue(), child.GetValue())
}

func (e *element) GetValue() Value {
	return e.Value
}
