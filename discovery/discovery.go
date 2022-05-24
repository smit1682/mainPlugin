package Discovery

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

func Discovery(credentialsMap map[string]interface{}) {

	switch credentialsMap["metric.type"] {
	case "linux":
		sshDiscovery(credentialsMap)
	case "windows":
		winRmDiscovery(credentialsMap)
	case "network.device":
		snmpDiscovery(credentialsMap)
	default:
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Unknown Matrix Type",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)

	}

}
