package main

import (
	"encoding/json"
	"github.com/alexflint/go-arg"
	efinance "github.com/denysvitali/postfinance-sync/pkg"
	"github.com/denysvitali/postfinance-sync/pkg/models"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	Debug      *bool  `arg:"-D,--debug"`
	ExportFrom string `arg:"-f,--from,required"`
	ExportTo   string `arg:"-t,--to,required"`
	OutputFile string `arg:"-o,--output,required"`
	Msilad     string `arg:"env:EF_MSILAD,required"`
	CsrfPid    string `arg:"env:EF_CSRFPID,required"`
	AccountId  string `arg:"-a,--account-id,env:EF_ACCOUNT_ID,required"`
}

var logger = logrus.New()

func main() {
	arg.MustParse(&args)
	if args.Debug != nil && *args.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	c, err := efinance.New(logger)
	if err != nil {
		logger.Fatalf("unable to create client: %v", err)
	}

	c.SetMsilad(args.Msilad)
	c.SetCsrfPid(args.CsrfPid)

	f, err := os.OpenFile(args.OutputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		logger.Fatalf("unable to open file: %v", err)
	}

	var bookings []models.Booking
	dateFrom := args.ExportFrom
	dateTo := args.ExportTo
	forwardId := ""
	page := 1
	for {
		logger.Debugf("page %d, forwardId=%s", page, forwardId)
		bookingsResult, err := c.GetBookings(args.AccountId, dateFrom, dateTo, forwardId)
		if err != nil {
			logger.Errorf("unable to get bookings: %v", err)
			break
		}
		bookings = append(bookings, bookingsResult.Bookings...)
		forwardId = bookingsResult.ForwardId
		if forwardId == "" {
			logger.Infof("reached the end")
			break
		}
		page++
	}

	// Given the bookings JSON, encode it to file
	enc := json.NewEncoder(f)
	err = enc.Encode(bookings)
	if err != nil {
		logger.Fatalf("unable to encode: %v", err)
	}
	_ = f.Close()
}
