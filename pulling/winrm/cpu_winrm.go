package WINRM

import (
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
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
		"host": host,
		"cpu":  cpuList,
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
