package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/theofficialnar/quiz_game/models"
)

func main() {
	var name string
	score := 0

	showWelcome()

	fmt.Printf("Please enter your name: ")
	fmt.Scan(&name)

	clear()

	var qs = models.GetQuestions()

	for i, v := range qs {
		var answer string

		fmt.Printf("Question %v. %v\n", i+1, v.Question)
		for _, choice := range v.Choices {
			fmt.Println(choice)
		}
		fmt.Printf("\nType the letter of your answer: ")
		fmt.Scan(&answer)

		if strings.ToLower(answer) == v.Answer {
			score++
		}
		clear()
	}

	fmt.Printf("%v, your score is: %v\n", name, score)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showWelcome() {
	fmt.Printf("\n")
	fmt.Println("########################################################")
	fmt.Println("## Welcome to the Linux Quiz! 🔍                      ##")
	fmt.Println("## Time to test just how much you know about Linux 🐧 ##")
	fmt.Println("########################################################")
	fmt.Printf("\n")
}
