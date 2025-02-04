package local

import (
	"fmt"
	"gorm.io/gorm"
)

// Icons
var iconAws Icon
var iconEc2 Icon
var iconS3 Icon
var iconMariadb Icon
var iconChatgpt Icon
var iconLockupb Icon
var iconLogomarkb Icon
var iconLockupw Icon
var iconLogomarkw Icon
var iconMysql Icon
var iconStripe Icon

func InsertIcons(db *gorm.DB) {
	fmt.Println("Insert Icons")

	iconAws = Icon{Code: "aws.aws", Directory: "aws", Filename: "aws.svg"}
	iconEc2 = Icon{Code: "aws.ec2", Directory: "aws", Filename: "ec2.svg"}
	iconS3 = Icon{Code: "aws.s3", Directory: "aws", Filename: "s3.svg"}
	iconMariadb = Icon{Code: "mariadb.mariadb", Directory: "mariadb", Filename: "mariadb.svg"}
	iconChatgpt = Icon{Code: "openai.chatgpt", Directory: "openai", Filename: "chatgpt.svg"}
	iconLockupb = Icon{Code: "openai.lockup.b", Directory: "openai", Filename: "openai-lockup.svg"}
	iconLogomarkb = Icon{Code: "openai.logomark.b", Directory: "openai", Filename: "openai-logomark.svg"}
	iconLockupw = Icon{Code: "openai.lockup.w", Directory: "openai", Filename: "white-lockup.svg"}
	iconLogomarkw = Icon{Code: "openai.logomark.w", Directory: "openai", Filename: "white-logomark.svg"}
	iconMysql = Icon{Code: "oracle.mysql", Directory: "oracle", Filename: "mysql.svg"}
	iconStripe = Icon{Code: "stripe.stripe", Directory: "stripe", Filename: "stripe.svg"}

	db.Create(&iconAws)
	db.Create(&iconEc2)
	db.Create(&iconS3)
	db.Create(&iconMariadb)
	db.Create(&iconChatgpt)
	db.Create(&iconLockupb)
	db.Create(&iconLogomarkb)
	db.Create(&iconLockupw)
	db.Create(&iconLogomarkw)
	db.Create(&iconMysql)
	db.Create(&iconStripe)
}
