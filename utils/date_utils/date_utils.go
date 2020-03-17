package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05.05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

//GetNow ...
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString ...
func GetNowString() string {
	now := GetNow()
	//format of date in yy/mm/dd format
	//try this in standard time zone so it
	//remains even throughout all the countries  of the world
	return now.Format(apiDateLayout)
}

//GetNowDBFormat ...
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
