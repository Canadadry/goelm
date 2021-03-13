package common

import (
	"github.com/canadadry/golem"
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

func Run(w, h int, css string, i golem.InitFunc, v golem.ViewFunc, u golem.UpdateFunc) error {
	ui, err := lorca.New("", "", w, h)
	if err != nil {
		return err
	}
	defer ui.Close()

	app := golem.App{
		T: LorcaAsTarget{ui},
	}

	app.JsDom.AddStyle(css)

	err = app.Start(i, v, u)
	if err != nil {
		return err
	}

	<-ui.Done()
	return nil
}
