package geome

import "math"
import "fmt"

type Point struct {
  X, Y float64
}

type Line struct {
  M, Q float64
}

func Distance(p1, p2 Point) float64 {
  return math.Sqrt((p1.X - p2.X) * (p1.X - p2.X) + (p1.Y - p2.Y) * (p1.Y - p2.Y))
}

func LineThru(p1, p2 Point) (r Line, ok bool) {
  if p1 == p2 || p1.X == p2.X {
    return
  }
  ok = true
  r.M = (p2.Y - p1.Y) / (p2.X - p1.X)
  r.Q = p1.Y - r.M * p1.X
  return
}

func LinePassesThru(r Line, p Point) bool {
  return r.M * p.X + r.Q == p.Y
}



func String(p Point) string {
  return fmt.Sprintf("PUNTO (%.3f,%.3f)", p.X, p.Y)
}

