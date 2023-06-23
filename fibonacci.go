// This program prints the Fibonacci sequence up to the 10th number.
package main
import "fmt"

func main() {
  fibs := []int{0, 1}
  for i := 2; i <= 10; i++ {
    fibs = append(fibs, fibs[i-1] + fibs[i-2])
  }

  // Print the Fibonacci sequence.
  for _, fib := range fibs {
    fmt.Println(fib)
  }
}
