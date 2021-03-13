package golem

import (
	"fmt"
	"sync"
)

type Target interface {
	Bind(string, func()) error
	Eval(string) error
}

type App struct {
	T         Target
	JsDom     JavascriptDOM
	debugCmpt bool
	model     interface{}
	sync.Mutex
}

func (a *App) Render(e Element) (map[string]Object, error) {
	refs, err := Render(a.JsDom.GetBody(), e, &a.JsDom)
	if err != nil {
		return nil, err
	}
	out, err := a.JsDom.Apply(a.T)
	if a.debugCmpt {
		fmt.Println(out)
	}
	return refs, err
}

type InitFunc func() interface{}
type ViewFunc func(model interface{}) Element
type UpdateFunc func(m interface{}, e string) (interface{}, string)

func (a *App) Start(i InitFunc, v ViewFunc, u UpdateFunc) error {
	a.model = i()
	vdom := v(a.model)

	a.JsDom.EventFunc = a.eventRouter(u, v)
	_, err := a.Render(vdom)
	return err
}

func (a *App) eventRouter(u UpdateFunc, v ViewFunc) func(string) func() {
	return func(name string) func() {
		return func() {
			a.Lock()
			defer a.Unlock()

			current := name
			for current != "" {
				m, event := u(a.model, current)
				a.model = m
				if current == event {
					break
				}
				current = event
			}
			a.JsDom.Raw("document.body.innerHTML='';\n")
			vdom := v(a.model)
			a.Render(vdom)
		}
	}
}
