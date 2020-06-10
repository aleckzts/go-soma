package main
import "fmt"

func main() {
    // var c int
    var c := soma(5, 5)
    fmt.Println("GO: soma de valores")
    fmt.Println("5 + 5 = %d", c)
}

func soma(a int, b int) int {
  // var resultado int
  var resultado := a + b
  return resultado
}
