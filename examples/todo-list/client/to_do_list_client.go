package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/client/todos"
)

// The Default to do list HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new to do list HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ToDoList {
	swaggerSpec, err := spec.New(swaggerJSON, "")
	if err != nil {
		// the swagger spec is valid because it was used to generated this code.
		panic(err)
	}
	if formats == nil {
		formats = strfmt.Default
	}
	return New(httptransport.New(swaggerSpec), formats)
}

// New creates a new to do list client
func New(transport client.Transport, formats strfmt.Registry) *ToDoList {
	cli := new(ToDoList)

	cli.Todos = todos.New(transport, formats)

	return cli
}

// ToDoList is a client for to do list
type ToDoList struct {
	Todos *todos.Client
}
