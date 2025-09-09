package main

import (
  "fmt"
)

func main() {
  var grade string

  fmt.Println("Enter your grade (A-F):")
  fmt.Scan(&grade)

  switch grade:
  case "A":
    fmt.Println("Great job!")
    break
  case "B":
    fmt.Println("Good job!")
    break
  case "C":
    fmt.Println("Okay job...")
    break
  case "D":
    fmt.Println("You need some work!")
    break
  case "F":
    fmt.Println("You need to re-evaluate life choices")
    break
  default:
    fmt.Println("Invalid.")


// Can declare variable inside the if stmt 
  if success:= true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}
}
