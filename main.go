package main

import (
	"encoding/json"
	"github.com/alexflint/go-arg"
	efinance "github.com/denysvitali/postfinance-sync/pkg"
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

	dateFrom := args.ExportFrom
	dateTo := args.ExportTo

	movements, err := c.GetAllMovements(100,
		args.AccountId,
		0,
		[]int{200, 202},
		dateFrom,
		dateTo,
	)

	if err != nil {
		logger.Fatalf("unable to get movements: %v", err)
	}

	// Given the bookings JSON, encode it to file
	enc := json.NewEncoder(f)
	err = enc.Encode(movements)
	if err != nil {
		logger.Fatalf("unable to encode: %v", err)
	}
	_ = f.Close()
}
