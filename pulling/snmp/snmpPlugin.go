package SNMP

import (
	g "github.com/gosnmp/gosnmp"
	"strconv"
	"time"
)

var mapList []map[string]string
var index = 0
var indexList []string

func SnmpPlugin(credentialsMap map[string]interface{}) {

	port, err := strconv.Atoi(credentialsMap["port"].(string)) //port
	params := &g.GoSNMP{
		Target:    credentialsMap["host"].(string),
		Port:      uint16(port),
		Community: credentialsMap["password"].(string),
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		//Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}
	err = params.Connect()
	if err != nil {
		//log.Fatalf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	switch credentialsMap["metric.group"] {
	case "system":
		system_snmp(params)
	case "interface":
		interfaceSnmp(params)

	}
	/*	oids := []string{"1.3.6.1.2.1.1.5.0", "1.3.6.1.2.1.1.6.0", "1.3.6.1.2.1.1.1.0", "1.3.6.1.2.1.1.2.0", "1.3.6.1.2.1.1.3.0"} // sysName , sysLocation, sysDiscription,sysOID,

		result, err2 := params.Get(oids) // Get() accepts up to g.MAX_OIDS
		if err2 != nil {
			log.Fatalf("Get() err: %v", err2)
		}
		var sysName string
		var sysLocation string
		var sysDiscryption string
		var sysOID string
		var sysUpTime string
		for _, variable := range result.Variables {

			switch variable.Name {
			case ".1.3.6.1.2.1.1.5.0":
				sysName = string(variable.Value.([]byte))
			case ".1.3.6.1.2.1.1.6.0":
				sysLocation = string(variable.Value.([]byte))
			case ".1.3.6.1.2.1.1.1.0":
				sysDiscryption = string(variable.Value.([]byte))
			case ".1.3.6.1.2.1.1.2.0":
				str := fmt.Sprintf("%v", variable.Value)
				sysOID = str
			case ".1.3.6.1.2.1.1.3.0":
				str := fmt.Sprintf("%v", variable.Value)
				sysUpTime = str
			}
		}

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
			"sysName":        sysName,
			"sysDescription": sysDiscryption,
			"sysLocation":    sysLocation,
			"sysOID":         sysOID,
			"sysUpTime":      sysUpTime,
			"interface":      mapList,
		}

		b, _ := json.Marshal(result1)
		fmt.Println(string(b))
	*/
}
