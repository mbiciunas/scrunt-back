package local

import (
	"fmt"
	"scrunt-back/models/scrunt/icon"
)

// Icons
var iconIdAws uint
var iconIdEc2 uint
var iconIdS3 uint
var iconIdMariadb uint
var iconIdChatgpt uint
var iconIdLockupb uint
var iconIdLogomarkb uint
var iconIdLockupw uint
var iconIdLogomarkw uint
var iconIdMysql uint
var iconIdStripe uint

func InsertIcons() (err error) {
	fmt.Println("Insert Icons")

	iconIdAws, err = icon.GormInsertIcon("aws.aws", "aws", "aws.svg")
	iconIdEc2, err = icon.GormInsertIcon("aws.ec2", "aws", "ec2.svg")
	iconIdS3, err = icon.GormInsertIcon("aws.s3", "aws", "s3.svg")
	iconIdMariadb, err = icon.GormInsertIcon("mariadb.mariadb", "mariadb", "mariadb.svg")
	iconIdChatgpt, err = icon.GormInsertIcon("openai.chatgpt", "openai", "chatgpt.svg")
	iconIdLockupb, err = icon.GormInsertIcon("openai.lockup.b", "openai", "openai-lockup.svg")
	iconIdLogomarkb, err = icon.GormInsertIcon("openai.logomark.b", "openai", "openai-logomark.svg")
	iconIdLockupw, err = icon.GormInsertIcon("openai.lockup.w", "openai", "white-lockup.svg")
	iconIdLogomarkw, err = icon.GormInsertIcon("openai.logomark.w", "openai", "white-logomark.svg")
	iconIdMysql, err = icon.GormInsertIcon("oracle.mysql", "oracle", "mysql.svg")
	iconIdStripe, err = icon.GormInsertIcon("stripe.stripe", "stripe", "stripe.svg")

	return err
}
