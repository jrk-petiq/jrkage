package main

import (
	"fmt"
	"time"
)

func main() {
	var dob string
	layout := "2006-01-02"
	var dobTime time.Time
	var err error

	for {
		fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
		fmt.Scanln(&dob)

		dobTime, err = time.Parse(layout, dob)
		if err != nil {
			fmt.Println("Please enter a valid date in the format YYYY-MM-DD: ")
		} else {
			break
		}
	}

	currentTime := time.Now()
	diff := currentTime.Sub(dobTime)
	age := diff.Hours() / 24 / 365
	fmt.Printf("Your age: %d\n", int(age))

}
