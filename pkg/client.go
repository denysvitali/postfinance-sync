package efinance

import (
	"encoding/json"
	"fmt"
	"github.com/denysvitali/postfinance-sync/pkg/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Client struct {
	http    *http.Client
	logger  *logrus.Logger
	cookies []http.Cookie
	csrfPid string
}

const BaseUrl = "https://www.postfinance.ch/ap/ba/ef/api/v2/accounts/bookings"
const UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0"

func New(logger *logrus.Logger) (*Client, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}

	c := Client{
		http:   http.DefaultClient,
		logger: logger,
	}

	return &c, nil
}

func (c *Client) SetCookie(cookie *http.Cookie) {
	if c.cookies == nil {
		c.cookies = []http.Cookie{}
	}

	if cookie != nil {
		c.cookies = append(c.cookies, *cookie)
	}
}

func (c *Client) GetBookings(
	accountId string,
	dateFrom string,
	dateTo string,
	forwardId string,
) (models.BookingResult, error) {
	qs := url.Values{
		"productUniqueKey": []string{accountId},
	}

	if dateFrom != "" {
		qs.Add("dateFrom", dateFrom)
	}

	if dateTo != "" {
		qs.Add("dateTo", dateTo)
	}

	if forwardId != "" {
		qs.Add("forwardId", forwardId)
	}

	var result models.Response[models.BookingResult]
	err := Json(c, BaseUrl, qs, &result)
	return result.Result, err
}

func Json[T any](c *Client, baseUrl string, qs url.Values, r *T) error {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	u.RawQuery = qs.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("unable to create request: %v", err)
	}

	c.decorateRequest(req)
	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status %s received", res.Status)
	}

	d := json.NewDecoder(res.Body)
	err = d.Decode(&r)
	if err != nil {
		return fmt.Errorf("unable to decode JSON: %v", err)
	}

	return nil
}

func (c *Client) decorateRequest(req *http.Request) {
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.postfinance.ch/ap/ba/ob/html/finance/assets/movements-overview")
	req.Header.Set("csrfpId", c.csrfPid)
	for _, c := range c.cookies {
		req.AddCookie(&c)
	}
}

func (c *Client) SetMsilad(msilad string) {
	c.SetCookie(&http.Cookie{Name: "MSILAD", Value: msilad})
}

func (c *Client) SetCsrfPid(csrfPid string) {
	c.csrfPid = csrfPid
}
