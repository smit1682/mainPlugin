package Discovery

import (
	"encoding/base64"
	"encoding/json"
	g "github.com/gosnmp/gosnmp"
	"log"
	"strconv"
	"strings"
	"time"
)

func snmpDiscovery(credentialsMap map[string]interface{}) {

	port, err := strconv.ParseInt(credentialsMap["port"].(string), 10, 64)
	version := g.Version2c
	if err != nil {
		log.SetFlags(0)
		log.Fatal("port", err)
	}
	if credentialsMap["version"] == "v1" {
		version = g.Version1
	}
	params := &g.GoSNMP{
		Target:    credentialsMap["host"].(string),
		Port:      uint16(port),
		Community: credentialsMap["community"].(string),
		Version:   version,
		Timeout:   time.Duration(3) * time.Second,
	}

	err = params.Connect()

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}

	outputResult, oidErr := params.Get([]string{".1.3.6.1.2.1.1.5.0"})

	if oidErr != nil {
		if strings.Contains(oidErr.Error(), "connection refused") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "ERR_CONNECTION_REFUSED | reason: port is not available or Host does not provide SNMP services",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		} else if strings.Contains(oidErr.Error(), "request timeout") {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "REQUEST_TIMEOUT | reason: community string is not valid or SNMP version is not supported",
				"status.code": 408,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		} else {
			log.SetFlags(0)
			log.Fatal(oidErr)
		}

	} else {
		statusMap := map[string]interface{}{
			"status":       "success",
			"status.code":  200,
			"monitor.name": string(outputResult.Variables[0].Value.([]byte)),
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)
	}

}
