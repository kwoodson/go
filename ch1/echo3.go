

package main

import (
    "fmt"
    "os"
)

func main() {
  for ind, arg := range os.Args {
      out_str := fmt.Sprintf("%d: %s", ind, arg)
      fmt.Println(out_str)
  }
}
