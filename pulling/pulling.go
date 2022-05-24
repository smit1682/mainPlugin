package Polling

import (
	SNMP "mainPlugin/pulling/snmp"
	SSH "mainPlugin/pulling/ssh"
	WINRM "mainPlugin/pulling/winrm"
)

func Pulling(credentialsMap map[string]interface{}) {

	switch credentialsMap["metric.type"] {

	case "linux":
		SSH.SshPlugin(credentialsMap)
	case "network.device":
		SNMP.SnmpPlugin(credentialsMap)
	case "windows":
		WINRM.WinrmPlugin(credentialsMap)

	}

}
