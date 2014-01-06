package main

import (
  "github.com/NOX73/go-neural"
  "github.com/NOX73/go-neural/lern"
  "github.com/NOX73/go-neural/persist"
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"
)

const (
  jsonFile = "json/descrement.json"
  speed = 0.1
  sampleFile = "json/sample/"
)

type Sample struct {
  In  []float64
  Out []float64
}

func DiscernmentMain () {
  //createNetwork()
  n := loadNetwork()

  testNetwork(n)
  //lernNetwork(n)
  //testNetwork(n)

  saveNetwork(n)
}

func testNetwork (n *neural.Network) {
  //sample := loadSample("plus")
  //sample := loadSample("plus2")
  //sample := loadSample("minus")
  //sample := loadSample("multiple")
  sample := loadSample("multiple2")
  //sample := loadSample("divide")

  res := n.Calculate(sample.In)
  log.Println(res)
}
func lernNetwork (n *neural.Network) {
  samples := make([]*Sample,0,10)
  //samples = append(samples, loadSample("plus"))
  //samples = append(samples, loadSample("minus"))
  samples = append(samples, loadSample("multiple"))
  //samples = append(samples, loadSample("divide"))

  for _, s := range samples {
    for i := 0; i < 1000000; i++ {
      lern.Lern(n, s.In, s.Out, speed)
    }
  }

}

func loadSample (name string) *Sample {

  s := &Sample{}

  fileName := fmt.Sprint(sampleFile, name, ".json")
  b, _ := ioutil.ReadFile(fileName)
  json.Unmarshal([]byte(b), s)

  return s
}


func loadNetwork() *neural.Network {
  return persist.FromFile(jsonFile)
}

func saveNetwork(n *neural.Network) {
  persist.ToFile(jsonFile, n)
}

func createNetwork() {

  n := neural.NewNetwork(9, []int{9,9,4})
  n.RandomizeSynapses()

  persist.ToFile(jsonFile, n)
}