package auth

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	utils "github.com/dombartenope/viewnotifs/Utils"
	"github.com/joho/godotenv"
)

func CheckForAuth() (string, string) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, generating a new one to store your user auth key")
	}

	//Check for the App ID
	appId, exists := os.LookupEnv("APP_ID")
	if !exists {
		fmt.Println("APP_ID not found, Please neter a new APP_ID : ")

		reader := bufio.NewReader(os.Stdin)
		auth, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		appId = strings.TrimSpace(auth)

		fmt.Printf("New App ID set %s\n", appId)

		//save the new AUTH_KEY to .env file
		utils.SaveAuthKeyToFile("APP_ID", appId)

	} else {
		fmt.Printf("APP_ID found\n")
	}
	//Check for the User Auth Key
	authKey, exists := os.LookupEnv("AUTH_KEY")
	if !exists {
		fmt.Println("AUTH_KEY not found, Please neter a new AUTH_KEY : ")

		reader := bufio.NewReader(os.Stdin)
		auth, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		authKey = strings.TrimSpace(auth)

		fmt.Printf("New AUTH_KEY set %s\n", authKey)

		//save the new AUTH_KEY to .env file
		utils.SaveAuthKeyToFile("AUTH_KEY", authKey)

	} else {
		fmt.Printf("AUTH_KEY found\n")
	}

	return appId, authKey
}
