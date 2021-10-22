package main

import (
	"fmt"
  "goxls"
)

func main() {
  xls, err := goxls.Open("testdata/float.xls", "utf-8")
  if err != nil { panic(err) }
  
  sheet := xls.GetSheet(0)
  
  data, ok := sheet.Row(0).GetData(0)

  fmt.Printf("-> %d (%b)\n", data, ok)

  fmt.Println("done.")
}