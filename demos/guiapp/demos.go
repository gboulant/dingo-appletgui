package main

import (
	"fmt"
	"time"
)

func DEMO00_logscale() error {
	fmt.Println("Executing demo DEMO00_logscale")
	return nil
}

func DEMO01_quintes() error {
	fmt.Println("Executing demo DEMO01_quintes")
	for i := range 10 {
		fmt.Printf("action %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

func DEMO02_vibrato() error {
	fmt.Println("Executing demo DEMO02_vibrato")
	return nil
}

func DEMO03_amplitude_modulation() error {
	fmt.Println("Executing demo DEMO03_amplitude_modulation")
	return nil
}

func DEMO04_frequency_modulation() error {
	fmt.Println("Executing demo DEMO04_frequency_modulation")
	return nil
}

func DEMO05_sounds_like_a_laser() error {
	fmt.Println("Executing demo DEMO05_sounds_like_a_laser")
	return nil
}

func DEMO06_musicalscale() error {
	fmt.Println("Executing demo DEMO06_musicalscale")
	return nil
}
