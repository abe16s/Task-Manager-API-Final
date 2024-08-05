package main

import "fmt"

func calculateAverage(grades []float32) float32 {
    if len(grades) == 0 {
        return 0
    }
    var total float32
    for _, grade := range grades {
        total += grade
    }
    return total / float32(len(grades))
}

func main() {
	var studentName string
    fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName)

	var subj int
	fmt.Print("How many subjects have you taken: ")
	fmt.Scanf("%d\n", &subj)

	subjects := make(map[string]float32)
	var s string
	var g float32
	grades := []float32{}

	for i := 0; i < subj; i++ {
		fmt.Printf("Enter subject name #%d: ", i+1)
		fmt.Scanln(&s)
		for {
			fmt.Print("Enter its obtained grade: ")
			fmt.Scanf("%f\n", &g)
			if 0.0 <= g && g <= 100.0 {
				subjects[s] = g
				grades = append(grades, g)
				break
			} else {
				fmt.Println("Invalid grade")
			}
		}

	}

	average := calculateAverage(grades)
	fmt.Println("\nStudent Name:", studentName)
    fmt.Println("Subject Grades:")
    for subj, grd := range subjects {
        fmt.Printf("  %s: %.2f\n", subj, grd)
    }
    fmt.Printf("Average Grade: %.2f\n", average)
}