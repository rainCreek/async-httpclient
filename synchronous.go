package entities

import (
  "fmt"
  "time"
)

func Synchronous() [10]Infomation {
  start := time.Now()
  customer := GetCustomers()
  destinations := GetDestinations(customer)
  var infos [10]Infomation
  for index, dest := range destinations {
    q := GetQuoting(dest)
    w := GetWeather(dest)
    infos[index] = Infomation{W : w,D : dest,Q : q}
  }
  fmt.Println(time.Since(start))
  return infos
}



