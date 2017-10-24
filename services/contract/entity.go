package contract

import (
	"fmt"
	"time"

	"github.com/kwri/go-workflow/modules/setting"
)

type StdTime struct {
	time.Time
}

func (t StdTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.Format(setting.Config.TimeFormat))

	return []byte(stamp), nil
}
