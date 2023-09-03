package efinance

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/h2non/gock.v1"
	"os"
	"testing"
)

func createClient(t *testing.T) *Client {
	var logger = logrus.New()
	c, err := New(logger)
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	c.SetMsilad(os.Getenv("EF_MSILAD"))
	c.SetCsrfPid(os.Getenv("EF_CSRFPID"))
	return c
}

func TestClient_GetMovements(t *testing.T) {
	// Mock the http client
	gock.New(BaseUrl).
		Get(MovementsPath).
		Reply(200).
		File("../test/fixtures/v4/movements.json")

	c := createClient(t)
	res, err := c.GetMovements(
		100,
		os.Getenv("EF_ACCOUNT_ID"),
		0,
		[]int{200, 202},
		"",
		"",
		"",
	)

	if err != nil {
		t.Fatalf("unable to get movements: %v", err)
	}

	if len(res.Result.Movements) != 100 {
		t.Fatalf("expected 100 movements, got %d", len(res.Result.Movements))
	}
}

func TestClient_GetAllMovements(t *testing.T) {
	c := createClient(t)
	c.logger.SetLevel(logrus.DebugLevel)

	// Mock two responses, one with 100 movements and the restartKey and one with 100 movements
	gock.New(BaseUrl).Get(MovementsPath).Reply(200).File("../test/fixtures/v4/movements.json")
	gock.New(BaseUrl).Get(MovementsPath).Reply(200).File("../test/fixtures/v4/movements-2.json")

	res, err := c.GetAllMovements(
		100,
		os.Getenv("EF_ACCOUNT_ID"),
		0,
		[]int{200, 202},
	)
	if err != nil {
		t.Fatalf("unable to get movements: %v", err)
	}

	if len(res) != 101 {
		t.Fatalf("expected 101 movements, got %d", len(res))
	}
}
