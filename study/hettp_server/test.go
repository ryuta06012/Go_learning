/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/11 07:11:56 by hryuuta           #+#    #+#             */
/*   Updated: 2022/01/11 15:45:01 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Ping struct {
	//Status  int    `json:"Status"`
	Fortune string `json:"fortune"`
	Study   string `json:"study"`
	//Result  string `json:"result"`
}

func (p Ping) Init() {
	//p.Status = 0
	p.Fortune = ""
	//p.Result = ""
}

var fortuneList = []string{"Dai-kichi", "Kichi", "Chuu-kichi", "Sho-kichi", "Sue-kichi", "Kyo", "Dai-kyo"}

func selectContentRandom(fortune string) Ping {
	var p Ping
	switch fortune {
	case "Dai-kichi":
		p = Ping{Fortune: fortune, Study: "aaaaaa"}
	case "Kichi":
		p = Ping{Fortune: fortune, Study: "bbbbbb"}
	case "Chuu-kichi":
		p = Ping{Fortune: fortune, Study: "cccccc"}
	case "Sho-kichi":
		p = Ping{Fortune: fortune, Study: "ddddddd"}
	case "Sue-kichi":
		p = Ping{Fortune: fortune, Study: "eeeeeee"}
	case "Kyo":
		p = Ping{Fortune: fortune, Study: "ffffff"}
	case "Dai-kyo":
		p = Ping{Fortune: fortune, Study: "gggggg"}

	}
	return p
}

func (p Ping) selectFortuneRandom() Ping {
	rand.Seed(time.Now().UnixNano())
	content := fortuneList[rand.Int()%len(fortuneList)]

	return selectContentRandom(content)
}

func oracle() Ping {
	var p Ping
	p.Init()
	res := p.selectFortuneRandom()
	return res
}

func pingHandler(w http.ResponseWriter, r *http.Request) {

	ping := oracle()
	//res, err := json.Marshal(ping)

	/* if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} */
	enc := json.NewEncoder(os.Stdout)
	//d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(ping)

	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(ping.Fortune))
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8000", nil)
}
