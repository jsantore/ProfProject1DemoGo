package main

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
	"log"
	"os"
)

func main() {
	outFile, err := os.Create("Sprint1Output.txt")
	if err != nil {
		log.Fatal("Not able to open the output file: ", err)
	}
	defer outFile.Close()
	for page := 0; page < 5; page++ {
		jobsData := GetData(page)
		saveData(jobsData, outFile)
	}

}

func GetData(page int) []interface{} {
	parameter := map[string]string{
		"api_key":       serpapiKey,
		"engine":        "google_jobs",
		"google_domain": "google.com",
		"q":             "Software Developer",
		"hl":            "en",
		"gl":            "us",
		"location":      "Boston, Massachusetts, United States",
		"lrad":          "100",
		"start":         fmt.Sprintf("%d", page*10),
	}

	search := g.NewGoogleSearch(parameter, serpapiKey)
	results, err := search.GetJSON()
	if err != nil {
		log.Fatal("Error getting data: ", err)
	}
	usefulResults := results["jobs_results"].([]interface{})
	return usefulResults
}

func saveData(data []interface{}, outputFile *os.File) {
	for _, data := range data {
		output := fmt.Sprintf("%s", data)
		outputFile.WriteString(output)
	}
	outputFile.Sync()
}
