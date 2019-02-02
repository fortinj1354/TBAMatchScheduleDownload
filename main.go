package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/fortinj1354/TBAMatchScheduleDownload/models"
	"github.com/parnurzeal/gorequest"
)

func main() {
	var apiKey = flag.String("key", "", "API key for The Blue Alliance Read API v3, available at https://www.thebluealliance.com/apidocs")
	var eventCode = flag.String("event", "", "Event code from The Blue Alliance.\nExample: For https://www.thebluealliance.com/event/2019gadal the event code is 2019gadal")
	var filterTeam = flag.String("team", "", "Filter the schedule for a specific team number\nExample: 2974")
	flag.Parse()

	if *apiKey == "" {
		println("API key is required, pass it in with the -key flag")
	} else if *eventCode == "" {
		println("Event code is required, pass it in with the -event flag")
	} else {
		matches := makeTBARequest(*eventCode, *filterTeam, *apiKey)
		if matches != nil {
			fileName := writeToCSV(&matches, *eventCode, *filterTeam)

			println("Results available in " + fileName + ".csv")
		}
	}
}

func makeTBARequest(eventCode string, teamKey string, apiKey string) models.Match {
	var uri string
	if teamKey == "" {
		uri = "/event/" + eventCode + "/matches/simple"
	} else {
		uri = "/team/frc" + teamKey + "/event/" + eventCode + "/matches/simple"
	}

	request := gorequest.New()
	resp, _, err := request.Get("https://www.thebluealliance.com/api/v3"+uri).
		AppendHeader("X-TBA-Auth-Key", apiKey).
		End()
	if err != nil {
		panic(err)
	}

	//Dump body to generic
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var matches models.Match
	decodeErr := json.Unmarshal(bodyBytes, &matches)
	if decodeErr != nil {
		var errorMessage models.Error
		decodeErr = json.Unmarshal(bodyBytes, &errorMessage)

		if decodeErr != nil {
			panic(decodeErr)
		}

		println("Errors encountered:")

		for _, row := range errorMessage.Errors {
			for key, value := range row {
				println(key + ": " + value)
			}
		}
	}

	return matches
}

func writeToCSV(matches *models.Match, eventCode string, teamKey string) string {
	var rows [][]string

	rows = append(rows, []string{
		"Match Number",
		"Time",
		"Blue 1 Team",
		"Blue 1 Scout",
		"Blue 2 Team",
		"Blue 2 Scout",
		"Blue 3 Team",
		"Blue 3 Scout",
		"Red 1 Team",
		"Red 1 Scout",
		"Red 2 Team",
		"Red 2 Scout",
		"Red 3 Team",
		"Red 3 Scout"})

	for _, row := range *matches {
		alliances := row.Alliances
		blueTeams := alliances.Blue.TeamKeys
		redTeams := alliances.Red.TeamKeys

		rows = append(rows, []string{
			row.CompLevel + "-" + strconv.FormatInt(row.MatchNumber, 10),
			time.Unix(row.Time, 0).String(),
			blueTeams[0][3:],
			"",
			blueTeams[1][3:],
			"",
			blueTeams[2][3:],
			"",
			redTeams[0][3:],
			"",
			redTeams[1][3:],
			"",
			redTeams[2][3:],
			""})
	}

	var fileName string
	if teamKey == "" {
		fileName = eventCode
	} else {
		fileName = eventCode + "-" + teamKey
	}

	csvFile, err := os.Create(fileName + ".csv")

	if err != nil {
		panic(err)
	}

	csvWriter := csv.NewWriter(csvFile)

	for _, row := range rows {
		err = csvWriter.Write(row)
		if err != nil {
			panic(err)
		}
	}

	csvWriter.Flush()

	_ = csvFile.Close()

	return fileName
}
