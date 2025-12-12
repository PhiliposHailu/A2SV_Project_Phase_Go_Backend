package main

import "fmt"

type subjectGrade struct {
	name string 
	grade float64
}

func main() {
	var userName string
	var n int
	
	fmt.Print("Enter your name and number of sumbjects (eg. Philip 3): ")
	fmt.Scan(&userName, &n)

	data := map[string]subjectGrade {}
	var totalScore float64
	for i := 0; i < n; i++ {

		// Separation of concerns of the score and the subject name 
		// Subject name 
		var subject string
		var score float64

		fmt.Printf("Enter subject %d name: ", i + 1)
		fmt.Scan(&subject)

		_, ok := data[subject]
		for ok {
			fmt.Printf("Note: subject already exits! \nEnter subject %d name: ", i + 1)
			fmt.Scan(&subject)
			_, ok = data[subject]
		}

		// Subject score 
		fmt.Printf("Enter score for subject %d: ", i + 1)
		fmt.Scan(&score)

		for score < 0 || score > 100 {
			fmt.Printf("Note: score must be with in range of 0 - 100 \nEnter score for subject %d: ", i + 1)
			fmt.Scan(&score)
		}

		totalScore += score
		data[subject] = subjectGrade{
			name: subject,
			grade : score,
		}
	}

	// display the data collected in a formated way 
	fmt.Println("\n-----------------------------------")
	fmt.Printf("%s's Grade Report of %d subjects \n___________________________________\n", userName, n)
	i := 1
	for _, value :=  range data {
		fmt.Printf("|%-10d|%-10s| %-10.2f|\n",i ,value.name ,value.grade )
		i++
	}
	fmt.Printf("|Average Score: %0.2f\n", totalScore / float64(n))
	fmt.Println("-----------------------------------")
}