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
	Current  int `json:"current"`
	Max      int `json:"max"`
	Timeout  int `json:"timeout"`
	Modifier int `json:"modifier"`
	Cooldown int `json:"cooldown"`
	Start    int `json:"start"`
}

func GetChainCooldown(client *http.Client, apiKey string) (int, error) {
	url := fmt.Sprintf("https://api.torn.com/faction/?selections=chain&key=%s", apiKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	s := &selection{}
	err = json.Unmarshal(body, s)
	if err != nil {
		return -1, err
	}

	return s.Chain.Cooldown, nil
}
