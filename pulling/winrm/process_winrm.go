package WINRM

import (
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
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
		"host":    host,
		"process": processList,
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
