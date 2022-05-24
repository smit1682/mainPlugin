package SSH

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func SshPlugin(credentialsMap map[string]interface{}) {

	sshHost := credentialsMap["host"].(string)
	sshUser := credentialsMap["username"].(string)
	sshPassword := credentialsMap["password"].(string)

	sshPort := credentialsMap["port"] //port

	/*if err != nil {

		log.Fatal("Error: Port conversion string to Int", err)
	}*/

	// Create SSHP login configuration
	config := &ssh.ClientConfig{
		Timeout:         10 * time.Second, //ssh connection time out time is one second, if SSH validation error returns in one second
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Config: ssh.Config{Ciphers: []string{
			"aes128-ctr", "aes192-ctr", "aes256-ctr",
		}},
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	// dial gets SSH client
	addr := fmt.Sprintf("%s:%v", sshHost, sshPort)

	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)

	}
	defer sshClient.Close()

	switch credentialsMap["metric.group"] {
	case "cpu":
		cpu_ssh(sshClient, credentialsMap["host"].(string))
	case "system":
		system_ssh(sshClient, credentialsMap["host"].(string))
	case "disk":
		disk_ssh(sshClient, credentialsMap["host"].(string))
	case "process":
		process_ssh(sshClient, credentialsMap["host"].(string))
	case "memory":
		memory_ssh(sshClient, credentialsMap["host"].(string))

	}
	//cpu_ssh(sshClient, credentialsMap["host"])
	/*memory_ssh(sshClient, credentialsMap["host"])
	system_ssh(sshClient, credentialsMap["host"])
	disk_ssh(sshClient, credentialsMap["host"])
	process_ssh(sshClient, credentialsMap["host"])*/

	// Create ssh-session
	/*session, err := sshClient.NewSession()
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

		if flag1 <= 3 {
			flag1++
			continue
		}
		cpuEachWorld := strings.Split(standardizeSpaces(v), " ")
		if len(cpuEachWorld) <= 13 {
			continue
		}

		temp1 := map[string]string{
			"cpu":                   cpuEachWorld[3],
			"cpu.user.percentage":   cpuEachWorld[4],
			"cpu.system.percentage": cpuEachWorld[6],
			"cpu.idle.percentage":   cpuEachWorld[13],
		}
		cpuList = append(cpuList, temp1)

	}

	session, err = sshClient.NewSession()
	var diskList []map[string]string
	if err != nil {
	}
	diskUtilization, err := session.Output("df -h")
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
		temp := map[string]string{
			"disk":      diskEachWorld[0],
			"size":      diskEachWorld[1],
			"used":      diskEachWorld[2],
			"available": diskEachWorld[3],
		}
		diskList = append(diskList, temp)

	}
	//	fmt.Println("Using df -h command fetch Disk usage")
	//fmt.Println(diskUtilizationString)
	session.Close()

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	psaux, err := session.Output("ps aux")
	if err != nil {
		panic(err)
	}
	psauxString := string(psaux)

	myStringArray := strings.Split(psauxString, "\n")

	var processList []map[string]string
	flag := 1
	for _, v := range myStringArray {
		if flag == 1 {
			flag = 0
			continue
		}
		smitEachWorld := strings.SplitN(standardizeSpaces(v), " ", 11)
		if len(smitEachWorld) <= 10 {
			break
		}

		temp1 := map[string]string{
			"user": smitEachWorld[0],
			"pid":  smitEachWorld[1],
			"cpu":  smitEachWorld[2],
			"mem":  smitEachWorld[3],
		}
		processList = append(processList, temp1)

	}

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	freeMemoryP, err := session.Output("free | grep Mem | awk '{ printf(\"%.4f\\n\", $4) }'") //free memory
	session.Close()

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	usedMemoryP, err := session.Output("free | grep Mem | awk '{ printf(\"%.4f %%\\n\", $3) }'") // used memory
	session.Close()

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	totalMemoryP, err := session.Output("free | grep Mem | awk '{ printf(\"%i\\n\", $2) }'") // used memory
	session.Close()

	result := map[string]interface{}{
		"Device":      "linux",
		"process":     processList,
		"cpu":         cpuList,
		"disk":        diskList,
		"freeMemory":  string(freeMemoryP),
		"usedMemory":  string(usedMemoryP),
		"totalMemory": string(totalMemoryP),
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))
	//channel <- string(b)

	session.Close()
	*/
}
