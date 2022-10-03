# postfinance-sync

A small CLI to get you bookings in JSON format.

## Disclaimer

This tool is not officially provided nor connected
in any way with PostFinance Ltd. The use of this tool can constitute a breach
of PostFinance's Terms of Service.

### Friendly Disclaimer

Although I haven't found anything related to scraping your own data,
PostFinance might not like you to scrape them like we're doing with this
tool.  
  
Please use the tool responsibly and limit your requests.

## What is this for?

Given a PostFinance account, this tool helps you to
export the transactions in a JSON format, so that you can
then import them in your favorite tool.

The resulting JSON is an array of `Booking`s that can be later used
to perform your analysis.

Here is an anonymized JSON:

```json
{
  "actions": [
    "MovementDetailSimple"
  ],
  "type": 0,
  "bancsUniqueSeqNr": "19c1e114-d8ae-4296-a852-33bae64f99e1",
  "bookingRequestId": "becdca95-4de9-4ebe-9268-0ab360e204bf",
  "date": "2022-10-04",
  "valueDate": "2022-10-03",
  "bookingType": 2,
  "storno": false,
  "debit": {
    "amount": 58.2,
    "currency": 756
  },
  "balance": {
    "amount": 0,
    "currency": 0
  },
  "shorttext": "Acquisto/servizio del 03.10.2022, Lidl 111",
  "text": "ACQUISTO/SERVIZIO DEL 03.10.2022\nCARTA N° XXXX1234\nLIDL 111\nZURICH\nSVIZZERA\n",
  "compiledShortText": "Acquisto/servizio del 03.10.2022, Lidl 111",
  "categorisation": {},
  "location": {
    "address": {
      "city": ""
    },
    "country": 0
  },
  "transactionPartner": {
    "name": ""
  },
  "credit": {
    "amount": 0,
    "currency": 0
  }
}

```

## Requirements

- Go
- Make
- A [PostFinance](https://www.postfinance.ch/) account
- Access to [E-Finance](https://www.postfinance.ch/ap/ba/ob/html/finance/home)

## Building

```bash
make bin
```

## Getting started

1. Connect to your E-Finance account
2. Open the network tab of your browser
3. Retrieve the `MSILAD` cookie value and the `csrfpid` header value

Using these two information, you're ready to start

```bash
read -s -r EF_MSILAD # paste the `MSILAD` cookie value and press enter
export EF_MSILAD

read -s -r EF_CSRFPID # paste the `csrfpid` value and press enter
export EF_CSRFPID 

# Start downloading your bookings!

# --from cannot be older than 2 years from the current date
./bin/postfinance-sync --from "2020-01-01" --to "2022-01-01" --output bookings.json
```