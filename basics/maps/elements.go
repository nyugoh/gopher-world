package main

import "fmt"

func main() {
  elements := make(map[string]string)

  elements["He"] = "Hellium"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["Ag"] = "Argon"
  elements["Hg"] = "Mercury"

  for key, value := range elements {
    fmt.Printf("Chemical sysmbol for %s is %s \n", value, key)
  }

}
