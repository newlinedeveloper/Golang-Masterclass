# Golang-Masterclass
Golang Masterclass - Basic and Advanced Concepts

### prerequisite

```
https://go.dev/doc/install

go version

```

##### Create go project 

```
mkdir go-project

cd go-project

go mod init github.com/newlinedeveloper/go-project

touch main.go

```

###### Hello world program

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

```



### Topics


#### Values

Go has various value types including strings, integers, floats, booleans


```
package main
import "fmt"
func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

```

#### Variables

Variables are explicitly declared 


```
package main
import "fmt"
func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)
    
    var e int
    fmt.Println(e)
   

    f := "apple"
    fmt.Println(f)
}

```




