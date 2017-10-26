## Variabel Go
-> penulisan variabel dengan menambahkan type di belallang variabel 
``` var c, python bool ```
``` var i, java int ```
- jika tidak menuliskan type variabel bisa menggunakan insialisasi langsung
``` var i, j = 1, 2 ```
``` var c, java = true, "no!" ```
_ variabel i dan j memiliki type int, variabel c bertipe boolean dan variabel
java bertipe string _
_ tipe ini diberikan otomatis oleg GO_

- > variabel di dalam fungsi bisa menggunakan short variabel declaration
``` k := 3 ```

or 

``` package main
import "fmt"
func main() {
 var i, j int = 1, 2
 k := 3
 c, python, java := true, false, "no!"
fmt.Println(i, j, k, c, python, java)
}```

_ Hanya bekerja di dalam fungsi tidak bisa untuk di luar fungsi _
