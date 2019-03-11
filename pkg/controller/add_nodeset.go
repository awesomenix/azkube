package controller

import (
	"github.com/awesomenix/azk/pkg/controller/nodeset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodeset.Add)
}
