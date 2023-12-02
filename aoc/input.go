package aoc

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func FetchInput(token string, n int) (io.Reader, error) {
	req, err := http.NewRequest(http.MethodGet, "https://adventofcode.com/2023/day/"+strconv.Itoa(n)+"/input", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: token,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}

	return res.Body, nil
}
