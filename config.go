package main

import (
	"fmt"
	"strings"
)

// Config the test config.
type Config struct {
	URL        string
	APIKey     string
	Threads    int
	Iterations int
	Headers    []string
}

func (c Config) String() string {
	return fmt.Sprintf("[%d:%d] - %s [%s]{%s}", c.Threads, c.Iterations, c.URL, c.APIKey, strings.Join(c.Headers, ","))
}
