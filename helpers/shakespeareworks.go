package helper

type ShakespeareWork struct {
	Title     string `json:"Title,omitempty"`
	ReadCount int    `json:"ReadCount,omitempty"`
}

type ResultsFromLebenshtein struct {
	Title     string `json:"Title,omitempty"`
	ReadCount int    `json:"ReadCount,omitempty"`
	Distance  int    `json:"Distance,omitempty"`
}

func Copy(from ShakespeareWork, distance int) ResultsFromLebenshtein {
	var to = ResultsFromLebenshtein{}
	to.Title = from.Title
	to.ReadCount = from.ReadCount
	to.Distance = distance
	return to
}
