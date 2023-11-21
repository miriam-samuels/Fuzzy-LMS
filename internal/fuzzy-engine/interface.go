package fis

type FIS interface {
	fuzzify() [3]float64
	Defuzzify() float64
}
