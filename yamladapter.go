package yamladapter

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"gopkg.in/yaml.v3"
)

func init() {
	caddyconfig.RegisterAdapter("yaml", Adapter{})

}

type Adapter struct{}

func (a Adapter) Adapt(body []byte, _ map[string]any) ([]byte, []caddyconfig.Warning, error) {
	m := make(map[string]any)
	err := yaml.Unmarshal(body, &m)
	if err != nil {
		return nil, nil, err
	}

	js, err := json.Marshal(m)
	if err != nil {
		return nil, nil, err
	}

	return []byte(js), nil, err
}
