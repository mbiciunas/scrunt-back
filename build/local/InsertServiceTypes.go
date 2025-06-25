package local

import (
	"fmt"
	"scrunt-back/models/scrunt/servicetype"
)

var serviceTypeIdEC2 uint
var serviceTypeIdMySQL uint
var serviceTypeIdChatGPT uint

func InsertServiceTypes() (err error) {
	fmt.Println("Insert Scripts")

	serviceTypeIdEC2 = insertServiceType("EC2", iconCodeEc2, "Amazon EC2 instance", "This service is for an AWS EC2 instance.")
	serviceTypeIdMySQL = insertServiceType("MySql", iconCodeMysql, "MySql server", "This service is for a MySQL database instance.")
	serviceTypeIdChatGPT = insertServiceType("ChatGPT", iconCodeChatgpt, "Char GPT", "This service is for a ChatGTP connection.")

	return err
}

func insertServiceType(name string, icon string, descShort string, descLong string) uint {
	uuid := servicetype.GenerateServiceTypeUUID(name, descShort)
	scriptId, err := servicetype.GormInsertServiceType(uuid, name, icon, descShort, descLong)

	if err != nil {
		panic(err)
	}

	return scriptId
}
