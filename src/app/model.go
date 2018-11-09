package main

// Number holds the Destination to call to
type Number struct {
	Destination string `json:"destination"`
	Language    string `json:"language"`
}

type callRequest struct {
	ToNumber    []string `json:"to"`
	FromNumber  string   `json:"from"`
	NCCOURL     string   `json:"answer_url"`
	CallbackURL string   `json:"event_url"`
}

type vapiResponse struct {
	UUID string `json:"uuid"`
}

type result {
	Text string `json:"text"`
}

type callback struct {
	Speech struct {
		Results []result
	}
}
