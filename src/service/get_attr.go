package service

import (
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    "strings"
)

// use wordsapi
//
// details
// https://www.wordsapi.com
//
// mashape
// https://market.mashape.com/wordsapi/wordsapi#word

var (
    StemUrl = "https://wordsapiv1.p.mashape.com/words/"
    MashapeKey string
)

type response struct {
    Results []result `json:"results"`
    Frequency float64 `json:"frequency"`
}

type result struct {
    PartOfSpeech string `json:"partOfSpeech"`
}

func GetAttr(word string) (float64, string) {

    req := createRequest(word)
    res := doRequest(req)
    frequency, partOfSpeech := parseResponse(res)

    return frequency, partOfSpeech
}

func createRequest(word string) *http.Request {
    url := StemUrl + word

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("create request failed : ", err)
    }

    req.Header.Set("Accept", "application/json")
    if MashapeKey != "" {
        req.Header.Set("X-Mashape-Key", MashapeKey)
    }

    return req
}

func doRequest(req *http.Request) []byte {
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Fatal("http request failed : ", err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)

    if res.StatusCode != http.StatusOK {
        log.Fatal("http request failed : ", res.StatusCode, " ", string(body))
    }
    if err != nil {
        log.Fatal("http read response failed : ", err)
    }

    return body
}

func parseResponse(res []byte) (float64, string) {
    var out response
    err := json.Unmarshal(res, &out)
    if err != nil {
        log.Fatal("stem unmarshal failed : ", err)
    }

    frequency := out.Frequency
    partOfSpeech := collectPartOfSpeech(out.Results)

    return frequency, partOfSpeech
}

func collectPartOfSpeech(results []result) string {
    set := map[string]string{}
    for _, result := range results {
        set[result.PartOfSpeech] = result.PartOfSpeech
    }

    parts := make([]string, 0, len(set))
    for k := range set {
        parts = append(parts, k)
    }

    return strings.Join(parts, ",")
}