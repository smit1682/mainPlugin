package WINRM

import (
	"github.com/masterzen/winrm"
	"strconv"
	"strings"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
func WinrmPlugin(credentialsMap map[string]interface{}) {

	port, err := strconv.Atoi(credentialsMap["port"].(string)) //port
	if err != nil {

		panic(err)
	}
	host := credentialsMap["host"].(string)
	endpoint := winrm.NewEndpoint(credentialsMap["host"].(string), port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, credentialsMap["username"].(string), credentialsMap["password"].(string))

	if err != nil {
		panic(err)
	}

	switch credentialsMap["metric.group"] {
	case "disk":
		disk_winrm(client, host)
	case "process":
		process_winrm(client, host)
	case "cpu":
		cpu_winrm(client, host)
	case "memory":
		memory_winrm(client, host)
	default:

	}
}
