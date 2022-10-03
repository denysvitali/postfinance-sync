package efinance_test

import (
	efinance "github.com/denysvitali/postfinance-sync/pkg"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"os"
	"testing"
	"time"
)

var logger = logrus.New()

var AccountId = "1234"
var DateFrom = "2020-10-04"
var DateTo = "2022-10-04"

func TestMockedGetBookings(t *testing.T) {
	gock.New(efinance.BaseUrl).Get("").Reply(http.StatusOK).File("./test/fixtures/bookings-1.json")
	c := createClient(t)
	br, err := c.GetBookings(AccountId, DateFrom, DateTo, "")
	if err != nil {
		t.Fatalf("unable to get bookings: %v", err)
	}

	assert.Len(t, br.Bookings, 100)
	assert.Equal(t, br.Bookings[0].CardNumber, "")
	assert.Equal(t, br.Bookings[0].Balance.Amount, 0.0)
	assert.Equal(t, br.Bookings[0].Debit.Amount, 214.3)
	assert.Equal(t, br.Bookings[0].Debit.Currency, 756)
	assert.Equal(t, br.Bookings[0].Categorisation.IndustryId, 0)
	assert.Equal(t, br.Bookings[0].Categorisation.MainCategoryId, 0)
	assert.Equal(t, br.Bookings[0].Date, "2022-10-04")
	assert.Equal(t, br.Bookings[0].ValueDate, "2022-10-03")
	assert.Equal(t, br.Bookings[0].Storno, false)
	assert.Equal(t, br.Bookings[0].Actions, []string{"MovementDetailSimple"})

	gock.New(efinance.BaseUrl).Get("/ap/ba/ef/api/v2/accounts/bookings").MatchParams(
		map[string]string{
			"productUniqueKey": AccountId,
			"dateFrom":         br.SearchParams.MinDate,
			"dateTo":           br.SearchParams.MaxDate,
			"forwardId":        "\\[(\\d+),(\\d+),\"(\\d+)|T|(\\d+)|(\\d+)\"]#(\\d+)",
		},
	).Reply(http.StatusOK).File("./test/fixtures/bookings-2.json")

	br, err = c.GetBookings(AccountId, DateFrom, DateTo, br.ForwardId)
	assert.Nil(t, err)
	assert.Len(t, br.Bookings, 100)
}

func TestGetBookings(t *testing.T) {
	c := createClient(t)

	accountId := os.Getenv("EF_ACCOUNT_ID")
	msilad := os.Getenv("EF_MSILAD")
	csrfPid := os.Getenv("EF_CSRFPID")
	c.SetMsilad(msilad)
	c.SetCsrfPid(csrfPid)

	format := "2006-01-02"
	today := time.Now()
	todayS := today.Format(format)
	twoYearsAgo := today.Add(-time.Hour * 24 * 365 * 2)
	twoYearsAgoS := twoYearsAgo.Format(format)

	br, err := c.GetBookings(accountId, twoYearsAgoS, todayS, "")
	if err != nil {
		t.Fatalf("unable to get bookings: %v", err)
	}

	assert.Len(t, br.Bookings, 100)
}

func createClient(t *testing.T) *efinance.Client {
	c, err := efinance.New(logger)
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}
	return c
}
