### Package for geometry on the Cartesian plane

This package provides a set of types and functions to manipulate points, lines etc. on the Cartesian plane.

To define a `Point` you can do like this:

```
  import "github.com/boldip/geom"
  func main() {
    var p1,p2 geom.Point
    ...
  }
```
A `Point` is a struct with two fields, called `X` and `Y`.
Once you have a `Point` you can do things like computing distance between points:
```
  geom.Distance(p1, p2)
```
