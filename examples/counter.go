package main

import (
	"fmt"
	"github.com/canadadry/goelm"
	"github.com/zserge/lorca"
)

type LorcaAsTarget struct {
	l lorca.UI
}

func (ui LorcaAsTarget) Bind(name string, fn func()) error {
	return ui.l.Bind(name, fn)
}
func (ui LorcaAsTarget) Eval(js string) error {
	v := ui.l.Eval(js)
	return v.Err()
}

func Init() interface{} {
	return int(0)
}

func View(model interface{}) goelm.Element {
	return goelm.Div(goelm.Props{Class: "counter-container"},
		goelm.Div(goelm.Props{Class: "counter", Id: "counter", InnerText: fmt.Sprintf("%v", model)}),
		goelm.Div(goelm.Props{Class: "btn-row"},
			goelm.Div(goelm.Props{
				Class:     "btn btn-incr",
				InnerText: "+1",
				EventListener: map[string]string{
					"click": "Plus",
				},
			}),
			goelm.Div(goelm.Props{
				Class:     "btn btn-decr",
				InnerText: "-1",
				EventListener: map[string]string{
					"click": "Minus",
				},
			}),
		),
	)
}

func Update(m interface{}, event string) (interface{}, string) {
	c, ok := m.(int)
	if !ok {
		c, ok = Init().(int)
		if !ok {
			panic("init func should return an int")
		}
	}

	switch event {
	case "Plus":
		c += 1
	case "Minus":
		c -= 1
	}
	return c, ""
}

func run() error {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		return err
	}
	defer ui.Close()

	app := goelm.App{
		T: LorcaAsTarget{ui},
	}

	app.JsDom.AddStyle(`
		* { margin: 0; padding: 0; box-sizing: border-box; user-select: none; }
		body { height: 100vh; display: flex; align-items: center; justify-content: center; background-color: #f1c40f; font-family: 'Helvetika Neue', Arial, sans-serif; font-size: 28px; }
		.counter-container { display: flex; flex-direction: column; align-items: center; }
		.counter { text-transform: uppercase; color: #fff; font-weight: bold; font-size: 3rem; }
		.btn-row { display: flex; align-items: center; margin: 1rem; }
		.btn { cursor: pointer; min-width: 4em; padding: 1em; border-radius: 5px; text-align: center; margin: 0 1rem; box-shadow: 0 6px #8b5e00; color: white; background-color: #E4B702; position: relative; font-weight: bold; }
		.btn:hover { box-shadow: 0 4px #8b5e00; top: 2px; }
		.btn:active{ box-shadow: 0 1px #8b5e00; top: 5px; }
	`)

	err = app.Start(Init, View, Update)
	if err != nil {
		return err
	}

	<-ui.Done()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
