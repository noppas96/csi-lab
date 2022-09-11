# Golang Playground

Define varaiables is the same result but different way 

```golang
var foo int = 10

foo := 10
```

## Structures in Golang


Declear pattern

```golang
 type Address struct {
      name string 
      city string
      Pincode int
}
func main() {

    var a = Address{"Muang", "Surathani", 84000} 
    	
		fmt.Println(a)
}
```
Result should be like this
```
{Muang Surathani 84000}
```
