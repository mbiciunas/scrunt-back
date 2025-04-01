package local

import (
	"fmt"
	"scrunt-back/models/scrunt/script"
	"time"
)

// Scripts
var scriptIdEC2 uint
var scriptIdMySQL uint
var scriptIdChatGPT uint

func InsertScripts() (err error) {
	fmt.Println("Insert Scripts")

	scriptIdEC2, err = script.GormInsertScript("Create EC2 Instance", iconCodeEc2, "Local - Creates an EC2 instance", "This script will create an EC2 instance in your AWS account.  The instance will be set up in the specified region using the specified instance type and operating system.  If specified, an EBS instance will be allocated and attached as well.\\n\\nIf the required region, instance type, operating system or EBS size is not available, the script will not create the EC2 Instance and will return an error.\\n\\nUpon successful operation, the instance will be available via cryptographic key (provided by user).  Also, a new service will be added to Scrunt representing the EC2 instance.", time.Date(2023, 8, 1, 8, 34, 47, 0, time.Local))
	scriptIdMySQL, err = script.GormInsertScript("MySql Backup/Restore", iconCodeMysql, "Local - Backup a MySQL database to the local machine.  The script will ask for the database to backup", "This script will backup an existing MySQL database using the MySQL dump  utility.  Both the schema and all data will be backed up.  The backup file will be downloaded to the local machine (running Scrunt).\\n\\nIf a database is specified in the parameters, that database will be backed up.  If no database is specified, all databases will be backed up in individual files and transferred to the local machine.\\n\\nUpon successful operation, the backup files will be available in the backup section of scrunt.", time.Date(2023, 4, 16, 14, 53, 19, 0, time.Local))
	scriptIdChatGPT, err = script.GormInsertScript("ChatGPT Interface", iconCodeChatgpt, "Local - Send a question to ChatGPT and receive answer.  A valid ChatGPT API key is required", "Connect to your ChatGPT instance via the API.  The question parameter value will be sent to Chat GPT and the results displayed in the output.", time.Date(2023, 9, 24, 21, 4, 17, 0, time.Local))

	return err
}
