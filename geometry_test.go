package geome

import "math"
import "testing"

/*
  go build
  go test
  go test -coverprofile cover.out 
  go tool cover -html=cover.out
*/

func TestDistance(t *testing.T) {
  p1 := Point{0,0}
  p2 := Point{1,1}
  actual := Distance(p1, p2)
  expected := math.Sqrt(2)
  if actual != expected {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestDistance2(t *testing.T) {
  p1 := Point{0,0}
  actual := Distance(p1, p1)
  expected := 0.0
  if actual != expected {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestPassThru(t *testing.T) {
  p0 := Point{0,0}
  p1 := Point{1,1}
  r, ok := LineThru(p0, p1)
  if !ok {
    t.Error("Expected line not to return an error")
  }
  r_expected := Line{1, 0}
  if r != r_expected {
    t.Error("Expected line ", Line{1,0}, " found ", r)
  }
}

func TestPassThru2(t *testing.T) {
  p0 := Point{0,0}
  _, ok := LineThru(p0, p0)
  if ok {
    t.Error("Expected line to return an error")
  }
}

func TestPassThru3(t *testing.T) {
  p0 := Point{0,0}
  p1 := Point{0,1}
  _, ok := LineThru(p0, p1)
  if ok {
    t.Error("Expected line to return an error")
  }
}
