package initializers

import (
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/imtiaz246/codera_oj/initializers/db"
)

func Initialize() error {
	if err := config.LoadConfigs(); err != nil {
		return err
	}
	if err := db.InitializeDB(); err != nil {
		return err
	}

	return nil
}
