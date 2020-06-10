package main
import "fmt"

func main() {
    c := soma(5, 5)
    fmt.Println("GO: soma de valores")
    fmt.Printf("5 + 5 = %d\n", c)
}

func soma(a int, b int) int {
  var resultado int
  resultado = a + b
  return resultado
}
