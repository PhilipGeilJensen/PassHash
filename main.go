package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Username string
	Password string
}

type UserError struct {
	ErrorMessage string
}

func (e *UserError) Error() string {
	return e.ErrorMessage
}

func main() {
	//Initiate db connection - connection string saved in environment variable
	srv := CreateDbConn(os.Getenv("DB_CONN"))
	//Initiate new reader to listen for input
	reader := bufio.NewReader(os.Stdin)
	//Initiate context
	ctx := context.Background()

	//Create infinite loop for running the program
	for {
		//Print out options
		fmt.Println("1 - Create new user")
		fmt.Println("2 - Login")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("1", text) == 0 {
			//Get the result of the inputs
			user, err := GetCreateUserCredentials()
			if err != nil {
				log.Println(err)
			} else {
				//Create the user
				_, err := CreateUser(ctx, user, srv.DB)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(user.Username + " has been succesfully created")
			}
		} else if strings.Compare("2", text) == 0 {
			for i := 0; i < 5; i++ {
				//Get the result of the inputs
				user, err := GetLoginCredentials()
				if err != nil {
					log.Println(err)
				} else {
					//Login the user
					ok, err := Login(ctx, srv.DB, user)
					if err != nil {
						log.Println(err)
					}
					if ok {
						fmt.Println(user.Username + " has been succesfully logged in")
						break
					}
				}
				fmt.Println("Attempts left: " + strconv.Itoa(5-(i+1)))
			}
		}
	}
}

//GetLoginCredentials listen for user input and create a User struct with the entered username and password
func GetLoginCredentials() (User, error) {
	var user User
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter username")
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	// convert CRLF to LF
	user.Username = strings.Replace(text, "\r\n", "", -1)

	fmt.Println("Enter password")
	fmt.Print("-> ")
	text, err = reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	// convert CRLF to LF
	user.Password = strings.Replace(text, "\r\n", "", -1)

	return user, nil
}

//GetCreateUserCredentials listen for user input and create a User struct with the entered username and password
func GetCreateUserCredentials() (User, error) {
	var user User
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Create a new user")
	fmt.Println("Enter username")
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	// convert CRLF to LF
	user.Username = strings.Replace(text, "\r\n", "", -1)

	fmt.Println("Enter password")
	fmt.Print("-> ")
	text, err = reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	// convert CRLF to LF
	user.Password = strings.Replace(text, "\r\n", "", -1)

	fmt.Println("repeat password")
	fmt.Print("-> ")
	text, err = reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	// convert CRLF to LF
	text = strings.Replace(text, "\r\n", "", -1)

	if text == user.Password {
		return user, nil
	}

	return user, &UserError{ErrorMessage: "Password mismatch"}
}
