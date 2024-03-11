package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	auth "github.com/dombartenope/viewnotifs/Auth"
	filters "github.com/dombartenope/viewnotifs/Filters"
	"github.com/dombartenope/viewnotifs/Notification"
	utils "github.com/dombartenope/viewnotifs/Utils"
)

func main() {
	// Define a boolean flag
	clearFlag := flag.Bool("clear", false, "Clear the .env file before proceeding")
	flag.Parse()

	// Check the flag value
	if *clearFlag {
		utils.ClearEnvFile() // Clear the .env file if flag is set
	}

	out, err := os.Create("out.csv")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer out.Close()

	writer := csv.NewWriter(out)
	defer writer.Flush()

	//TODO Put behind a clear flag
	defer os.Clearenv()

	client := &http.Client{}
	base_url := "https://api.onesignal.com/notifications?app_id="
	app_id, api_key := auth.CheckForAuth()
	var os int

	headers := []string{
		"DelayedOption",
		"DeliveryTimeOfDay",
		"Errored",
		"Failed",
		"ID",
		"IncludePlayerIds",
		"IncludeExternalUserIds",
		"IncludeAliases",
		"IncludedSegments",
		"SendAfter",
		"CompletedAt",
		"Successful",
		"Received",
		"TemplateId",
		"TimeSentSpending",
	}
	writer.Write(headers)

	param, choiceValue := filters.UserChoice()

	for {
		//Convert int to string for url param
		offset := strconv.Itoa(os)

		url := fmt.Sprintf("%s%s&offset=%s", base_url, app_id, offset)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		apiKey := fmt.Sprintf("Basic %s", api_key)

		req.Header.Add("Authorization", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		var apiResp Notification.ApiResponse
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&apiResp); err != nil {
			log.Fatalf("error decoding API response: %s", err)
		}

		if apiResp.TotalCount == 0 {
			break
		}
		fmt.Println(apiResp.TotalCount)
		fmt.Println(url)

		switch param {
		case "seg":
			filters.FilterBySegment(choiceValue, apiResp, writer)
		case "sid":
			filters.FilterById(choiceValue, apiResp, writer)
		case "eid":
			filters.FilterByEid(choiceValue, apiResp, writer)
		default:
			filters.NoFilter(apiResp, writer)
		}

		//Take offset and add 50 each loop
		os += 50
		resp.Body.Close()
	}

}
