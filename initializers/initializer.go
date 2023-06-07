package initializers

import (
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/imtiaz246/codera_oj/initializers/db"
	"github.com/imtiaz246/codera_oj/initializers/session_cache"
)

func Initialize() error {
	if err := config.LoadConfigs(); err != nil {
		return err
	}
	if err := db.InitializeDB(); err != nil {
		return err
	}
	if err := session_cache.LoadSessionCache(); err != nil {
		return err
	}

	return nil
}
