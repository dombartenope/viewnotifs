package filters

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	utils "github.com/dombartenope/viewnotifs/Utils"
)

func UserChoice() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nWhat would you like to search by?")
	fmt.Println(`
'sid' for subscription id 
'eid' for external id
'seg' for segment
Leave blank for no filter`)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading user choice input: %s", err)
	}

	trimmed := strings.TrimSpace(input)

	switch trimmed {
	//Add in more cases here for a user to type in "template id, eid, etc."

	case "seg":
		userId := searchBySegment()
		return "ex", userId
	case "sid":
		userId := searchBySubId()
		return "sid", userId

	case "eid":
		userId := searchBySubEid()
		return "eid", userId

	default:
		return "", ""
	}

}

func searchBySubId() string {

	userId, exists := os.LookupEnv("SUBSCRIPTION_ID")
	if !exists {
		fmt.Println("Sub ID not found, Please neter a new one : ")

		reader := bufio.NewReader(os.Stdin)
		auth, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		userId = strings.TrimSpace(auth)

		fmt.Printf("New Sub ID set %s\n", userId)

		//save the new AUTH_KEY to .env file

		utils.SaveAuthKeyToFile("SUBSCRIPTION_ID", userId)

	} else {
		fmt.Printf("Sub ID found\n")
	}
	return userId
}

func searchBySubEid() string {

	userId, exists := os.LookupEnv("EXTERNAL_ID")
	if !exists {
		fmt.Println("External ID not found, Please neter a new one : ")

		reader := bufio.NewReader(os.Stdin)
		auth, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		userId = strings.TrimSpace(auth)

		fmt.Printf("New External ID set %s\n", userId)

		//save the new AUTH_KEY to .env file

		utils.SaveAuthKeyToFile("EXTERNAL_ID", userId)

	} else {
		fmt.Printf("External ID found\n")
	}
	return userId
}

func searchBySegment() string {
	userId, exists := os.LookupEnv("SEGMENT_NAME")
	if !exists {
		fmt.Println("Segment name not found, Please neter a new one : ")

		reader := bufio.NewReader(os.Stdin)
		auth, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		userId = strings.TrimSpace(auth)

		fmt.Printf("Segment input set %s\n", userId)

		//save the new AUTH_KEY to .env file

		utils.SaveAuthKeyToFile("SEGMENT_NAME", userId)

	} else {
		fmt.Printf("Segment found\n")
	}
	return userId
}
