package Discovery

import (
	"encoding/base64"
	"encoding/json"
	"github.com/masterzen/winrm"
	"log"
	"strconv"
	"strings"
)

//wron user      eyJtZXRyaWMudHlwZSI6IndpbmRvd3MiLCJob3N0IjoiMTcyLjE2LjguMTEzIiwicG9ydCI6IjU5ODUiLCJ1c2VybmFtZSI6InZhbmkiLCJwYXNzd29yZCI6Ik1pbmRAMTIzIn0=
// wrong port    eyJtZXRyaWMudHlwZSI6IndpbmRvd3MiLCJob3N0IjoiMTcyLjE2LjguMTEzIiwicG9ydCI6IjU1IiwidXNlcm5hbWUiOiJkaHZhbmkiLCJwYXNzd29yZCI6Ik1pbmRAMTIzIn0=
//wrong host     eyJtZXRyaWMudHlwZSI6IndpbmRvd3MiLCJob3N0IjoiMTcyLjE2LjguMTE0IiwicG9ydCI6IjU5ODUiLCJ1c2VybmFtZSI6ImRodmFuaSIsInBhc3N3b3JkIjoiTWluZEAxMjMifQ==
//wrong pass     eyJtZXRyaWMudHlwZSI6IndpbmRvd3MiLCJob3N0IjoiMTcyLjE2LjguMTEzIiwicG9ydCI6IjU5ODUiLCJ1c2VybmFtZSI6ImRodmFuaSIsInBhc3N3b3JkIjoiS2luZEAxMjMifQ==
//right         eyJtZXRyaWMudHlwZSI6IndpbmRvd3MiLCJob3N0IjoiMTcyLjE2LjguMTEzIiwicG9ydCI6IjU5ODUiLCJ1c2VybmFtZSI6ImRodmFuaSIsInBhc3N3b3JkIjoiTWluZEAxMjMifQ==

func winRmDiscovery(credentialsMap map[string]interface{}) {

	port, err := strconv.ParseInt(credentialsMap["port"].(string), 10, 64)
	if err != nil {
		panic(err)
	}

	endpoint := winrm.NewEndpoint(credentialsMap["host"].(string), int(port), false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, credentialsMap["username"].(string), credentialsMap["password"].(string))

	if err != nil {
		log.SetFlags(0)
		log.Fatal("hello ", err)
	}

	_, errShell := client.CreateShell()

	if errShell != nil {
		log.SetFlags(0)
		if strings.Contains(errShell.Error(), "invalid content type") {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "UNAUTHORIZED | reason: wrong credentials(Username or Password)",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		} else if strings.Contains(errShell.Error(), "connection refused") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "ERR_CONNECTION_REFUSED | reason: port is not available",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		} else {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       errShell.Error(),
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		}
	} else {

		output, _, _, err := client.RunPSWithString("whoami", "")

		if err != nil {

		}
		statusMap := map[string]interface{}{
			"status":       "success",
			"status.code":  200,
			"monitor.name": strings.TrimRight(output, "\r\n"),
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)
	}

}
