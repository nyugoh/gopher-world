## Arrays

### Declaration

We declare arrays using []  and specifying the length as follows

```go
//
var scores[5] float64

```
The above array will contain only zeros ```0```
You can declare and initialize the array immediately as follows

```go
// Define and initialize variables
var scores[5] float64{60, 70, 50, 90, 60}
```
 or you can initialize them later as follows

 ```go
 // Declaration
 var scores[5] float64
// Initialization
 scores[0] = 70
 scores[1] = 60
 scores[2] = 80
 scores[3] = 90
 scores[4] = 100

 ```
### Methods

```go

var scores[5] int
len(scores) // Should return 5

```

### Looping arrays

We use for loops to iterate through items in an array. The following are ways of looping an array.

```go

var scores[5] float64{60, 70, 50, 90, 60}
var total float64 = 0

// len() function returns the length of the array
for i := 0; i < len(scores); i++ {
  total += scores[i]
}

for _, value := range scores {
  // use _ to represent the unused var, tells compiler we don't need this
  total += value
}

```
