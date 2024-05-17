# Go Lang Overview
## Data Types
### Basic Types
- **bool:** true or false
- **string:** sequence of characters
- **int:** signed integers (int8, int16, int32, int64, int)
- **uint:** unsigned integers (uint8, uint16, uint32, uint64, uint)
- **float:** floating-point numbers (float32, float64)
- **complex:** complex numbers (complex64, complex128)
- **byte:** alias for uint8
- **rune:** alias for int32, represents a Unicode code point
## Composite Types
- **array:** fixed-size sequence of elements, **[n]Type**
- **slice:** dynamic-size sequence of elements, **[]Type**
- **map:** collection of key-value pairs, **map[KeyType]ValueType**
- **struct:** collection of fields, **struct{ Field1 Type1; Field2 Type2 }**
## Pointer Types
**pointer:** holds the memory address of a variable, **`*Type`**
## Function Types
**function:** func(parameterList) (resultList)
## Operators
### Arithmetic Operators
- `+` addition
- `-` subtraction
- `*` multiplication
- `/` division
- `%` modulus
### Relational Operators
- `==` equal to
- `!=` not equal to
- `<` less than
- `<=` less than or equal to
- `>` greater than
- `>=` greater than or equal to
### Logical Operators
- `&&` logical AND
- `||` logical OR
- `!` logical NOT
### Bitwise Operators  
- `&` less than or equal to
- `|` greater than
- `^` greater than or equal to
- `&^` bit clear (AND NOT)
- `<<` left shift
- `>>` right shift
### Assignment Operators
- `=` assign
- `+=` add and assign
- `-=` subtract and assign
- `*=` multiply and assign
- `/=` divide and assign
- `%=` modulus and assign
- `&=` bitwise AND and assign
- `|=` bitwise OR and assign
- `^=` bitwise XOR and assign
- `&^=` bit clear and assign
- `<<=` left shift and assign
- `>>=` right shift and assign
  
## Conditions
### if-Else Statement
```go-lang
if condition {
    // code to execute if condition is true
} else if anotherCondition {
    // code to execute if anotherCondition is true
} else {
    // code to execute if both conditions are false
}
```
## Switch Statement
```go-lang
switch expression {
case value1:
    // code to execute if expression == value1
case value2:
    // code to execute if expression == value2
default:
    // code to execute if no case matches
}
```
## Loops
### For Loop
```go-lang
for initialization; condition; post {
    // code to execute repeatedly
}
```
### While Loop (For Loop as While)
```go-lang
for condition {
    // code to execute while condition is true
}
```
### infinite Loop 
```go-lang
for {
    // code to execute infinitely
}
```

### Range Loop (For Each) 
```go-lang
for index, value := range collection {
    // code to execute for each element
}
```
## Functions
### Basic Function
```go-lang
func functionName(parameterList) returnType {
    // function body
}
```
### Function with Multiple Return Values
```go-lang
func functionName(parameterList) (returnType1, returnType2) {
    // function body
}
```
### Variadic Function
```go-lang
func functionName(param1 type1, param2 ...type2) {
    // function body
}
```
## Struct
### Defining a Struct
```go-lang
type StructName struct {
    Field1 Type1
    Field2 Type2
}
```
### Creating an Instance
```go-lang
var s StructName
s := StructName{Field1: value1, Field2: value2}
```
## Pointers
### Defining a Pointer
```go-lang
var s StructName
s := StructName{Field1: value1, Field2: value2}
```
### Dereferencing a Pointer
```go-lang
value := *p // value at address stored in p
```
## Interfaces 
### Defining an Interface
```go-lang
type InterfaceName interface {
    Method1(param1 type1) returnType1
    Method2(param2 type2) returnType2
}
```
### Implementing an Interface
```go-lang
type StructName struct{}

func (s StructName) Method1(param1 type1) returnType1 {
    // method implementation
}

func (s StructName) Method2(param2 type2) returnType2 {
    // method implementation
}
```
### Using an Interface
```go-lang
var i InterfaceName
i = StructName{}
```
## Error Handling
### Basic Error Handling
```go-lang
import "errors"

func functionName() error {
    if someCondition {
        return errors.New("error message")
    }
    return nil
}
```
### Using Custom Errors
```go-lang
import "fmt"

type CustomError struct {
    Message string
}

func (e *CustomError) Error() string {
    return e.Message
}

func functionName() error {
    if someCondition {
        return &CustomError{Message: "custom error message"}
    }
    return nil
}
```
## Concurrency
### Goroutines
```go-lang
go functionName()
```
### Channels
```go-lang
ch := make(chan Type)

ch <- value // send value to channel
receivedValue := <-ch // receive value from channel
```
### Buffered Channels
```go-lang
ch := make(chan Type, bufferSize)
```
### Select Statement
```go-lang
select {
case msg1 := <-ch1:
    // code to execute if message received from ch1
case msg2 := <-ch2:
    // code to execute if message received from ch2
default:
    // code to execute if no case is ready
}
```


