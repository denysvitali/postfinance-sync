package efinance

import (
	"encoding/json"
	"fmt"
	"github.com/denysvitali/postfinance-sync/pkg/models/movement"
	"net/http"
	"net/url"
)

func (c *Client) GetAllMovements(
	numberOfBookings int,
	account string,
	bookingType int,
	movementTypeId []int,
	dateFrom string,
	dateTo string,
) ([]movement.Movement, error) {
	// While GetMovements returns a result with MoreRecords=true, we need to
	// continue to fetch the next page using the restartKey
	var movements []movement.Movement
	var restartKey string

	for {
		res, err := c.GetMovements(
			numberOfBookings,
			account,
			bookingType,
			movementTypeId,
			restartKey,
			dateFrom,
			dateTo,
		)
		if err != nil {
			return nil, err
		}

		c.logger.Debugf("Fetched %d movements", len(res.Result.Movements))
		movements = append(movements, res.Result.Movements...)
		if !res.Result.MoreRecords {
			break
		}
		restartKey = res.Result.RestartKey
	}

	return movements, nil
}

// GetMovements returns the movements for the given account
func (c *Client) GetMovements(
	numberOfBookings int,
	account string,
	bookingType int,
	movementTypeId []int,
	restartKey string,
	bookingDateFrom string,
	bookingDateTo string,
) (*movement.Response, error) {
	u, err := url.Parse(fmt.Sprintf("%s%s", BaseUrl, MovementsPath))
	if err != nil {
		return nil, err
	}

	values := url.Values{
		"numberOfBookings": []string{fmt.Sprintf("%d", numberOfBookings)},
		"accountList":      []string{account},
		"bookingType":      []string{fmt.Sprintf("%d", bookingType)},
		"movementTypeId":   intArrayToString(movementTypeId),
	}

	if restartKey != "" {
		values.Add("restartKey", restartKey)
	}

	if bookingDateFrom != "" {
		values.Add("bookingDateFrom", bookingDateFrom)
	}

	if bookingDateTo != "" {
		values.Add("bookingDateTo", bookingDateTo)
	}

	u.RawQuery = values.Encode()

	var movementsResult movement.Response
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}

	c.decorateRequest(req)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	d := json.NewDecoder(res.Body)
	err = d.Decode(&movementsResult)
	if err != nil {
		return nil, fmt.Errorf("unable to decode JSON: %v", err)
	}

	return &movementsResult, nil
}

func intArrayToString(id []int) []string {
	var res []string
	for _, v := range id {
		res = append(res, fmt.Sprintf("%d", v))
	}
	return res
}
