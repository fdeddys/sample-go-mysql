package repository

import (
	"context"
	"fmt"

	"com.ddabadi.connection/database"
	"com.ddabadi.connection/model"
)

func GetUser() {

	db := database.GetConnection()
	if db == nil {
		fmt.Println("Error get connection")
	}
	if err := db.Ping(); err != nil {
		panic("PING : " + err.Error())
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Error begin")
	}

	fmt.Printf("Select User")
	rows, err := tx.Query("select ID, CREATED_BY, CREATED_DATE, CREATED_FROM, EMAIL, FULL_NAME from SEC_USER limit 1,10 FOR UPDATE ")
	if err != nil {
		fmt.Println("Error get data !", err.Error())
	}

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.CreatedBy, &user.CreatedDate, &user.CreatedFrom, &user.Email, &user.FullName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// log.Printf(user.Id, user.CreatedBy, user.CreatedDate, user.CreatedFrom, user.Email, user.FullName)
		fmt.Println("Nama ", user.FullName)
	}
	rows.Close()

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error commit")
	}

}
