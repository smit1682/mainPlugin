package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"log"
	"os"
	"strconv"
	"strings"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
func WinrmPlugin(credentialsMap map[string]interface{}) {

	port, err := strconv.Atoi(credentialsMap["port"].(string)) //port

	winrmUser, errUsername := credentialsMap["username"].(string)
	winrmPassword, errPassword := credentialsMap["password"].(string)
	if !errUsername || !errPassword {
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Request",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		fmt.Println(encode)
		os.Exit(0)

	}

	host := credentialsMap["host"].(string)
	endpoint := winrm.NewEndpoint(credentialsMap["host"].(string), port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, winrmUser, winrmPassword)

	if err != nil {
		panic(err)
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
			fmt.Println(encode)
			os.Exit(0)
		} else if strings.Contains(errShell.Error(), "connection refused") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "ERR_CONNECTION_REFUSED | reason: port is not available",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			fmt.Println(encode)
			os.Exit(0)
		} else {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       errShell.Error(),
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			fmt.Println(encode)
			os.Exit(0)
		}
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
	case "system":
		system_winrm(client, host)
	default:
		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Unknown Matrix Group",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		fmt.Println(encode)
		os.Exit(0)

	}
}
