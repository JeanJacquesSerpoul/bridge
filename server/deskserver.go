package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/JeanJacquesSerpoul/bridge/distribution"
	"github.com/JeanJacquesSerpoul/bridge/libdds"
	"github.com/gorilla/mux"
)

// POST Using a lot ...
const POST = "POST"

var version = "undefined"

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(version))
}

func handlerParPbn(w http.ResponseWriter, r *http.Request) {
	pbn, okPbn := r.URL.Query()["pbn"]
	sPbn := checkParams(okPbn, pbn, "")
	vul, okVul := r.URL.Query()["vul"]
	sVul := checkParams(okVul, vul, "NONE")
	s, err := libdds.CallParDDS(sPbn, sVul)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		v := fmt.Sprintf("%v", err)
		_, _ = w.Write([]byte(v))

		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(s))
}

func checkPostParam(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func checkParams(oK bool, value []string, defaultValue string) string {
	if !oK || len(value[0]) < 1 {
		return defaultValue
	}

	return value[0]
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	index, okIndex := r.URL.Query()["value"]
	sIndex := checkParams(okIndex, index, "")
	s := "{\"index\":\"" + sIndex + "\"}"
	board, err := distribution.PbnGenerateFromJSONIndex(s)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		v := fmt.Sprintf("%v", err)
		_, _ = w.Write([]byte(v))

		return
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(board))
}

func handlerMaskMultiPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		sh := new(distribution.Random)
		sDealer := checkPostParam(r.FormValue("dealer"), distribution.Position[0])
		sVulnerable := checkPostParam(r.FormValue("vulnerable"), distribution.Vulnerable[0])
		sMask := checkPostParam(r.FormValue("mask"), distribution.EMPTYDESK)
		sComment := checkPostParam(r.FormValue("comment"), "")
		sCount := checkPostParam(r.FormValue("count"), "1")
		s := "{\"count\":" + sCount + ",\n\"mask\":\"" + sMask +
			"\",\n\"comment\":\"" + sComment + "\",\n\"dealer\":\"" + sDealer +
			"\",\n\"Vulnerable\":\"" + sVulnerable + "\"\n}"
		board, err := distribution.PbnDataGenerateFromJSON(sh, s)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			v := fmt.Sprintf("%v", err)
			_, _ = w.Write([]byte(v))

			return
		}

		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(board))
	}
}

func handlerPointMultiPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		sh := new(distribution.Random)
		sDealer := checkPostParam(r.FormValue("dealer"), distribution.Position[0])
		sVulnerable := checkPostParam(r.FormValue("vulnerable"), distribution.Vulnerable[0])
		sMask := checkPostParam(r.FormValue("mask"), distribution.EMPTYPOINTS)
		sComment := checkPostParam(r.FormValue("comment"), "")
		sCount := checkPostParam(r.FormValue("count"), "1")

		var result string

		if sMask == distribution.EMPTYPOINTS {
			sMask = distribution.EMPTYDESK
			s := "{\"count\":" + sCount + ",\n\"mask\":\"" + sMask +
				"\",\n\"comment\":\"" + sComment + "\",\n\"dealer\":\"" +
				sDealer + "\",\n\"Vulnerable\":\"" + sVulnerable + "\"\n}"
			result, err = distribution.PbnDataGenerateFromJSON(sh, s)
		} else {
			s := "{\"count\":" + sCount + ",\n\"mask\":\"" + sMask +
				"\",\n\"comment\":\"" + sComment + "\",\n\"dealer\":\"" + sDealer +
				"\",\n\"Vulnerable\":\"" + sVulnerable + "\"\n}"
			result, err = distribution.PbnPointDataGenerateFromJSON(sh, s)
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			v := fmt.Sprintf("%v", err)
			_, _ = w.Write([]byte(v))

			return
		}

		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(result))
	}
}

func handlerMaskPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		sh := new(distribution.Random)
		sMask := checkPostParam(r.FormValue("mask"), distribution.EMPTYDESK)
		result, err := distribution.PbnAndIndexGenerateFromMask(sh, nil, sMask)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			v := fmt.Sprintf("%v", err)
			_, _ = w.Write([]byte(v))

			return
		}

		w.WriteHeader(http.StatusOK)

		r.Header.Set("Content-Type", "application/json")

		_, _ = w.Write([]byte(result))
	}
}

func handlerPointPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		var result string

		err := r.ParseForm()

		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		sh := new(distribution.Random)
		sMask := checkPostParam(r.FormValue("mask"), distribution.EMPTYPOINTS)

		if sMask == distribution.EMPTYPOINTS {
			result, err = distribution.PbnAndIndexGenerateFromMask(sh, nil, distribution.EMPTYDESK)
		} else {
			result, err = distribution.GetPbnHandsFromPoints(sh, sMask)
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			v := fmt.Sprintf("%v", err)
			_, _ = w.Write([]byte(v))

			return
		}

		w.WriteHeader(http.StatusOK)

		r.Header.Set("Content-Type", "application/json")

		_, _ = w.Write([]byte(result))
	}
}

func handlerSuitPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		sh := new(distribution.Random)
		sMask := checkPostParam(r.FormValue("mask"), distribution.EMPTYSUIT)
		result, err := distribution.PbnAndIndexGenerateFromSuits(sh, sMask)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			v := fmt.Sprintf("%v", err)
			_, _ = w.Write([]byte(v))

			return
		}

		w.WriteHeader(http.StatusOK)

		r.Header.Set("Content-Type", "application/json")

		_, _ = w.Write([]byte(result))
	}
}

func main() {
	portPtr := flag.String("p", "3000", "API port")
	flag.Parse()

	port := *portPtr
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/maskmultipbn", handlerMaskMultiPbn)
	r.HandleFunc("/api/v1/maskpbn", handlerMaskPbn)
	r.HandleFunc("/api/v1/index", handlerIndex)
	r.HandleFunc("/api/v1/version", handlerVersion)
	r.HandleFunc("/api/v1/suitpbn", handlerSuitPbn)
	r.HandleFunc("/api/v1/pointpbn", handlerPointPbn)
	r.HandleFunc("/api/v1/pointmultipbn", handlerPointMultiPbn)
	r.HandleFunc("/api/v1/parpbn", handlerParPbn)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
