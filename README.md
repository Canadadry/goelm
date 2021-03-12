# GoElm

This is a toy project to bring [Elm](https://elm-lang.org) way of build web app into go. 


It require a web browser to render but does not need to bundle it. You can use [lorca](https://github.com/zserge/lorca)

> A very small library to build modern HTML5 desktop apps in Go. 
> It uses Chrome browser as a UI layer. Unlike Electron it doesn't bundle Chrome into the app package, but rather reuses the one that is already installed. 
> Lorca establishes a connection to the browser window and allows calling Go code from the UI and manipulating UI from Go in a seamless manner.

You can use any library you want if you can provide this interface 

```go 
type Target interface {
	Bind(string, func()) error // to get event from the browser
	Eval(string) error // to execute js on the browser
}
```

## Feature 

 - Pure Go library (no cgo) with a very simple API (only 3 function to write)
 - You don't need to write HTML
 - You don't need to write JS
 - Elm architecture => so immutability in mind
 - All feature of Lorca or any other remote browser library

