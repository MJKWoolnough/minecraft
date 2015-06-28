package generator

import "github.com/MJKWoolnough/minecraft"

type Rand interface {
	Int31n(int32) int32
}

type Generator struct {
	*minecraft.Level
}

func New(l *minecraft.Level) Generator {
	return Generator{l}
}

func NewFromPath(p minecraft.Path) (Generator, error) {
	l, err := minecraft.NewLevel(p)
	if err != nil {
		return Generator{}, err
	}
	return Generator{l}, nil
}
