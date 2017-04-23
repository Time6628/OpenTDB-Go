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

func (getter *Getter) GetTrivia(i int) (q Question, err error) {
	stat, body, err := getter.Client.Get(nil, "https://opentdb.com/api.php?amount=" + strconv.Itoa(i))
	if err != nil || stat != 200 {
		return
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
	if err != nil {
		panic(err)
	}
	return
}