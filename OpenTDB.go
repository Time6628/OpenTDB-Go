package OpenTDB_Go

import (
	"github.com/valyala/fasthttp"
	"time"
	"strconv"
	"errors"
	"encoding/json"
	"bytes"
)

var (
	client = fasthttp.Client{ReadTimeout: time.Second * 10, WriteTimeout: time.Second * 10}
)

func New() (trivia *Trivia, err error){
	trivia = &Trivia{
		Getter: Getter{Client: client},
	}
	return
}

func (getter *Getter) GetTrivia(i int) (*Question, error) {
	stat, body, err := client.Get(nil, "https://opentdb.com/api.php?amount=" + strconv.Itoa(i))
	if err != nil || stat != 200 {
		return nil, errors.New("Could not obtain json response")
	}
	q := &Question{}
	json.NewDecoder(bytes.NewReader(body)).Decode(q)
	return q, nil
}