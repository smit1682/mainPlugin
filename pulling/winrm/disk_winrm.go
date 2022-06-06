package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
	"strings"
)

func disk_winrm(client *winrm.Client, host string) {

	commandForDisk := "Get-WmiObject win32_logicaldisk | Foreach-Object {$_.DeviceId,$_.Freespace,$_.Size -join \" \"}" //disksize

	disk, _, _, err := client.RunPSWithString(commandForDisk, "")

	if err != nil {
	}
	var diskList []map[string]string
	diskStringArray := strings.Split(disk, "\n")

	for _, v := range diskStringArray {
		diskEachWord := strings.Split(standardizeSpaces(v), " ")

		if len(diskEachWord) == 0 {
			break
		}
		if len(diskEachWord) == 3 {
			temp := map[string]string{
				"disk": diskEachWord[0],
				"free": diskEachWord[1],
				"size": diskEachWord[2],
			}
			diskList = append(diskList, temp)

		}
		if len(diskEachWord) == 1 {
			temp := map[string]string{
				"disk": diskEachWord[0],
				"free": "0",
				"size": "0",
			}
			diskList = append(diskList, temp)
		}

	}

	result := map[string]interface{}{
		"host":        host,
		"status":      "success",
		"status.code": 200,
		"disk":        diskList,
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
