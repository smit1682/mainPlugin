package SSH

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

func cpu_ssh(sshClient *ssh.Client, host string) {

	session, err := sshClient.NewSession()
	if err != nil {
	}

	// Execute remote commands

	//combo, err := session.Output("df -h")                    //pwd; free -m; snmpwalk -H Combined output
	cpuUtilization, err := session.CombinedOutput("mpstat -P ALL") // it will show you all 4 processors usage

	var cpuList []map[string]string
	cpuUtilizationString := string(cpuUtilization)
	cpuStringArray := strings.Split(cpuUtilizationString, "\n")

	//fmt.Println(len(cpuStringArray))

	flag1 := 1
	for _, v := range cpuStringArray {

		if flag1 <= 4 {
			flag1++
			continue
		}
		cpuEachWorld := strings.Split(standardizeSpaces(v), " ")
		if len(cpuEachWorld) <= 13 {
			continue
		}

		temp1 := map[string]string{
			"cpu.name":              cpuEachWorld[3],
			"cpu.user.percentage":   cpuEachWorld[4],
			"cpu.system.percentage": cpuEachWorld[6],
			"cpu.idle.percentage":   cpuEachWorld[13],
		}
		cpuList = append(cpuList, temp1)

	}

	session, err = sshClient.NewSession()
	if err != nil {
	}

	scalercpu, err := session.CombinedOutput("mpstat | grep \"all\"")
	scalercpuString := string(scalercpu)

	eachWord := strings.Split(standardizeSpaces(scalercpuString), " ")

	session, err = sshClient.NewSession()
	if err != nil {
	}
	cpup, err := session.Output("mpstat | awk 'NR == 4 {print ($5+$6+$7+$8+$9+$10+$11+$12+$13)}'")
	cpuString := string(cpup)
	smit := strings.Split(standardizeSpaces(cpuString), " ")

	result := map[string]interface{}{
		"host":                      host,
		"status":                    "success",
		"status.code":               200,
		"cpu.all.user.percentage":   eachWord[4],
		"cpu.all.system.percentage": eachWord[6],
		"cpu.all.idle.percentage":   eachWord[13],
		"cpu.all.percentage":        smit[0],
		"cpu":                       cpuList,
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
