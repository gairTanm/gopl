package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	response, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("%s", response.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()
		return nil, err
	}
	response.Body.Close()
	return &result, nil
}
