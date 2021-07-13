package logic

import (
	"fmt"

	"github.com/matt-hoiland/mocking/matching/service"
)

type API struct {
	Service   service.Service
	UserAgent string
	Session   string
}

func (api *API) Create(data string) (int, error) {
	req := service.Request{
		Method: "POST",
		Headers: map[string][]string{
			"UserAgent":  {api.UserAgent},
			"SessionKey": {fmt.Sprintf("s=%s", api.Session)},
		},
		Body: &service.DataBody{
			Data: &data,
		},
	}
	res := api.Service.MakeRequest(&req)
	if res.Error != nil {
		return -1, res.Error
	}
	return *res.Body.ID, nil
}

func (api *API) Retrieve(id int) (string, error) {
	req := service.Request{
		Method: "GET",
		Headers: map[string][]string{
			"UserAgent":  {api.UserAgent},
			"SessionKey": {fmt.Sprintf("s=%s", api.Session)},
		},
		Body: &service.DataBody{
			ID: &id,
		},
	}
	res := api.Service.MakeRequest(&req)
	if res.Error != nil {
		return "", res.Error
	}
	return *res.Body.Data, nil
}

func (api *API) Update(id int, data string) error {
	return nil
}

func (api *API) Delete(id int) error {
	return nil
}
