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

var lastMessageID int64

func SetLastMessageID(ID int64) { // toDO: Push it in redis or not?
	lastMessageID = ID
}

func GetLastMessageID() int64 {
	return lastMessageID
}

func (r *Receiver) Updates(offset, limit int64) ([]client.Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.FormatInt(offset, 10))
	q.Add("limit", strconv.FormatInt(limit, 10))

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
