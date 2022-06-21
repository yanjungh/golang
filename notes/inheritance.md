# Inheritance

Go does not support type inheritance but instead has a concept of interface.

## Method vs function

### In appearance:

   * a method has _a receiver_, between the `func` keyword and the method name.
   * the `receiver` is an `object` you would normally pass to a function.

```go
type rectangle struct {
    length, width float32
}

// A regular function
func area2(r rectangle) float32 {
    return r.length * r.width
}

// A method: receiver is a struct.
func (r rectangle) area() float32 {
    return r.length * r.width
}

// How to invoke:
r := rectangle{3, 4}
aByM := r.area()    // A method call
aByF := area2(r)    // A procedural call
```

### In functionality:

A method is an OOP term and operates on an object of a class.  Unlike a Java
class, methods and data (struct in this case) are loosely coupled and defined
separately.
