package WINRM

import (
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
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
		"freeMemory":         memoryStringArray[1],
		"freeVirtualMemory":  memoryStringArray[3],
		"totalMemory":        memoryStringArray[0],
		"totalVirtualMemory": memoryStringArray[2],
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
