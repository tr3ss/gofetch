package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getcve(url string) {

	// Get request
	req, err := http.Get(url)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// JSON unmarshalling
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	data := result["data"].([]interface{})

	for _, value := range data {
		cveDetail := value.(map[string]interface{})["cve"]
		description := value.(map[string]interface{})["description"]
		assigner := value.(map[string]interface{})["assigner"]
		severity := value.(map[string]interface{})["severity"]

		fmt.Print("<*>-----------------------------------------<*>\n")
		fmt.Printf("CVE No: %v\n",cveDetail)
		fmt.Printf("Description: %v\n",description)
		fmt.Printf("Severity: %v\n",severity)
		fmt.Printf("Assigner: %v\n\n",assigner)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: gofetch day|week")
		return
	} else if os.Args[1] == "day" {
		url := "https://cvetrends.com/api/cves/24hrs"
		getcve(url)
	} else if os.Args[1] == "week" {
		url := "https://cvetrends.com/api/cves/order-by-tweets-7days"
		getcve(url)
	} else {
		fmt.Println("Usage: gofetch day|week")
		return
	}
}
