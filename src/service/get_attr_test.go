package service

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
    "io/ioutil"
)

func TestGetAttr(t *testing.T) {
    // set up
    var header string
    handler := func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, testResult)

        request, _ := ioutil.ReadAll(r.Body)
        header = string(request)
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    StemUrl = server.URL
    expectedFrequency := 4.47
    expectedPart := "verb,noun"

    // exercise
    actualFrequency, actualPart := GetAttr("")

    // verify
    if expectedFrequency != actualFrequency {
        t.Errorf("\nexpected: %v\nactual: %v", expectedFrequency, actualFrequency)
    }
    if expectedPart != actualPart {
        t.Errorf("\nexpected: %v\nactual: %v", expectedPart, actualPart)
    }
}

/*
func TestGetAttrError(t *testing.T) {
    // set up
    handler := func (w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Bad Request")
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    StemUrl = server.URL
    input := "test"

    // exercise
    GetAttr(input)

    // verify log.Fatal()
}
*/

const (
    testResult = `{
  "word": "request",
  "results": [
    {
      "definition": "express the need or desire for",
      "partOfSpeech": "verb",
      "synonyms": [
        "ask for",
        "bespeak",
        "call for",
        "quest"
      ],
      "typeOf": [
        "pass",
        "pass along",
        "pass on",
        "put across",
        "communicate"
      ],
      "hasTypes": [
        "invite out",
        "beg",
        "beg off",
        "book",
        "call",
        "call for",
        "challenge",
        "claim",
        "demand",
        "desire",
        "encore",
        "excuse",
        "hold",
        "invite",
        "invoke",
        "lay claim",
        "order",
        "petition",
        "reserve",
        "solicit",
        "supplicate",
        "take out",
        "tap",
        "appeal",
        "apply",
        "arrogate",
        "ask",
        "ask in",
        "ask out",
        "ask over",
        "ask round"
      ],
      "verbGroup": [
        "invite",
        "call for"
      ],
      "examples": [
        "She requested an extra bed in her room"
      ]
    },
    {
      "definition": "the verbal act of requesting",
      "partOfSpeech": "noun",
      "synonyms": [
        "asking"
      ],
      "typeOf": [
        "speech act"
      ],
      "hasTypes": [
        "inquiring",
        "orison",
        "indirect request",
        "entreaty",
        "order",
        "charge",
        "questioning",
        "callback",
        "prayer",
        "appeal",
        "billing",
        "notification",
        "petition",
        "trick or treat",
        "wish",
        "notice",
        "call",
        "invitation",
        "recall"
      ]
    },
    {
      "definition": "a formal message requesting something that is submitted to an authority",
      "partOfSpeech": "noun",
      "synonyms": [
        "petition",
        "postulation"
      ],
      "typeOf": [
        "subject matter",
        "content",
        "message",
        "substance"
      ],
      "hasTypes": [
        "collection",
        "ingathering",
        "solicitation",
        "appeal",
        "application",
        "demand"
      ]
    },
    {
      "definition": "ask (a person) to do something",
      "partOfSpeech": "verb",
      "typeOf": [
        "ask"
      ],
      "hasTypes": [
        "tell",
        "propose",
        "pop the question",
        "enjoin",
        "declare oneself",
        "call",
        "bid",
        "say",
        "order",
        "offer",
        "invite"
      ],
      "examples": [
        "I requested that she type the entire manuscript"
      ]
    },
    {
      "definition": "inquire for (information)",
      "partOfSpeech": "verb",
      "typeOf": [
        "wonder",
        "enquire",
        "inquire"
      ],
      "hasTypes": [
        "seek"
      ],
      "examples": [
        "I requested information from the secretary"
      ]
    }
  ],
  "syllables": {
    "count": 2,
    "list": [
      "re",
      "quest"
    ]
  },
  "pronunciation": {
    "all": "rɪ'kwɛst"
  },
  "frequency": 4.47
}`
)