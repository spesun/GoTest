package github

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
)

import (
	"time"
)
const IssuesURL = "https://api.github.com/search/issues"
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items          []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string, second int) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}


	//根据时间过滤
	var filterResult IssuesSearchResult
	slice := []*Issue{};
	now := time.Now()

	count := 0
	for _, issue := range result.Items {
		stime := issue.CreatedAt
		if now.Unix() - stime.Unix() < int64(second){
			slice = append(slice, issue)
			count++
		}
	}

	filterResult.Items = slice
	filterResult.TotalCount = count

	resp.Body.Close()
	return &result, nil
}
