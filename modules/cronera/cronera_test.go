package cronera

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func Test_OneTimeJob(t *testing.T) {
	curTime := time.Now()
	firstTime := fmt.Sprintf("%s:%s", strconv.Itoa(curTime.Hour()), strconv.Itoa(curTime.Minute()))

	curTime = curTime.Add(time.Minute)
	secondTime := fmt.Sprintf("%s:%s", strconv.Itoa(curTime.Hour()), strconv.Itoa(curTime.Minute()))

	JobFunc := func(arg int) int {
		return arg
	}

	sigStream, err := New().Every(1).Day().At(firstTime, secondTime).MustExec(1).Do(context.Background(), JobFunc, 10)
	require.Equal(t, err, nil)

	jobChecked := false
	for signal := range sigStream {
		if signal.status == fmt.Sprintf(string(StatusTaskSuccessful), 1) {
			output := signal.returnValues[0].Interface().(int)
			require.Equal(t, output, 10)
			jobChecked = true
		}
	}

	require.Equal(t, jobChecked, true)
}
