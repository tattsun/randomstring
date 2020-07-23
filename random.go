package randomstring

import "math/rand"

type Generator interface {
	Generate(length uint) string
}

var stdrunes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewAlnumGenerator() *AlnumGenerator {
	return &AlnumGenerator{}
}

type AlnumGenerator struct {
}

func (g *AlnumGenerator) Generate(length uint) string {
	runes := make([]rune, length)

	for i:=uint(0); i<length; i++ {
		runes[i] = rune(stdrunes[rand.Intn(len(stdrunes)-1)])
	}

	return string(runes)
}