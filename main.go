package main

import (
	"time"
)

type Spring struct {
	K float64 // Spring Constant
	M float64 // Spring Mass
	X float64 // Offset
	V float64 // Offset
}

func NewSpring(k, m float64) *Spring {
	return &Spring{K: k, M: m, X: 0}
}

func (s *Spring) Step(dt float64) {
	a := -(s.K / s.M) * (s.X)
	s.V += a * dt
	s.X += s.V * dt
}

func (s *Spring) KE() float64 {
	return .5 * s.M * s.V * s.V
}

type SpringSystem struct {
	Springs []*Spring

	J float64 // Linear coupling strength
}

func NewSpringSystem(count int, k, m, j float64) *SpringSystem {
	system := SpringSystem{
		J: j}

	for i := 0; i < count; i++ {
		system.Springs = append(system.Springs, NewSpring(k, m))
	}

	return &system
}

func (system *SpringSystem) KE() float64 {
	out := 0.0
	for _, s := range system.Springs {
		out += s.KE()
	}
	return out
}

func (system *SpringSystem) PE() float64 {
	out := 0.0
	for i, s := range system.Springs {
		prevIndex := (i - 1 + len(system.Springs)) % len(system.Springs)
		nextIndex := (i + 1) % len(system.Springs)

		sPrev := system.Springs[prevIndex]
		sNext := system.Springs[nextIndex]
		out += .5 * s.K * s.X * s.X
		out += .5 * system.J * (sPrev.X - s.X) * (sPrev.X - s.X)
		out += .5 * system.J * (sNext.X - s.X) * (sNext.X - s.X)
	}
	return out
}

func (system *SpringSystem) Step(dt float64) {
	a := make([]float64, len(system.Springs))
	for i, s := range system.Springs {
		prevIndex := (i - 1 + len(system.Springs)) % len(system.Springs)
		nextIndex := (i + 1) % len(system.Springs)

		sPrev := system.Springs[prevIndex]
		sNext := system.Springs[nextIndex]

		fLeft := (sPrev.X - s.X) * system.J
		fRight := (sNext.X - s.X) * system.J
		fSpring := -s.K * s.X

		a[i] = (fLeft + fSpring + fRight) / s.M
	}
	for i, s := range system.Springs {
		s.V += a[i] * dt
		s.X += s.V * dt
	}
}

func main() {
	count := 70
	s1 := NewSpringSystem(count, 1, 1, 10)
	s1.Springs[count/4].X = 1
	s1.Springs[count*3/4].X = 1

	for true {
		s1.Step(1 / 20.)
		DrawSpringSystem(s1, 10)
		time.Sleep(time.Second / 30.0)
	}
}
