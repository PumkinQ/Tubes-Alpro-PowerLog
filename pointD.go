package main

import "fmt"

type elektronik struct {
	perangkat string
	watt      int
	durasi    int
}

type daftarElektronik [1000]elektronik

func SelectionSortPerangkat(perangkat string) {
	
	var idx_min int
	var T elektronik

	for i:= 0; i < length(T[].perangkat); i++ {
		idx_min = i - 1
		j := i
		for j < n {
			if [idx_min].perangkat < T[j].perangkat {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}

}

func SelectionSortWatt(watt int) {
	
	for i:= 0; i < length(T[].watt); i++ {
		idx_min = i - 1
		j := i
		for j < n {
			if [idx_min].watt < T[j].watt {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}

}