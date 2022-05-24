package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"mainPlugin/discovery"
	Polling "mainPlugin/pulling"
	"os"
)

func main() {

	recevedARG := os.Args[1]
	jsonDecodedString, err := base64.StdEncoding.DecodeString(recevedARG)
	credentialsMap := make(map[string]interface{})
	err = json.Unmarshal(jsonDecodedString, &credentialsMap)
	if err != nil {
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Json",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)
		log.SetFlags(0)
		log.Fatal(string(b))
	}

	switch credentialsMap["category"] {
	case "discovery":
		Discovery.Discovery(credentialsMap)
	case "pulling":
		Polling.Pulling(credentialsMap)

	default:
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Unknown Category Type",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)

	}

}
