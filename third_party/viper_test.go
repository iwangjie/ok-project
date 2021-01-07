package third_party

import (
	"github.com/spf13/viper"
	"testing"
)

func TestSetupConfiguration(t *testing.T) {
	SetupConfiguration("../")
	if viper.Get("app.name")!="ok_project" {
		t.Error("app.name match failed")
	}
}
