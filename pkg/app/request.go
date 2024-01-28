package app

import (
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(error []*validation.Error) {
	for _, err := range error {
		logging.Info(err.Key, err.Message)
	}

	return
}
