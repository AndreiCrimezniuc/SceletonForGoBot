package receiver

import (
	"context"
	"encoding/json"
	"net/url"
	"nokogiriwatir/notifierbot/pkg/client"
	"strconv"
)

type UpdatesResponse struct {
	Ok     bool            `json:"ok"`
	Result []client.Update `json:"result"`
}

type Receiver struct {
	client *client.Client
}

func New(client client.Client) *Receiver {
	return &Receiver{
		client: &client,
	}
}

const getUpdatesMethod = "getUpdates"

func (r *Receiver) Updates(offset, limit int) ([]client.Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := r.client.DoRequest(context.TODO(), getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}
