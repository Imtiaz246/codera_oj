package cronera

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getWaitingDuration(t time.Time) time.Duration {
	curTime := time.Now()
	dur := t.Sub(curTime)
	if dur > 0 {
		return dur
	} else {
		return time.Duration(0)
	}
}

func parseTodayTimes(st []string, today time.Time) []time.Time {
	todayTimes := make([]time.Time, 0)
	for _, s := range st {
		todayTimes = append(todayTimes, parseForToday(s, today))
	}
	sort.Slice(todayTimes, func(i, j int) bool {
		return todayTimes[i].Before(todayTimes[j])
	})
	return todayTimes
}

func parseForToday(st string, today time.Time) time.Time {
	ss := strings.Split(st, ":")
	if len(ss) == 1 {
		ss = append(ss, "00")
	}
	hour, err := strconv.Atoi(ss[0])
	if err != nil {
		hour = 0
	}
	minute, err := strconv.Atoi(ss[1])
	if err != nil {
		minute = 0
	}

	return time.Date(today.Year(), today.Month(), today.Day(), hour%24, minute%60, 0, 0, today.Location())
}

func waitUntil(t time.Time) <-chan time.Time {
	dur := getWaitingDuration(t)
	return time.After(dur)
}

func execTaskFunc(taskFunc any, args ...any) []reflect.Value {
	in := make([]reflect.Value, 0)
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	return reflect.ValueOf(taskFunc).Call(in)
}

func truncateToStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

func parseDate(d string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, d)
	if err != nil {
		return time.Time{}, err
	} else {
		return t, nil
	}
}

func checkTaskFuncValidity(taskFunc any, args ...any) error {
	ft := reflect.TypeOf(taskFunc)
	if ft.Kind() != reflect.Func {
		return fmt.Errorf("taskFunc is not a func type")
	}
	if ft.NumIn() != len(args) {
		return fmt.Errorf("input no doesn't match argument list")
	}

	for i := 0; i < ft.NumIn(); i++ {
		if ft.In(i) != reflect.TypeOf(args[i]) {
			return fmt.Errorf("mismatch argument no %v. Expected %v, given %v", i, ft.In(i).Kind(), reflect.TypeOf(args[i]).Kind())
		}
	}

	return nil
}
