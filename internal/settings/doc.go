// Package settings set up various global variables from config file
// for internal usage. This variable will be accessible from anywhere.
// But especially for internal usages.
package settings

import (
	"github.com/imtiaz246/codera_oj/internal/codera/judger"
)

var (
	judgerHub = judger.NewJudgerHub()
)
