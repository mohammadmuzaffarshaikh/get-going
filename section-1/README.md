# GO Basic

---

## Data Types

### Atomic Data Types

1. String
2. Integer (int, int32, int64, uint, uint32, uint64)
3. Float (float32, folat64)
4. complex - not used much

### Unsafe Data Types

1. Pointers

### Abstract Data Types

1. Array
2. Slices
3. Maps
4. Structs

### Some other Types

1. Interface
2. Functions - functions can also be treated as type in go
3. Channels

---

## Error Handling

### Comma ok/err syntax

```go
input,_ := <some condition>
input,err := <some condition>
_,_ := <some condition>
```

---

## Time in go

```go
	// 3. Time in go
	// follows a wierd syntax
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) // format follows a wierd syntax that must always be like this.

	createDate := time.Date(2025, time.December, 20, 7, 23, 0, 0, time.Local)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
```

---

## Memory Management

Memory management is taken care by go itself. Garbage collection happens automatically.

1. new():

   - Allocate memory but no INIT
   - Zeroed storage

2. make():
   - Allocate memory with INIT
   - Non-zeroed storage

---

## Pointers

    Refrence to memory address just like C, C++.

```go
	fmt.Println("Pointers:")
	num := 24
	var ptr *int = &num
	fmt.Println("Pointer address: ", ptr)
	fmt.Println("Pointer value: ", *ptr)
```

---

## Array

    Array is not used much in go.

```go
	var fruits [3]string // var fruits = [3]string{"apple","banana","mango"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Println(fruits)
	fmt.Println("length: ", len(fruits))
```

## Slices

    Slices are used more than arrays, they just implement array but have more built functions and they do not have a specific length

```go
var fruits = []string{"apple", "mango"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	// adding elements to slice
	fruits = append(fruits, "banana", "peach")
	fmt.Println(fruits)

	// Picking specific elements in range
	fruits = append(fruits[1:2])
	fmt.Println("now fruits: ", fruits)

	// another way to make slice
	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 3263
	highScores[2] = 1623
	highScores[3] = 990

	highScores = append(highScores, 240, 557)

	// sorting the slice
	sort.Ints(highScores)
	fmt.Println(highScores)

	// deleting a specific index element
	var index int = 2
	var courses = []string{"React", "Node", "JS", "GO", "Python", "Rust"}
	fmt.Println("Before removing element: ", courses)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After deleting the value: ", courses)
```
---