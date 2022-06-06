package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
	"strings"
)

func memory_winrm(client *winrm.Client, host string) {

	commandForMemory := "Get-WmiObject win32_OperatingSystem |%{\"{0} {1} {2} {3}\" -f $_.totalvisiblememorysize, $_.freephysicalmemory, $_.totalvirtualmemorysize, $_.freevirtualmemory} "
	memory, _, _, err := client.RunPSWithString(commandForMemory, "")
	if err != nil {
	}
	memoryStringArray := strings.Split(standardizeSpaces(memory), " ")

	result := map[string]interface{}{
		"host":               host,
		"status":             "success",
		"status.code":        200,
		"freeMemory":         memoryStringArray[1],
		"freeVirtualMemory":  memoryStringArray[3],
		"totalMemory":        memoryStringArray[0],
		"totalVirtualMemory": memoryStringArray[2],
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
