package internal

import (
	"encoding/json"
	"net/http"
)

const mdc_url = "https://mdcargas000.gear.host/api.php"

type APIResponse struct {
	Estado string `json:"estado"`
}

func FetchStatus(tipo string, suc string, numero string) (APIResponse, error) {
	response := APIResponse{}
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, mdc_url, nil)
	q := r.URL.Query()
	q.Add("tipo", tipo)
	q.Add("suc", suc)
	q.Add("numero", numero)
	r.URL.RawQuery = q.Encode()

	rsp, err := client.Do(r)
	if err != nil {
		return response, err
	}
	err = json.NewDecoder(rsp.Body).Decode(&response)
	return response, err
}
