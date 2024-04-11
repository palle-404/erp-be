package service

import "github.com/palle-404/erp-be/src/db"

type layer struct {
}

type Layer interface {
}

func NewServiceLayer(dbLayer db.Layer) Layer {
	return layer{}
}
