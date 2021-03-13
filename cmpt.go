package golem

const Empty Object = ""

type Object string

type DOM interface {
	CreateElement(TagName string) (Object, error)
	AddClass(o Object, class string) error
	AddId(o Object, id string) error
	SetInnerText(o Object, txt string) error
	AppendChild(parent, child Object) error
	AddEventListner(o Object, event string, name string) error
	GetBody() Object
}

type Element struct {
	TagName       string
	Id            string
	Content       string
	Class         string
	EventListener map[string]string
	Children      []Element
}

func Merge(a, b map[string]Object) map[string]Object {
	out := map[string]Object{}
	for id, o := range a {
		out[id] = o
	}
	for id, o := range b {
		out[id] = o
	}
	return out
}

func Render(parent Object, e Element, d DOM) (map[string]Object, error) {
	rt := map[string]Object{}

	if parent == Empty {
		parent = d.GetBody()
	}

	o, err := d.CreateElement(e.TagName)
	if err != nil {
		return nil, err
	}
	if e.Content != "" {
		d.SetInnerText(o, e.Content)
	}
	if e.Class != "" {
		d.AddClass(o, e.Class)
	}
	if e.Id != "" {
		d.AddId(o, e.Id)
		rt[e.Id] = o
	}
	for e, name := range e.EventListener {
		d.AddEventListner(o, e, name)
	}
	err = d.AppendChild(parent, o)
	if err != nil {
		return nil, err
	}
	for _, c := range e.Children {
		ids, err := Render(o, c, d)
		if err != nil {
			return nil, err
		}
		rt = Merge(rt, ids)
	}

	return rt, nil
}
