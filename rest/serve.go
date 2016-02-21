package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IServe interface {
	ServeJson(
		w http.ResponseWriter,
		r *http.Request,
	)
	GetRanks() ([]byte, error)
}

type Serve struct {
	IServe
}

func GetServe() *Serve {
	var serve = Serve{}
	return &serve
}

func (serve *Serve) ServeJson(
	w http.ResponseWriter,
	r *http.Request,
) {

	response, err := serve.GetRanks()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(response))
}

func (serve *Serve) GetRanks() ([]byte, error) {

	rank1 := GenerateRank(
		11,
		1,
	)

	rank2 := GenerateRank(
		12,
		1,
	)

  charts := []Rank{}
	charts = append(
		charts,
		rank1,
		rank2,
	)

	return json.MarshalIndent(charts, "", "  ")
}
