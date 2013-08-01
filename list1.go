// Geordie Bilkey - 2013
package main

import (
  "fmt"
  "sort"
)

// Finds smallest missing term greater or equal to one
func FindSmallest(list []int) int {
  // Sort
  sort.Sort(sort.IntSlice(list))

  // tempStore
  smallest := 0
  for _,elem := range list {
    if elem == smallest {
      continue
    } else if elem - 1 != smallest {
      break
    }
    smallest = elem
  }

  smallest += 1

  return smallest
}

// More processing efficient through use of hashes O(n)
func FindSmallest_imp(list []int) int {
  m := make(map[int]bool)

  // Insert all elements
  biggest := 0
  for _,elem := range list {
    m[elem] = true
    if elem > biggest {
      biggest = elem
    }
  }

  // Iterate until smallest is found
  i := 1
  for ; i <= biggest; i++ {
    if m[i] == false {
      break
    }
  }

  // Catch continuous
  if i == biggest {
    i += 1
  }

  return i
}

func main() {
  // List some test cases
  tests := [][]int{
    {1,2,3},
    {7},
    {2,3},
    {1,2,3,5,9},
    {1,2,2,3,3,5,9},
    {},
  }

  // Loop through test cases
  for _, test := range tests {
    fmt.Println("Smallest missing from ", test, " is: ", FindSmallest_imp(test))
  }
}
