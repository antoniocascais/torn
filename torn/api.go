package torn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type selection struct {
	Chain chain `json:"chain"`
}

type chain struct {
	Current  int     `json:"current"`
	Max      int     `json:"max"`
	Timeout  int     `json:"timeout"`
	Modifier float64 `json:"modifier"`
	Cooldown int     `json:"cooldown"`
	Start    int     `json:"start"`
}

func getChainData(client *http.Client, apiKey string) (chain, error) {
	url := fmt.Sprintf("https://api.torn.com/faction/?selections=chain&key=%s", apiKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return chain{}, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return chain{}, err
	}

	s := &selection{}
	err = json.Unmarshal(body, s)
	if err != nil {
		fmt.Println("Error while unmarshalling body")
		return chain{}, err
	}

	return s.Chain, nil
}

func GetChainTimeout(client *http.Client, apiKey string) (int, error) {
	d, err := getChainData(client, apiKey)

	if err != nil {
		return -1, err
	}

	return d.Timeout, nil
}

func GetChainModifier(client *http.Client, apiKey string) (float64, error) {
	d, err := getChainData(client, apiKey)

	if err != nil {
		return -1, err
	}

	return d.Modifier, nil
}
