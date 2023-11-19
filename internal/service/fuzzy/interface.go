package service

type FIS interface {
	Fuzzify() [3]float64
	Defuzzify() float64
}
