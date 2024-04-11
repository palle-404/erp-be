package db

type dbLayer struct {
}

type Layer interface {
}

func NewDBLayer() Layer {
	return &dbLayer{}
}
