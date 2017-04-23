package OpenTDB_Go

import "github.com/valyala/fasthttp"

type Trivia struct {
	Getter Getter
}

type Getter struct {
	Client fasthttp.Client
}

type Question struct {
	Response int `json:"response_code"`
	Results Results `json:"results"`
}

type Results struct {
	Category string `json:"category"`
	Type string `json:"type"`
	Difficulty string `json:"difficulty"`
	Question string `json:"question"`
	CorrectAnswer string `json:"correct_answers"`
	IncorrectAnswer []string `json:"incorrect_answers"`
}
