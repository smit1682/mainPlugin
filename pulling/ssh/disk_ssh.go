package SSH

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
)

func disk_ssh(sshClient *ssh.Client, host string) {

	session, err := sshClient.NewSession()
	var diskList []map[string]string
	if err != nil {
	}
	diskUtilization, err := session.Output("df")
	if err != nil {
		panic(err)
	}
	diskUtilizationString := string(diskUtilization)

	diskStringArray := strings.Split(diskUtilizationString, "\n")
	flag2 := 1
	for _, v := range diskStringArray {
		if flag2 == 1 {
			flag2++
			continue
		}
		diskEachWorld := strings.Split(standardizeSpaces(v), " ")
		if len(diskEachWorld) <= 5 {
			continue
		}

		totalDisk, err := strconv.ParseInt(diskEachWorld[1], 10, 64)
		if err != nil {
		}
		totalDisk = totalDisk * 1024

		usedDisk, err := strconv.ParseInt(diskEachWorld[2], 10, 64)
		if err != nil {
		}
		usedDisk = usedDisk * 1024

		freeDisk, err := strconv.ParseInt(diskEachWorld[3], 10, 64)
		if err != nil {
		}
		freeDisk = freeDisk * 1024

		/*usedPercent := float64((usedDisk / totalDisk) * 100)
		fmt.Println("used", usedPercent)
		freePercent := float64((freeDisk / totalDisk) * 100.00)
		fmt.Println("free", freePercent)

		usedPercent1 := fmt.Sprintf("%.3f", usedPercent)
		freePercent1 := fmt.Sprintf("%.3f", freePercent)*/
		//fmt.Println(diskEachWorld[4])
		usedper, err := strconv.ParseInt(strings.TrimRight(diskEachWorld[4], "%"), 10, 64)

		temp := map[string]string{
			"disk.volume.name":            diskEachWorld[0],
			"disk.volume.total.bytes":     strconv.FormatInt(totalDisk, 10),
			"disk.volume.used.bytes":      strconv.FormatInt(usedDisk, 10),
			"disk.volume.free.bytes":      strconv.FormatInt(freeDisk, 10),
			"disk.volume.used.percentage": strings.TrimRight(diskEachWorld[4], "%"),
			"disk.volume.free.percentage": strconv.FormatInt(100-usedper, 10),
		}
		diskList = append(diskList, temp)

	}

	session, err = sshClient.NewSession()
	if err != nil {
	}
	totaldisk, err := session.Output("df --total | grep 'total'")

	totaldiskparts := strings.Split(standardizeSpaces(string(totaldisk)), " ")

	tDisk, err := strconv.ParseInt(totaldiskparts[1], 10, 64)
	if err != nil {
	}
	tDisk = tDisk * 1024

	udisk, err := strconv.ParseInt(totaldiskparts[2], 10, 64)
	if err != nil {
	}
	udisk = udisk * 1024

	fDisk, err := strconv.ParseInt(totaldiskparts[3], 10, 64)
	if err != nil {
	}
	fDisk = fDisk * 1024

	u, err := strconv.ParseInt(strings.TrimRight(totaldiskparts[4], "%"), 10, 64)

	result := map[string]interface{}{
		"host":                 host,
		"disk":                 diskList,
		"disk.total.bytes":     strconv.FormatInt(tDisk, 10),
		"disk.used.bytes":      strconv.FormatInt(udisk, 10),
		"disk.free.bytes":      strconv.FormatInt(fDisk, 10),
		"disk.used,percentage": strings.TrimRight(totaldiskparts[4], "%"),
		"disk.free.percentage": strconv.FormatInt(100-u, 10),
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))

}
