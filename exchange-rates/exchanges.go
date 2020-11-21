package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

// define structs for json structure
type Rates struct {
  USDCAD float32 `json:"CAD"`
  USDCHF float32 `json:"CHF"`
  USDEUR float32 `json:"EUR"`
  USDGBP float32 `json:"GBP"`
}

type Currency struct {
  Base string `json:"base"`
  Rates Rates `json:"rates"`
}

func load_rates(config_file string) (Currency) {
  // read file
  data, err := ioutil.ReadFile("./rates.json")
  if err != nil {
    fmt.Print(err)
  }

  var obj Currency
  // unmarshall it
  err = json.Unmarshal(data, &obj)
  if err != nil {
      fmt.Println("error:", err)
  }
  return obj
}

func get_exchange_rate (currency_pair string) (float32) {
    rates := load_rates("./rates.json")
    var rate float32
    switch currency_pair {
  	case "USDCAD":
  		rate = rates.Rates.USDCAD
  	case "USDGBP":
  		rate = rates.Rates.USDGBP
    case "USDCHF":
  		rate = rates.Rates.USDCHF
    case "USDEUR":
  		rate = rates.Rates.USDEUR
    }
   return rate

}

func calc_exchange(from string, to string, amount float32) (float32) {
  return amount * get_exchange_rate(from + to)
}

func main() {
  args := os.Args[1:]
  amount, err := strconv.ParseFloat(args[2], 64)
  if err != nil {
      panic(err)
  }
  exchanged_amount := calc_exchange(args[0], args[1], float32(amount))
  fmt.Println(exchanged_amount)
}
