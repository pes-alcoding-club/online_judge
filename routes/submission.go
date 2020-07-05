package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber"
)

// Submission type
type Submission struct {
	logger *log.Logger
	token  string
}

type token struct {
	Token string `json:"token"`
}

// NewSubmission returns Submission type
func NewSubmission(logger *log.Logger) *Submission {
	return &Submission{logger, ""}
}

// UpdateToken will update the token provided by judge0
func (submission *Submission) UpdateToken(token string) {
	submission.token = token
}

// PostSubmission posts a new submission to judge0 and returns the token
func PostSubmission(c *fiber.Ctx) {
	logger := log.New(os.Stdout, "oj-api-v1-submission", log.LstdFlags)
	newSubmission := NewSubmission(logger)

	token, err := getToken(newSubmission, c.Body())
	if err != nil {
		c.Next(err)
	}

	newSubmission.logger.Println(token)
	result, err := newSubmission.postSubmission()
	if err != nil {
		c.Next(err)
	}

	c.Send(result)
}

func getToken(submission *Submission, payload string) (string, error) {
	url := "https://judge0.p.rapidapi.com/submissions"
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))

	if err != nil {
		submission.logger.Println("Error sending request.")
		return "", err
	}

	req.Header.Add("x-rapidapi-host", "judge0.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		submission.logger.Println("Error receiving response.")
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		submission.logger.Println("Error reading response body.")
		return "", err
	}

	var tempToken token
	err = json.Unmarshal(body, &tempToken)

	if err != nil {
		submission.logger.Println("Error unmarshaling JSON.")
		return "", err
	}

	submission.UpdateToken(tempToken.Token) // put token here
	return tempToken.Token, nil
}

func (submission *Submission) postSubmission() (string, error) {
	url := "https://judge0.p.rapidapi.com/submissions/" + submission.token

	submission.logger.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		submission.logger.Println("Error sending request.")
		return "", err
	}

	req.Header.Add("x-rapidapi-host", "judge0.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		submission.logger.Println("Error receiving response.")
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		submission.logger.Println("Error reading response body.")
		return "", err
	}

	submission.logger.Println(string(body))

	return string(body), nil
}
