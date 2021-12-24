package api

import (
	"net/http"
	"testing"

	"ftk8s/base/cfg"
	"ftk8s/router"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI(t *testing.T) {
	cfg.InitAll()
	handler := router.InitRouter()

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		// Printers: []httpexpect.Printer{
		// 	httpexpect.NewDebugPrinter(t, true),
		// },
	})

	// ####################### system ###########################
	{
		// Manage association: role - permission
		ReadPermissionByRole(e)
	}

	// ###################### resource ##########################

}
