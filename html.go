package golem

type Props struct {
	InnerText     string
	Id            string
	Class         string
	EventListener map[string]string
}

func (p *Props) Text(text string) Props {
	if p == nil {
		p = &Props{}
	}
	p.InnerText = text
	return *p
}

func createElem(kind string, props Props, children ...Element) Element {
	return Element{
		TagName:       kind,
		Content:       props.InnerText,
		Class:         props.Class,
		Id:            props.Id,
		EventListener: props.EventListener,
		Children:      children,
	}
}

func H1(props Props, children ...Element) Element {
	return createElem("h1", props, children...)
}

func H2(props Props, children ...Element) Element {
	return createElem("h2", props, children...)
}

func H3(props Props, children ...Element) Element {
	return createElem("h3", props, children...)
}

func H4(props Props, children ...Element) Element {
	return createElem("h4", props, children...)
}

func H5(props Props, children ...Element) Element {
	return createElem("h5", props, children...)
}

func P(props Props, children ...Element) Element {
	return createElem("p", props, children...)
}

func Div(props Props, children ...Element) Element {
	return createElem("div", props, children...)
}

func Input(props Props, children ...Element) Element {
	return createElem("input", props, children...)
}

func A(props Props, children ...Element) Element {
	return createElem("a", props, children...)
}

func Button(props Props, children ...Element) Element {
	return createElem("button", props, children...)
}

func Img(props Props, children ...Element) Element {
	return createElem("img", props, children...)
}

func Table(props Props, children ...Element) Element {
	return createElem("table", props, children...)
}

func THead(props Props, children ...Element) Element {
	return createElem("thead", props, children...)
}

func TBody(props Props, children ...Element) Element {
	return createElem("tbody", props, children...)
}

func Tr(props Props, children ...Element) Element {
	return createElem("tr", props, children...)
}

func Td(props Props, children ...Element) Element {
	return createElem("td", props, children...)
}

func Ul(props Props, children ...Element) Element {
	return createElem("ul", props, children...)
}

func Ol(props Props, children ...Element) Element {
	return createElem("ol", props, children...)
}

func Li(props Props, children ...Element) Element {
	return createElem("li", props, children...)
}
