## Fungsi
- > fungsi di go sama dengan fungsi yang lain 
- > type variabel diletakkan setelah nama variabel
- > jika ada 2 variabel dengan nama yang sama bisa dijadikan satu
- > nilai return bisa lebih dari 1
### Contoh
```
package main
import "fmt"
func add(x int, y int) int {
 return x + y
}
func main() {
 fmt.Println(add(423, 13))
}
```

or

```
package main
import "fmt"
func add(x, y int) int {
 return x + y
}
func main() {
 fmt.Println(add(42, 13))
}

```
