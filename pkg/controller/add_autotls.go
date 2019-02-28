package controller

import (
	"github.com/samueltauil/autotls/pkg/controller/autotls"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, autotls.Add)
}
