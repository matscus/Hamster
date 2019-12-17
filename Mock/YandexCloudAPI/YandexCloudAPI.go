package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var (
	listenport string
	mode       string
	pemPath    string
	keyPath    string
	scoreValue = []float64{0.7622091841744798, 0.4661361235810615, 0.5694269819653356}
)

type IntRange struct {
	min, max int
}

type ScoresRequest struct {
	ResponseIds struct {
		RequestID string `json:"request_id"`
	} `json:"response_ids"`
	UserIds struct {
		Cookies []struct {
			Cookie       string `json:"cookie"`
			CookieVendor string `json:"cookie_vendor"`
		} `json:"cookies"`
		Emails []struct {
			IDValue string `json:"id_value"`
		} `json:"emails"`
		Phones []struct {
			IDValue string `json:"id_value"`
		} `json:"phones"`
	} `json:"user_ids"`
	Scores []struct {
		ScoreName string `json:"score_name"`
	} `json:"scores"`
}

type ScoresResponse struct {
	ResponseIds struct {
		RequestID string `json:"request_id"`
	} `json:"response_ids"`
	Scores []Scores `json:"scores"`
}
type Scores struct {
	ScoreName  string  `json:"score_name"`
	ScoreValue float64 `json:"score_value,omitempty"`
	HasScore   bool    `json:"has_score"`
}
type Errors struct {
	Code              string `json:"code"`
	InternalRequestID string `json:"internal_request_id"`
}

func scores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ScoresRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(http.StatusUnauthorized)
		errors := Errors{Code: "score_read_forbidden", InternalRequestID: req.ResponseIds.RequestID}
		err = json.NewEncoder(w).Encode(errors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
		}
		return
	}
	res := ScoresResponse{}
	res.ResponseIds = req.ResponseIds
	l := len(req.Scores)
	res.Scores = make([]Scores, l, l)
	for i := 0; i < l; i++ {
		res.Scores[i].ScoreName = req.Scores[i].ScoreName
		res.Scores[i].HasScore = true
		res.Scores[i].ScoreValue = scoreValue[i]
	}
	rand := rand.New(rand.NewSource(55))
	ir := IntRange{20, 100}
	time.Sleep(time.Duration(ir.NextRandom(rand)) * time.Millisecond)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}

}
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func main() {
	flag.StringVar(&pemPath, "pempath", "./server.pem", "server sert")
	flag.StringVar(&keyPath, "keypath", "./server.key", "server sert")
	flag.StringVar(&mode, "mode", "https", "server mode")
	flag.StringVar(&listenport, "port", ":9443", "port to Listen")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/v1/accounts/partner_a/scores", scores).Methods("POST")
	http.Handle("/v1/accounts/", r)
	switch mode {
	case "http":
		log.Printf("Listen to http on port %s", listenport)
		log.Fatal(http.ListenAndServe(listenport, r))
	case "https":
		log.Printf("Listen to https on port %s", listenport)
		log.Fatal(http.ListenAndServeTLS(listenport, pemPath, keyPath, context.ClearHandler(http.DefaultServeMux)))
	}
}
