package main

import(
	"math"
	"math/rand"
	"time"
)

func BoxMüller(t float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	u1 := rand.Float64()
	u2 := rand.Float64()
	if u1*u2 > 0 {
		x1 := math.Sqrt(-2*math.Log(u1))*math.Cos(2*math.Pi*u2)
		x2 := math.Sqrt(-2*math.Log(u1))*math.Sin(2*math.Pi*u2)
		return math.Sqrt(t)*x1, math.Sqrt(t)*x2
	}
	return 0, 0
}

