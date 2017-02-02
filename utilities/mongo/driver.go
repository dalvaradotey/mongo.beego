package mongo

import (
	"bytes"
	"github.com/astaxie/beego"
)

var maxPool int

func init() {
	var err error

	maxPool, err = beego.AppConfig.Int("mongo::maxpool")

	if err != nil {
		println(err)
	}

	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.BaseSession == nil {
		var string_connection bytes.Buffer

		string_connection.WriteString("mongodb://")
		string_connection.WriteString(beego.AppConfig.String("mongo::user"))
		string_connection.WriteString(":")
		string_connection.WriteString(beego.AppConfig.String("mongo::password"))
		string_connection.WriteString("@")
		string_connection.WriteString(beego.AppConfig.String("mongo::host"))
		string_connection.WriteString(":")
		string_connection.WriteString(beego.AppConfig.String("mongo::port"))
		string_connection.WriteString("/")
		string_connection.WriteString(beego.AppConfig.String("mongo::database"))

		service.Url = string_connection.String()

		err := service.New()

		if err != nil {
			println(err)
		}
	}
}
