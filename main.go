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

var works []helper.ShakespeareWork

func showResult(w http.ResponseWriter, r *http.Request) {
	frag := r.URL.Query()["term"][0]
	fmt.Printf("The fragment is: %v\n", frag)

	// dist := helper.Distance(frag[0], "shakespeare")
	// fmt.Printf("Distance is %d", dist)
	results := []helper.ResultsFromLebenshtein{}
	threshold := len(frag) / 2

	for _, work := range works {

		strippedTitle, err := regexp.Compile("[^a-zA-Z ]+")
		if err != nil {
			log.Fatal(err)
		}
		newTitle := strings.ToLower(strippedTitle.ReplaceAllLiteralString(work.Title, ""))

		var dist int
		if len(newTitle) < len(frag) {
			dist = helper.LevenshteinDistance(newTitle, frag)
		} else {
			dist = helper.LevenshteinDistance(newTitle[:len(frag)], frag)
		}

		if dist <= threshold {
			results = append(results, helper.Copy(work, dist))
		}

		helper.Sort(results)
	}

	for _, res := range results {
		fmt.Printf("Results are: %v\n", res.Title)
		w.Write([]byte(res.Title + "\n"))
	}

}

func main() {

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
