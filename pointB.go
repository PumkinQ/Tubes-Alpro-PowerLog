package main

type elektronik struct {
	ID      int
	nama    string
	ruangan string
	watt    float64
	durasi  float64
}

func (d elektronik) konsumsienergi() float64 {
	return (d.watt * d.durasi) / 1000.0
}

var elektroniks []elektronik
var nextID = 1
