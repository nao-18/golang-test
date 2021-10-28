package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handlerRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content": "Create post", "author": "nao"}`)
	request, _ := http.NewRequest("POST", "/post/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handlerRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}
