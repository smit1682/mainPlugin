package SNMP

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gosnmp/gosnmp"
	g "github.com/gosnmp/gosnmp"
	"os"
)

func interfaceSnmp(params *gosnmp.GoSNMP, host string) {

	ifIndex := params.Walk(".1.3.6.1.2.1.2.2.1.1", printValueForIfIndex)
	if ifIndex != nil {
	}
	var ifAdmin string
	var ifType string
	var ifInError string
	var ifOutError string
	var ifOperStatus string
	var ifInOctets string
	var ifOutOctets string
	var ifSpeed string
	var ifPhyAddres string
	mainOIDS := []string{".1.3.6.1.2.1.2.2.1.2", ".1.3.6.1.2.1.2.2.1.3", ".1.3.6.1.2.1.2.2.1.7", ".1.3.6.1.2.1.2.2.1.14", ".1.3.6.1.2.1.2.2.1.20", ".1.3.6.1.2.1.2.2.1.8", "1.3.6.1.2.1.2.2.1.10", ".1.3.6.1.2.1.2.2.1.16", "1.3.6.1.2.1.2.2.1.5", ".1.3.6.1.2.1.2.2.1.7"}

	for _, v := range indexList {
		var tempofOIDs []string
		for _, v1 := range mainOIDS {

			v1 += "."
			v1 += v
			tempofOIDs = append(tempofOIDs, v1)
		}
		resultOfOIDs, err := params.Get(tempofOIDs)
		if err != nil {
			panic(err)
		}

		switch resultOfOIDs.Variables[2].Value {
		case 1:
			ifAdmin = "up"
		case 2:
			ifAdmin = "down"

		}

		switch resultOfOIDs.Variables[1].Value {
		case 6:
			ifType = "ethernetCsmacd"
		case 1:
			ifType = "other"
		case 135:
			ifType = "l2vlan"
		case 53:
			ifType = "propVirtual"
		case 24:
			ifType = "softwareLoopback"
		case 131:
			ifType = "tunnel"

		}

		switch resultOfOIDs.Variables[5].Value {
		case 1:

			ifOperStatus = "up"
		case 2:
			ifOperStatus = "down"

		}

		ifInError = fmt.Sprintf("%v", resultOfOIDs.Variables[3].Value)
		ifOutError = fmt.Sprintf("%v", resultOfOIDs.Variables[4].Value)

		ifInOctets = fmt.Sprintf("%v", resultOfOIDs.Variables[6].Value)
		ifOutOctets = fmt.Sprintf("%v", resultOfOIDs.Variables[7].Value)
		ifSpeed = fmt.Sprintf("%d", resultOfOIDs.Variables[8].Value)
		ifPhyAddres = fmt.Sprintf("%x", resultOfOIDs.Variables[9].Value)

		tempMap := map[string]string{
			"IfIndex(.1.3.6.1.2.1.2.2.1.1)":       v,
			"IfAdminStatus(.1.3.6.1.2.1.2.2.1.7)": ifAdmin,
			"IfType(1.3.6.1.2.1.2.2.1.3)":         ifType,
			"IfDescr(.1.3.6.1.2.1.2.2.1.2)":       string(resultOfOIDs.Variables[0].Value.([]byte)),
			"IfInError":                           ifInError,
			"IfOutError":                          ifOutError,
			"IfOperStatus":                        ifOperStatus,
			"IfOutOctets":                         ifOutOctets,
			"IfInOctets":                          ifInOctets,
			"IfSpeed":                             ifSpeed,
			"ifPhyAddres":                         ifPhyAddres,
		}
		mapList = append(mapList, tempMap)

	}

	result1 := map[string]interface{}{
		"host":        host,
		"interface":   mapList,
		"status":      "success",
		"status.code": 200,
	}

	marshal, marshalErr := json.Marshal(result1)
	if marshalErr != nil {

		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Polling Json",
			"status.code": 400,
		}

		marshal, _ := json.Marshal(statusMap)
		encode := base64.StdEncoding.EncodeToString(marshal)
		fmt.Println(encode)
		os.Exit(0)

	}
	encode := base64.StdEncoding.EncodeToString(marshal)
	fmt.Println(encode)

}
func printValueForIfIndex(pdu g.SnmpPDU) error {
	str := fmt.Sprintf("%v", pdu.Value)
	indexList = append(indexList, str)
	return nil
}
