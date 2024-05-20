package myip

import (
	_ "embed"
	"encoding/json"
)

//go:embed stun_servers.json
var stunServers []byte

// GetStunServers returns a list of public stun server baked into `myip`.
func GetStunServers() []string {
	servers := make([]string, 0, 5)
	err := json.Unmarshal(stunServers, &servers)
	if err != nil {
		// this means the build has a broken `stun_servers.json` embedded, so we can panic here 😱
		panic(err)
	}
	return servers
}
