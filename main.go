package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
)

var connectionError error

func main() {
	var err error

	database, connectionError = sql.Open("mssql", connectionString); if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}

	err = database.PingContext(dbContext); if err != nil {
		fmt.Printf("Error checking db connection: %v", err)
	}

	fmt.Println(err)

	fmt.Println("-> Welcome to Reminders Console App, built using Golang and Microsoft SQL Server")
	fmt.Println("-> Select a numeric option; \n [1] Create a new Reminder \n [2] Get a reminder \n [3] Delete a reminder")

	consoleReader := bufio.NewScanner(os.Stdin)
	consoleReader.Scan()
	userChoice := consoleReader.Text()

	switch userChoice {
	case "1":
		var (
			titleInput,
			descriptionInput,
			aliasInput string
		)
		fmt.Println("You are about to create a new reminder. Please provide the following details:")

		fmt.Println("-> What is the title of your reminder?")
		consoleReader.Scan()
		titleInput = consoleReader.Text()

		fmt.Println("-> What is the description of your reminder?")
		consoleReader.Scan()
		descriptionInput = consoleReader.Text()

		fmt.Println("-> What is an alias of your reminder? [ An alias will be used to retrieve your reminder ]")
		consoleReader.Scan()
		aliasInput = consoleReader.Text()

		err := createReminder(titleInput, descriptionInput, aliasInput); if err != nil {
			return
		}
		break

	case "2":
		fmt.Println("-> Please provide an alias for your reminder:")
		consoleReader.Scan()
		aliasInput := consoleReader.Text()

		getErr, data := retrieveReminder(aliasInput); if getErr != nil {
			fmt.Println(getErr)
	  }

	  fmt.Println(data)
		break

	case "3":
		fmt.Println("-> Please provide the alias for the reminder you want to delete:")

		consoleReader.Scan()
		deleteAlias := consoleReader.Text()

		getErr := deleteReminder(deleteAlias); if getErr != nil {
			fmt.Println(getErr)
		}
		break

		default:
		 fmt.Printf("-> Option: %v is not a valid numeric option. Try 1 , 2 , 3", userChoice)
	}
}
