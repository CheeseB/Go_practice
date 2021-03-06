package main

import "fmt"

func fibonacci() func() int {
  fn := 0
  fn1 := 0
  return func() int {
    fn2 := fn + fn1
    fn = fn1
    fn1 = fn2
    if fn1 == 0 {
      fn++
    }
    return fn2
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
