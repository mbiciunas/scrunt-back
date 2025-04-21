package local

import (
	"fmt"
	"scrunt-back/models/scrunt/icon"
)

// Icons
const iconCodeAws = "aws.aws"
const iconCodeEc2 = "aws.ec2"
const iconCodeS3 = "aws.s3"
const iconCodeMariadb = "mariadb.mariadb"
const iconCodeChatgpt = "openai.chatgpt"
const iconCodeLockupB = "openai.lockup.b"
const iconCodeLogoMarkB = "openai.logomark.b"
const iconCodeLockupW = "openai.lockup.w"
const iconCodeLogoMarkW = "openai.logomark.w"
const iconCodeMysql = "oracle.mysql"
const iconCodeStripe = "stripe.stripe"

func InsertIcons() (err error) {
	fmt.Println("Insert Icons")

	_, err = icon.GormInsertIcon(iconCodeAws, "aws", "aws.svg")
	_, err = icon.GormInsertIcon(iconCodeEc2, "aws", "ec2.svg")
	_, err = icon.GormInsertIcon(iconCodeS3, "aws", "s3.svg")
	_, err = icon.GormInsertIcon(iconCodeMariadb, "mariadb", "mariadb.svg")
	_, err = icon.GormInsertIcon(iconCodeChatgpt, "openai", "chatgpt.svg")
	_, err = icon.GormInsertIcon(iconCodeLockupB, "openai", "openai-lockup.svg")
	_, err = icon.GormInsertIcon(iconCodeLogoMarkB, "openai", "openai-logomark.svg")
	_, err = icon.GormInsertIcon(iconCodeLockupW, "openai", "white-lockup.svg")
	_, err = icon.GormInsertIcon(iconCodeLogoMarkW, "openai", "white-logomark.svg")
	_, err = icon.GormInsertIcon(iconCodeMysql, "oracle", "mysql.svg")
	_, err = icon.GormInsertIcon(iconCodeStripe, "stripe", "stripe.svg")

	return err
}
