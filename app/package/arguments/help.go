package arguments

import "fmt"

func PrintHelpMessage() {
	helpMessage := "Need help? Check the usecases!"
	optionAdd := "add, add a movie"
	optionList := "list, list all movies"
	optionDetails := "details, Details of a movie"
	optionDelete := "delete, delete a movie"
	fmt.Println(helpMessage, optionAdd, optionList, optionDetails, optionDelete)
}
