package local

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Scripts
var scriptEC2 Script
var scriptMySQL Script
var scriptChatGPT Script

func InsertScripts(db *gorm.DB) {
	fmt.Println("Insert Scripts")

	scriptEC2 = Script{IconId: iconEc2.Id, Name: "Create EC2 Instance", DescShort: "Local - Creates an EC2 instance", DescLong: "This script will create an EC2 instance in your AWS account.  The instance will be set up in the specified region using the specified instance type and operating system.  If specified, an EBS instance will be allocated and attached as well.\\n\\nIf the required region, instance type, operating system or EBS size is not available, the script will not create the EC2 Instance and will return an error.\\n\\nUpon successful operation, the instance will be available via cryptographic key (provided by user).  Also, a new service will be added to Scrunt representing the EC2 instance.", Created: time.Date(2023, 8, 1, 8, 34, 47, 0, time.Local)}
	scriptMySQL = Script{IconId: iconMysql.Id, Name: "MySql Backup/Restore", DescShort: "Local - Backup a MySQL database to the local machine.  The script will ask for the database to backup", DescLong: "This script will backup an existing MySQL database using the MySQL dump  utility.  Both the schema and all data will be backed up.  The backup file will be downloaded to the local machine (running Scrunt).\\n\\nIf a database is specified in the parameters, that database will be backed up.  If no database is specified, all databases will be backed up in individual files and transferred to the local machine.\\n\\nUpon successful operation, the backup files will be available in the backup section of scrunt.", Created: time.Date(2023, 4, 16, 14, 53, 19, 0, time.Local)}
	scriptChatGPT = Script{IconId: iconChatgpt.Id, Name: "ChatGPT Interface", DescShort: "Local - Send a question to ChatGPT and receive answer.  A valid ChatGPT API key is required", DescLong: "Connect to your ChatGPT instance via the API.  The question parameter value will be sent to Chat GPT and the results displayed in the output.", Created: time.Date(2023, 9, 24, 21, 4, 17, 0, time.Local)}

	db.Create(&scriptEC2)
	db.Create(&scriptMySQL)
	db.Create(&scriptChatGPT)
}
