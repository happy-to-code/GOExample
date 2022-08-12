package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Serve a controller based on the root Router, "/".
	mvc.New(app).Handle(new(ExampleController))

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhost:8080/custom_path
	app.Run(iris.Addr(":8080"))
}

// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

// BeforeActivation called once, before the controller adapted to the main application
// and of course before the server ran.
// After version 9 you can also add custom routes for a specific controller's methods.
// Here you can register custom method's handlers
// use the standard router with `ca.Router` to
// do something that you can do without mvc as well,
// and add dependencies that will be binded to
// a controller's fields or method function's input arguments.
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}

	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide",
		anyMiddlewareHere,
	)

	// or even add a global middleware based on this controller's router,
	// which in this example is the root "/":
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}
