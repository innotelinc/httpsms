package main

import (
	"fmt"
	"os"

	"github.com/innotelinc/httpsms/pkg/di"
)

// Version is injected at runtime
var Version string

// @title       HTTP SMS API
// @version     1.0
// @description API to send SMS messages using android [SmsManager](https://developer.android.com/reference/android/telephony/SmsManager) via HTTP
//
// @contact.name  HTTP SMS
// @contact.email support@httpsms.com
//
// @license.name AGPL-3.0
// @license.url  https://raw.githubusercontent.com/NdoleStudio/http-sms-manager/main/LICENSE
//
// @host     api.httpsms.com
// @schemes  https
// @BasePath /v1
//
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name x-api-Key
func main() {
	if len(os.Args) == 1 {
		di.LoadEnv()
	}

	container := di.NewContainer("http-sms", Version)
	container.Logger().Info(container.App().Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))).Error())
}
