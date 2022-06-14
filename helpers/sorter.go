package helper

import "sort"

func Sort(results []ResultsFromLebenshtein) {
	sort.Slice(results, func(i, j int) bool {
		var sortedByFrequency, sortedByDistance, sortedByTitle bool

		sortedByFrequency = results[i].ReadCount > results[j].ReadCount

		if results[i].ReadCount == results[j].ReadCount {
			sortedByDistance = results[i].Distance < results[j].Distance

			if results[i].Distance == results[j].Distance {
				sortedByTitle = results[i].Title < results[j].Title
				return sortedByTitle
			}
			return sortedByDistance
		}
		return sortedByFrequency
	})
}
