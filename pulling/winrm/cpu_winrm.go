package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
	"strings"
)

func cpu_winrm(client *winrm.Client, host string) {
	commandForCpu := "Get-WmiObject win32_Processor | select DeviceID, SystemName, LoadPercentage | Foreach-Object {$_.DeviceId,$_.SystemName,$_.LoadPercentage -join \" \"}"
	var cpu string
	var cpuList []map[string]string
	cpu, _, _, err := client.RunPSWithString(commandForCpu, "")
	if err != nil {

	}
	cpuStringArray := strings.Split(cpu, "\n")
	for _, v := range cpuStringArray {
		if len(cpuStringArray) == 0 {
			break
		}
		cpuEachWord := strings.Split(standardizeSpaces(v), " ")
		if len(cpuEachWord) <= 2 {
			break
		}
		temp := map[string]string{
			"cpu":            cpuEachWord[0],
			"sysName":        cpuEachWord[1],
			"loadPercentage": cpuEachWord[2],
		}
		cpuList = append(cpuList, temp)
	}

	result := map[string]interface{}{
		"host":        host,
		"status":      "success",
		"status.code": 200,
		"cpu":         cpuList,
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
