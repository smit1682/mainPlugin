package Discovery

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
	"time"
)

func sshDiscovery(credentialsMap map[string]interface{}) {

	sshHost := credentialsMap["host"]
	sshUser := credentialsMap["username"].(string)
	sshPassword := credentialsMap["password"].(string)
	sshPort := credentialsMap["port"]

	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second, //ssh connection time out time is one second, if SSH validation error returns in one second
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Config: ssh.Config{Ciphers: []string{
			"aes128-ctr", "aes192-ctr", "aes256-ctr",
		}},
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	addr := fmt.Sprintf("%s:%v", sshHost, sshPort)

	sshClient, errDial := ssh.Dial("tcp", addr, config)

	if errDial != nil {

		if strings.Contains(errDial.Error(), "handshake failed") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "UNAUTHORIZED | reason: wrong credentials(Username or Password)",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)

		} else if strings.Contains(errDial.Error(), "connection refused") {

			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       "ERR_CONNECTION_REFUSED | reason: port is not available",
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			log.SetFlags(0)
			log.Fatal(encode)
		} else {
			statusMap := map[string]interface{}{
				"status":      "error",
				"error":       errDial.Error(),
				"status.code": 400,
			}

			b, _ := json.Marshal(statusMap)

			encode := base64.StdEncoding.EncodeToString(b)
			fmt.Println(encode)
			os.Exit(0)
		}

	}

	session, errSession := sshClient.NewSession()

	if errSession != nil {
		log.SetFlags(0)
		log.Println(errSession)
	}

	commandOutput, errCommand := session.CombinedOutput("whoami")

	if errCommand != nil {

		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "NOT_SUPPORTED_DEVICE | reason: linux command `whoami` failed",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)

	} else {

		statusMap := map[string]interface{}{
			"status":       "success",
			"status.code":  200,
			"monitor.name": strings.TrimRight(string(commandOutput), "\n"),
		}

		b, _ := json.Marshal(statusMap)

		encode := base64.StdEncoding.EncodeToString(b)
		log.SetFlags(0)
		log.Fatal(encode)
	}

}
