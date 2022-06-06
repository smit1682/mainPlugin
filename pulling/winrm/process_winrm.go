package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
	"strings"
)

func process_winrm(client *winrm.Client, host string) {
	commandForProces := "get-process"
	process, _, _, err := client.RunPSWithString(commandForProces, "")

	if err != nil {
	}

	var processList []map[string]string
	processStringArray := strings.Split(process, "\n")

	flag := 1
	for _, v := range processStringArray {
		if flag <= 3 {
			flag++
			continue
		}
		processEachWorld := strings.SplitN(standardizeSpaces(v), " ", 8)
		if len(processEachWorld) <= 7 {
			break
		}

		temp := map[string]string{
			"process":        processEachWorld[7],
			"id":             processEachWorld[6],
			"cpu":            processEachWorld[5],
			"virtualMemory":  processEachWorld[4],
			"pageableMemory": processEachWorld[2],
			"handles":        processEachWorld[0],
		}
		processList = append(processList, temp)

	}

	result := map[string]interface{}{
		"host":        host,
		"status":      "success",
		"status.code": 200,
		"process":     processList,
	}
	b, marshalErr := json.Marshal(result)
	if marshalErr != nil {

		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Polling Json",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)
		encode := base64.StdEncoding.EncodeToString(b)
		fmt.Println(encode)
		os.Exit(0)

	}
	encode := base64.StdEncoding.EncodeToString(b)
	fmt.Println(encode)
}
