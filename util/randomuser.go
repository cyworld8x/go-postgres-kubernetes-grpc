package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type User struct {
	Email string `json:"email"`
	Name  struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Login struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	} `json:"login"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}

type JsonRandomUser struct {
	Results []User `json:"results"`
	Info    Info   `json:"info"`
}

const BASE_URI_RANDOM_USER_API = "https://randomuser.me/api/"

const (
	MAX_RESULTS = 5000
	MIN_RESULTS = 1
)

const (
	KEY_RESULTS = "results"
	KEY_SEED    = "seed"
)

type QueryConfig struct {
	MaxResults int
	Seed       string
}

func NewQueryConfig() *QueryConfig {
	return &QueryConfig{MIN_RESULTS, ""}
}

func (c *QueryConfig) encode(u *url.URL) {
	q := u.Query()

	if c.MaxResults < 0 {
		c.MaxResults = MIN_RESULTS
	} else if c.MaxResults > MAX_RESULTS {
		c.MaxResults = MAX_RESULTS
	}

	q.Set(KEY_RESULTS, strconv.Itoa(c.MaxResults))

	u.RawQuery = q.Encode()
}

func Generate(c *QueryConfig) ([]User, error) {
	u, err := url.Parse(BASE_URI_RANDOM_USER_API)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c.encode(u)

	response, err := http.Get(BASE_URI_RANDOM_USER_API + "?" + u.Query().Encode())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer response.Body.Close()

	rawBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var jsonRandomUser JsonRandomUser
	err = json.Unmarshal(rawBytes, &jsonRandomUser)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return jsonRandomUser.Results, nil
}
