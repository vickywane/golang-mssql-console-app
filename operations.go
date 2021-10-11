package main

import "fmt"

func retrieveReminder(alias string) (error, string) {
	sqlStatement := fmt.Sprintf("SELECT title FROM REMINDERS WHERE alias='%v';", alias)
	stmt, err := database.Prepare(sqlStatement); if err != nil {
		fmt.Printf("Query err: %v", err)
	}

	data, queryErr := stmt.Query(); if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}

	for data.Next() {
		var title string
		//var description, alias, title  string
		//var isCompleted int

		nErr := data.Scan(&title); if nErr != nil {
			fmt.Printf("Error: %v", nErr)
		}

		return nil, title
	}

	return nil, ""
}

func createReminder(titleInput, aliasInput, descriptionInput string) error {
	var err error

	err = database.PingContext(dbContext); if err != nil {
		fmt.Printf("Error checking db connection: %v", err)
	}

	queryStatement := fmt.Sprintf("INSERT INTO reminders ( title, description, alias ) VALUES ( '%v', '%v', '%v' );", titleInput, aliasInput, descriptionInput)

	query, err := database.Prepare(queryStatement); if err != nil {
		fmt.Printf("Query err: %v", err)
	}

	_, queryErr := query.QueryContext(dbContext)

	if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}

	return nil
}

func deleteReminder(alias string) error {
	var err error

	err = database.PingContext(dbContext); if err != nil {
		fmt.Printf("Error checking db connection: %v", err)
	}

	queryStatement := fmt.Sprintf("DELETE FROM reminders WHERE alias='%v';", alias)

	stmt, err := database.Prepare(queryStatement); if err != nil {
		fmt.Printf("Query err: %v", err)
	}

	_, queryErr := stmt.Query(); if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}

	return nil
}