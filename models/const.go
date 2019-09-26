package models

type NewsState int

const (
	Pendding NewsState = iota + 1
	Pushlish
	OffLine
)
