package clash

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type GoldPassSeason struct {
	StartTime time.Time
	EndTime   time.Time
}

func (c Client) CurrentGoldPassSeason() (*GoldPassSeason, *ClientError, error) {
	url := "https://api.clashofclans.com/v1/goldpass/seasons/current"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header = c.DefaultHeader()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	switch resp.StatusCode {
	case 200:
		type goldPass struct {
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
		}

		resp := goldPass{}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, nil, err
		}

		startTime, err := time.Parse("20060102T150405.000Z", resp.StartTime)
		if err != nil {
			return nil, nil, err
		}

		endTime, err := time.Parse("20060102T150405.000Z", resp.EndTime)
		if err != nil {
			return nil, nil, err
		}

		return &GoldPassSeason{
			StartTime: startTime,
			EndTime:   endTime,
		}, nil, err
	case 400, 403, 404, 429, 500, 503:
		clientErr := &ClientError{}
		err = json.Unmarshal(body, clientErr)
		return nil, clientErr, err
	}
	return nil, nil, nil
}
