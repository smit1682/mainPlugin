package Polling

import (
	"encoding/base64"
	"encoding/json"
	"log"
	SNMP "mainPlugin/pulling/snmp"
	SSH "mainPlugin/pulling/ssh"
	WINRM "mainPlugin/pulling/winrm"
)

func Pulling(credentialsMap map[string]interface{}) {

	_, errHost := credentialsMap["host"].(string)
	_, errPort := credentialsMap["port"] //port
	_, errMetricGroup := credentialsMap["metric.group"]
	metricType, errMetricType := credentialsMap["metric.type"]

	if !errHost || !errPort || !errMetricType || !errMetricGroup {
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Request",
			"status.code": 408,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)

	}

	switch metricType {

	case "linux":
		SSH.SshPlugin(credentialsMap)
	case "network.device":
		SNMP.SnmpPlugin(credentialsMap)
	case "windows":
		WINRM.WinrmPlugin(credentialsMap)
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
