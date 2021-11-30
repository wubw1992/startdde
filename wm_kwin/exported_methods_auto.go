// Code generated by "dbusutil-gen em -type Switcher"; DO NOT EDIT.

package wm_kwin

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (v *Switcher) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "AllowSwitch",
			Fn:      v.AllowSwitch,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:    "CurrentWM",
			Fn:      v.CurrentWM,
			OutArgs: []string{"outArg0"},
		},
		{
			Name: "RequestSwitchWM",
			Fn:   v.RequestSwitchWM,
		},
	}
}
