package goelm

import (
	"fmt"
)

type JavascriptDOM struct {
	events         map[Object]string
	out            string
	count          int
	styleFuncAdded bool
	EventFunc      func(string) func()
}

func (js *JavascriptDOM) Apply(t Target) (string, error) {
	if js.EventFunc != nil {
		for o, name := range js.events {
			err := t.Bind(string(o), js.EventFunc(name))
			if err != nil {
				return "", err
			}
		}
		js.events = nil
	}
	out := js.out
	js.out = ""
	err := t.Eval(out)
	return out, err
}

func (js *JavascriptDOM) newObject() Object {
	js.count++
	return Object(fmt.Sprintf("item%d", js.count))
}

func (js *JavascriptDOM) Raw(code string) {
	js.out += code
}

func (js *JavascriptDOM) CreateElement(TagName string) (Object, error) {
	o := js.newObject()
	js.out += fmt.Sprintf("var %v = document.createElement('%v')\n", o, TagName)
	return o, nil
}

func (js *JavascriptDOM) AddClass(o Object, class string) error {
	js.out += fmt.Sprintf("%v.className += ' '  + '%s'\n", o, class)
	return nil
}

func (js *JavascriptDOM) AddId(o Object, id string) error {
	js.out += fmt.Sprintf("%v.id += ' '  + '%s'\n", o, id)
	return nil
}

func (js *JavascriptDOM) SetInnerText(o Object, txt string) error {
	js.out += fmt.Sprintf("%v.innerText = '%s'\n", o, txt)
	return nil
}
func (js *JavascriptDOM) AppendChild(parent, child Object) error {
	js.out += fmt.Sprintf("%v.appendChild(%v)\n", parent, child)
	return nil
}

func (js *JavascriptDOM) AddEventListner(o Object, event string, name string) error {
	if js.events == nil {
		js.events = map[Object]string{}
	}
	eventObject := js.newObject()
	js.events[eventObject] = name
	js.out += fmt.Sprintf("%v.addEventListener('%s', async () => {await %v();});\n", o, event, eventObject)
	return nil
}

func (js *JavascriptDOM) GetBody() Object {
	return Object("document.body")
}

func (js *JavascriptDOM) AddStyle(style string) {
	if !js.styleFuncAdded {
		js.addStyleFunc()
	}
	js.out += "addStyle(`" + style + "`)\n"
}

func (js *JavascriptDOM) addStyleFunc() {
	js.styleFuncAdded = true
	js.out += `function addStyle(styles) { 
	var css = document.createElement('style'); 
	css.type = 'text/css'; 

	if (css.styleSheet)  
		css.styleSheet.cssText = styles; 
	else
		css.appendChild(document.createTextNode(styles)); 
	
	document.getElementsByTagName("head")[0].appendChild(css); 
}
`
}
