package domain

import (
	"fmt"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"os"
)

const (
	queryGetAllUsers = "SELECT id, full_name,email, gender FROM users"
	queryGetUserByID = "SELECT id, full_name,email,gender FROM users WHERE id=$1"
)

type UserRepositoryDB struct {
	db *sqlx.DB
}

func (r UserRepositoryDB) Populate() *errs.AppError {
	csvFile, err := os.Open("MOCK_DATA.csv")
	if err != nil {
		logger.Error("Error while opening file" + err.Error())
	}
	defer csvFile.Close()

	users := []User{}
	if err := gocsv.UnmarshalFile(csvFile, &users); err != nil {
		logger.Error("Error while unmarshalling" + err.Error())
	}
	for _, user := range users {
		insertStmt := `INSERT INTO users (id,email,full_name,gender,password) VALUES ($1, $2, $3, $4, $5)`
		_, err = r.db.Exec(insertStmt, user.ID, user.Email, user.FullName, user.Gender, user.Password)
		if err != nil {
			// Handle individual insertion errors gracefully (e.g., log or skip)
			logger.Error("Error while inserting user" + err.Error())
			fmt.Println("Error inserting user:", err)
		}
	}
	return nil
}

func (r UserRepositoryDB) GetAll() ([]dto.UserResponse, *errs.AppError) {
	var users []dto.UserResponse
	var err error
	err = r.db.Select(&users, queryGetAllUsers)
	if err != nil {
		logger.Error("Error while querying users" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return users, nil
}

func (r UserRepositoryDB) GetByID(id int) (*dto.UserResponse, *errs.AppError) {
	var user dto.UserResponse
	err := r.db.Get(&user, queryGetUserByID, id)
	if err != nil {
		logger.Error("Error while querying user" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &user, nil
}

func NewUserRepositoryDB(dbClient *sqlx.DB) UserRepositoryDB {
	return UserRepositoryDB{dbClient}
}
