package SNMP

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	g "github.com/gosnmp/gosnmp"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var mapList []map[string]string
var indexList []string

func SnmpPlugin(credentialsMap map[string]interface{}) {

	port, err := strconv.Atoi(credentialsMap["port"].(string)) //port

	snmpCommunity, errCommunity := credentialsMap["community"].(string)
	if !errCommunity {
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Request",
			"status.code": 400,
		}

		marshal, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(marshal)
		fmt.Println(encode)
		os.Exit(0)

	}

	params := &g.GoSNMP{
		Target:    credentialsMap["host"].(string),
		Port:      uint16(port),
		Community: snmpCommunity,
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		//Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}
	err = params.Connect()
	if err != nil {

	}
	defer params.Conn.Close()

	_, oidErr := params.Get([]string{".1.3.6.1.2.1.1.5.0"})

	if oidErr != nil {
		if strings.Contains(oidErr.Error(), "connection refused") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "ERR_CONNECTION_REFUSED | reason: port is not available or Host does not provide SNMP services",
				"status.code": 400,
			}

			marshal, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(marshal)
			fmt.Println(encode)
			os.Exit(0)
		} else if strings.Contains(oidErr.Error(), "request timeout") {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "REQUEST_TIMEOUT | reason: community string is not valid or SNMP version is not supported",
				"status.code": 408,
			}

			marshal, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(marshal)
			fmt.Println(encode)
			os.Exit(0)
		} else {
			log.SetFlags(0)
			log.Fatal(oidErr)
		}

	}

	switch credentialsMap["metric.group"] {
	case "system":
		systemSnmp(params, credentialsMap["host"].(string))
	case "interface":
		interfaceSnmp(params, credentialsMap["host"].(string))
	default:
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Unknown Matrix GROUP",
			"status.code": 400,
		}

		marshal, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(marshal)
		fmt.Println(encode)
		os.Exit(0)

	}
}
