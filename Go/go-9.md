## Perulangan 
-> For
```
package main
import "fmt"
func main() {
    for i := 0; i <= 5; i++ {
        fmt.Println("Angka ", i)
    }
}
```

- While
``` 
 package main
import "fmt"
func main() {
    var i = 0
    for i <= 5 {
        fmt.Println("Angka ", i)
        i++
    }
}
```

-> For Condition
```
  package main
import "fmt"
func main() {
    var i = 1
    for {
        fmt.Println("Angka ", i)
        i++
        if i == 5 {
            break
        }
    }
}
```

