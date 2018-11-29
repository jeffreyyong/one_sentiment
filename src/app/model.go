package main

// Number holds the Destination to call to
type Number struct {
	Destination string `json:"destination"`
	Language    string `json:"language"`
}

type vapiNumber struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type callRequest struct {
	ToNumber    []vapiNumber `json:"to"`
	FromNumber  vapiNumber   `json:"from"`
	NCCOURL     []string     `json:"answer_url"`
	CallbackURL []string     `json:"event_url"`
}

type vapiResponse struct {
	UUID string `json:"uuid"`
}

type result struct {
	Text       string `json:"text"`
	Confidence string `json:"confidence"`
}

type speechASR struct {
	Results []result `json:"results"`
}

type callback struct {
	UUID   string    `json:"uuid"`
	Speech speechASR `json:"Speech"`
}

// Result is stores the result
type Result struct {
	Word      string   `json:"word,omitempty"`
	Entities  []Entity `json:"entities ,omitempty"`
	Language  string   `json:"language,omitempty"`
	Sentiment string   `json:"sentiment,omitempty"`
}

// Entity holds entity of the syntax
type Entity struct {
	Type     string  `json:"entity,omitempty"`
	Salience float32 `json:"salience,omitempty"`
}
