package controller

import (
	"github.com/mkimuram/egress-mapper/pkg/controller/egressmapper"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, egressmapper.Add)
}
