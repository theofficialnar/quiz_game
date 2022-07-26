package models

type Question struct {
	Question string
	Choices  []string
	Answer   string
}

var questions = []Question{
	{
		Question: "How to delete a directory in Linux?",
		Choices:  []string{"a. delete", "b. remove", "c. rmdir", "d. ls"},
		Answer:   "c",
	},
	{
		Question: "Which of the following is a text editor that can be used in command mode to edit files on a Linux system?",
		Choices:  []string{"a. lsof", "b. open", "c. edit", "d. vi or vim"},
		Answer:   "d",
	},
	{
		Question: "Which of the following key combinations allows to terminate the current process in Linux shell?",
		Choices:  []string{"a. CTRL+Z", "b. CTRL+C", "c. CTRL+A", "d. CTRL+L"},
		Answer:   "b",
	},
	{
		Question: "Who designed the Linux OS?",
		Choices:  []string{"a. Linus Torvalds", "b. Steve Jobs", "c. Steve Wozniak", "d. Steve Linus"},
		Answer:   "a",
	},
}

func GetQuestions() []Question {
	return questions
}
