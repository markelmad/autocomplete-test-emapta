package main

import (
	helper "emapta-zestyio-golang/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const resultLimit int = 25 //the limit that we should print out
// Value accepts 0 to 1 only.
// This is a multiplier for the threshold that will determine the amount of levenshtein distance that will be accepted.
// 0.75 is the default, meaning, the distance should be less than or equal to the half of the length of the fragment (floor)
// Lower value means it will be stricter. Higher will accept more results but less accurate.
const thresholdMultiplier float32 = 0.75

var works []helper.ShakespeareWork

func showResult(w http.ResponseWriter, r *http.Request) {
	//extract only the fragment from URL query
	if !r.URL.Query().Has("term") {
		w.Write([]byte("No query received."))
		return
	}
	frag := r.URL.Query()["term"][0]

	fmt.Printf("The fragment is: %v\n", frag)

	results := []helper.ResultsFromLebenshtein{}

	for _, work := range works {

		regFrag, err := regexp.Compile("[^a-zA-Z ]+")
		if err != nil {
			log.Fatal(err)
		}
		//remove all non-letters from the fragment
		parsedFrag := strings.ToLower(regFrag.ReplaceAllLiteralString(frag, ""))
		threshold := int(float32(len(parsedFrag)) * thresholdMultiplier)

		if len(parsedFrag) == 0 {
			w.Write([]byte("No query received."))
			return
		}

		//get the distance of the fragment versus each title based on Levenshtein Distance algorithm
		var dist int
		//making sure there's no out of bounds error when the title is shorter than the fragment
		if len(work.Title) < len(parsedFrag) {
			dist = helper.LevenshteinDistance(strings.ToLower(work.Title), parsedFrag)
		} else {
			dist = helper.LevenshteinDistance(strings.ToLower(work.Title)[:len(parsedFrag)], parsedFrag)
		}

		// fmt.Printf("Title: %v -- Frag: %v -- Distance: %v\n", strings.ToLower(work.Title), parsedFrag, dist)
		//this is where we filter. put each title with a distance lower than the threshold to the results splice.
		if dist <= threshold {
			results = append(results, helper.Copy(work, dist))
		}
		//sort the results splice based on the ReadCount (frequency) then the distance, and lastly alphabetically which will deemed necessary.
		helper.Sort(results)
	}
	//making sure we all process when the results is not empty
	if len(results) > 0 {
		//making sure we only show results not more than the resultLimit
		if len(results) > resultLimit {
			results = results[:resultLimit]
		}
		//printing out each results
		for _, res := range results {
			fmt.Printf("Results are: %v\n", res.Title)
			w.Write([]byte(res.Title + "\n"))
		}
	} else {
		w.Write([]byte("No results found based on your query: " + frag))
	}

}

func main() {
	//loading the local text file (JSON) that contains all the necessary data to process
	content, err := ioutil.ReadFile("shakespeare_works.json")
	if err != nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal(content, &works)
	if err2 != nil {
		fmt.Println("Error unmarshalling JSON data: ", err2.Error())
	}

	http.HandleFunc("/autocomplete", showResult)

	if err3 := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err3)
	}
}
