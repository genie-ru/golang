package main
import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Printing all odd numbers %v", findAllOddNumbers(numbers))
}

func findAllOddNumbers(numbers []int) []int {
  result := []int{}
  for _, num := range numbers {
    if num % 2 != 0 {
      result = append(result, num)
    }
  }
  return result
}