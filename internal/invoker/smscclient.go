package invoker

import (
	"github.com/go-resty/resty/v2"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"os"
)

type SmscClient struct {
	client *resty.Client
	url    string
}

func NewSmscClient() *SmscClient {
	return &SmscClient{
		url:    os.Getenv("SMSC_CLIENT_URL"),
		client: resty.New(),
	}
}

func (c *SmscClient) SubmitSms(request *presentation.SmsRequest) error {
	_, err := c.client.R().SetBody(request).Post(c.url + "/submissions")
	return err
}
