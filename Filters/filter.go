package filters

import (
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dombartenope/viewnotifs/Notification"
	utils "github.com/dombartenope/viewnotifs/Utils"
)

func FilterById(subId string, apiResp Notification.ApiResponse, writer *csv.Writer) {
	fmt.Println("This is running")

	for _, record := range apiResp.Notifications {

		subscriptionIdSearch := subId
		playerIds := strings.Join(record.IncludePlayerIds, ";")

		if strings.Contains(playerIds, subscriptionIdSearch) {

			row := []string{
				utils.StringOrNil(record.DelayedOption),
				utils.StringOrNil(record.DeliveryTimeOfDay),
				strconv.Itoa(record.Errored),
				strconv.Itoa(record.Failed),
				record.ID,
				playerIds,
				strings.Join(record.IncludeExternalUserIds, ";"),
				strings.Join(record.IncludeAliases, ";"),
				strings.Join(record.IncludedSegments, ";"),
				strconv.FormatInt(record.SendAfter, 10),
				strconv.FormatInt(record.CompletedAt, 10),
				strconv.Itoa(record.Successful),
				utils.IntOrNilToString(record.Received),
				utils.StringOrNil(record.TemplateId),
				utils.SecondsToMinSec(record.CompletedAt - record.SendAfter),
			}

			// Write the record's data as a row in the CSV file
			if err := writer.Write(row); err != nil {
				log.Fatalf("error writing record to CSV: %s", err)
			}
		}
	}
}

func FilterByEid(subId string, apiResp Notification.ApiResponse, writer *csv.Writer) {

	for _, record := range apiResp.Notifications {

		subscriptionIdSearch := subId
		externalIds := strings.Join(record.IncludeExternalUserIds, ";")

		if strings.Contains(externalIds, subscriptionIdSearch) {

			row := []string{
				utils.StringOrNil(record.DelayedOption),
				utils.StringOrNil(record.DeliveryTimeOfDay),
				strconv.Itoa(record.Errored),
				strconv.Itoa(record.Failed),
				record.ID,
				strings.Join(record.IncludePlayerIds, ";"),
				externalIds,
				strings.Join(record.IncludeAliases, ";"),
				strings.Join(record.IncludedSegments, ";"),
				strconv.FormatInt(record.SendAfter, 10),
				strconv.FormatInt(record.CompletedAt, 10),
				strconv.Itoa(record.Successful),
				utils.IntOrNilToString(record.Received),
				utils.StringOrNil(record.TemplateId),
				utils.SecondsToMinSec(record.CompletedAt - record.SendAfter),
			}

			// Write the record's data as a row in the CSV file
			if err := writer.Write(row); err != nil {
				log.Fatalf("error writing record to CSV: %s", err)
			}
		}
	}
}

func NoFilter(apiResp Notification.ApiResponse, writer *csv.Writer) {

	for _, record := range apiResp.Notifications {

		row := []string{
			utils.StringOrNil(record.DelayedOption),
			utils.StringOrNil(record.DeliveryTimeOfDay),
			strconv.Itoa(record.Errored),
			strconv.Itoa(record.Failed),
			record.ID,
			strings.Join(record.IncludePlayerIds, ";"),
			strings.Join(record.IncludeExternalUserIds, ";"),
			strings.Join(record.IncludeAliases, ";"),
			strings.Join(record.IncludedSegments, ";"),
			strconv.FormatInt(record.SendAfter, 10),
			strconv.FormatInt(record.CompletedAt, 10),
			strconv.Itoa(record.Successful),
			utils.IntOrNilToString(record.Received),
			utils.StringOrNil(record.TemplateId),
			utils.SecondsToMinSec(record.CompletedAt - record.SendAfter),
		}

		// Write the record's data as a row in the CSV file
		if err := writer.Write(row); err != nil {
			log.Fatalf("error writing record to CSV: %s", err)
		}
	}
}
