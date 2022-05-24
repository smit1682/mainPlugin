package SNMP

import (
	"encoding/json"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"log"
)

func system_snmp(params *gosnmp.GoSNMP) {

	oids := []string{"1.3.6.1.2.1.1.5.0", "1.3.6.1.2.1.1.6.0", "1.3.6.1.2.1.1.1.0", "1.3.6.1.2.1.1.2.0", "1.3.6.1.2.1.1.3.0"} // sysName , sysLocation, sysDiscription,sysOID,

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

	result1 := map[string]interface{}{
		"sysName":        sysName,
		"sysDescription": sysDiscryption,
		"sysLocation":    sysLocation,
		"sysOID":         sysOID,
		"sysUpTime":      sysUpTime,
	}

	b, _ := json.Marshal(result1)
	fmt.Println(string(b))

}
