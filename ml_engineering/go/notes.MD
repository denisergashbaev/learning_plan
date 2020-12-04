# Learning GOlang

```go
import (
    "fmt"
    "math/rand"
)

//sets the seed
rand.Seed(866)

//prints out, adds newline afterwards
fmt.Println("My favorite number is", rand.Intn(10))

//print formatted
fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

// if function parameters share the same type, we can the 
// type from all but the last
func add(x, y int) int {
	return x + y
}

// function can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

a, b := swap("hello", "world")

// there are named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    /*
    this is a 'naked' return
    should be used only in the short functions 
    as it hurts readability
    */
    return
}

/*
var declares a list of variables
notice how the variable type is in the end
*/
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

/* variables with initializers,
type can be omitted,
the variable will take the type of initializer
*/
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

/*
Only inside of functions can we use the := assignment
statement (instead of var with implicit type) 
*/
func main() {
	var i, j int = 1, 2
	c, python, java := true, false, "no!"
}

/*
if statement can start with a short statement
to execute before the condition. The variable declared 
in the if statement is only available in the if-block (ie, inside {})
*/
if v := math.Pow(x, n); v < lim {
   return v
}

// comparing two slices
// source: https://stackoverflow.com/a/15312097/256002
func testEq(a, b []Type) bool {
    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

// Methods //
/* Methods are functions with a special
_receiver_ argument. Receiver appears between
func keyword and the method name (here: (v Vertex))
*/

type Vertex struct {
    X, Y float64
}

// method with a _value_ receiver
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// method with a _pointer_ receiver. 
// (* undoes &)
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// this is a function
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

[What is the meaning of '*' and '&' in Golang?
](https://stackoverflow.com/a/45993495/256002)

```go
a := 5 // a points to value 5
a_copy := a // copy of 5
p := &a // holds memory address of a
*p = 6 // changes val of a=6
```

The loop ``for i := range`` c receives values from the channel repeatedly until it is closed.
