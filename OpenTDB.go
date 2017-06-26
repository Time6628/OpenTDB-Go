package OpenTDB_Go

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"encoding/json"
	"bytes"
)

func New(client fasthttp.Client) (trivia *Trivia){
	trivia = &Trivia{
		Getter: Getter{Client: client},
	}
	return
}

func (getter *Getter) Request(url string) (body []byte) {
	stat, body, err := getter.Client.Get(nil, url)
	if err != nil || stat != 200 {
		return
	}
	return
}

func (getter *Getter) GetTrivia(i int) (q *Question, err error) {
	q = &Question{}
	body := getter.Request(base + strconv.Itoa(i))
	err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
	if err != nil {
		panic(err)
	}
	return
}

func (getter *Getter) GetTriviaWithCategory(i int, cat int) (q *Question, err error) {
	q = &Question{}
	body := getter.Request(base + strconv.Itoa(i) + with_cat + strconv.Itoa(cat))
	err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
	if err != nil {
		panic(err)
	}
	return
}

func (getter *Getter) GetTriviaWithToken(i int, token string) (q * Question, err error) {
	q = &Question{}
	body := getter.Request(base + strconv.Itoa(i) + with_token + token)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
	if err != nil {
		panic(err)
	}
	return
}

func (getter *Getter) GetToken() (token *Token){
	token = &Token{}
	body := getter.Request(req_token)
	err := json.NewDecoder(bytes.NewReader(body)).Decode(token)
	if err != nil {
		panic(err)
	}
	return
}