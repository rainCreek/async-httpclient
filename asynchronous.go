package entities

import (
  "fmt"
  "time"
)

func Asynchronous() [10]Infomation {
	start := time.Now()
	customer := GetCustomers()
	destinations := GetDestinations(customer)
	var infos [10]Infomation
	quotes := [10]chan Quoting{}
	weathers := [10]chan Weather{}

	for i, _ := range quotes {
		quotes[i] = make(chan Quoting)
	}

	for i, _ := range weathers {
		weathers[i] = make(chan Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- GetQuoting(d)
		}()

		go func() {
			weathers[i] <- GetWeather(d)
		}()
	}

	for index, dest := range destinations {
		infos[index] = Infomation{W:<-weathers[index],D:dest, Q:<-quotes[index]}
	}

	fmt.Println(time.Since(start))
	return infos
}