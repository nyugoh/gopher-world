package main

import "fmt"

func main() {
  mapCoordinates := make(map[string]map[string]float64)

  mapCoordinates["Kagwe"] = map[string]float64{"latitude": 34.4565, "longitude": 0.45322}
  mapCoordinates["Githunguri"] = map[string]float64{"longitude": 0.35667, "latitude": 34.12345}
  mapCoordinates["Cape Town"] = map[string]float64{"longitude": 1.35667, "latitude": 50.12345}

  fmt.Println(mapCoordinates["Kagwe"])
  fmt.Println(mapCoordinates["Kagwe"]["latitude"])
  fmt.Println(mapCoordinates["Kagwe"]["longitude"])

  for key, value := range mapCoordinates {
    fmt.Println("Town :: ", key)
    for _key, _value := range value {
      fmt.Printf("\t%s = %f \n", _key, _value)
    }
  }
}
