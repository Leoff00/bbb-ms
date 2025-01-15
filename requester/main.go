package main

import (
	"encoding/json"
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	type V struct {
		Voto int `json:"voto"`
	}
	v := V{Voto: 1}

	b, _ := json.Marshal(v)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:3001/api/voto",
		Body:   b,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	})

	rate := vegeta.Rate{Freq: 20, Per: time.Second}
	duration := 20 * time.Second
	attacker := vegeta.NewAttacker()
	results := attacker.Attack(targeter, rate, duration, "Load Test")

	for res := range results {
		fmt.Printf("Response Code: %d | Latency: %s | Error: %s | Message: %s\n", res.Code, res.Latency, res.Error, res.Body)
	}
}
