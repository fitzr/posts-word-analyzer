package service

import (
    "net/http"
    "net/url"
    "strings"
    "log"
    "io/ioutil"
    "encoding/json"
)

// use api to get stem
//
// details
// http://text-processing.com/docs/stem.html
//
// mashape
// https://market.mashape.com/japerk/text-processing#stem

var (
    stemUrl = "https://japerk-text-processing.p.mashape.com/stem/"
    mashapeKey string
)

type stem struct {
    Text string `json:"text"`
}

func GetStem(word string) string {

    req := createRequest(word)
    res := doRequest(req)
    stem := parseResponse(res)

    return stem
}

func createRequest(word string) *http.Request {
    body := url.Values{}
    body.Add("language", "english")
    body.Add("stemmer", "porter")
    body.Add("text", word)

    req, err := http.NewRequest("POST", stemUrl, strings.NewReader(body.Encode()))
    if err != nil {
        log.Fatal("create request failed : ", err)
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("X-Mashape-Key", mashapeKey)

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

func parseResponse(res []byte) string {
    var out stem
    err := json.Unmarshal(res, &out)
    if err != nil {
        log.Fatal("stem unmarshal failed : ", err)
    }

    return out.Text
}