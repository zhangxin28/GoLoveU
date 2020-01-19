package utils

import (
	"strconv"
	"time"
)

const (
	FMT_DATE_TIME    = "2006-01-02 15:04:05"
	FMT_DATE         = "2006-01-02"
	FMT_TIME         = "15:04:05"
	FMT_DATE_TIME_CN = "2006年01月02日 15时04分05秒"
	FMT_DATE_CN      = "2006年01月02日"
	FMT_TIME_CN      = "15时04分05秒"
)

// NowUnix represents second time stamp
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowTimestamp represents millisecond time stamp
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// Timestamp returns millisecond time stamp
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// TimeFromUnix converts a second time stamp to Time
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// TimeFromTimestamp converts a millisecond time stamp to Time
func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Microsecond))
}

// TimeFormat returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// A fractional second is represented by adding a period and zeros
// to the end of the seconds section of layout string, as in "15:04:05.000"
// to format a time stamp with millisecond precision.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

// TimeParse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time,
// defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be interpreted if it were the value; it serves as an example of
// the input format. The same interpretation will then be made to the
// input string.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
// Also, the executable example for Time.Format demonstrates the working
// of the layout string in detail and is a good reference.
//
// Elements omitted from the value are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
func TimeParse(value, layout string) (time.Time, error) {
	return time.Parse(layout, value)
}

// GetDay return a day format with yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

// WithTimeAsStartOfDay 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// PrettyTime represents
/**
 * 将时间格式换成 xx秒前，xx分钟前...
 * 规则：
 * 59秒--->刚刚
 * 1-59分钟--->x分钟前（23分钟前）
 * 1-24小时--->x小时前（5小时前）
 * 昨天--->昨天 hh:mm（昨天 16:15）
 * 前天--->前天 hh:mm（前天 16:15）
 * 前天以后--->mm-dd（2月18日）
 */
func PrettyTime(timestamp int64) string {
	_time := TimeFromTimestamp(timestamp)
	_duration := (NowTimestamp() - timestamp) / 1000
	if _duration < 60 {
		return "刚刚"
	} else if _duration < 3600 {
		return strconv.FormatInt(_duration/60, 10) + "分钟前"
	} else if _duration < 86400 {
		return strconv.FormatInt(_duration/3600, 10) + "小时前"
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24))) <= timestamp {
		return "昨天 " + TimeFormat(_time, FMT_TIME)
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24*2))) <= timestamp {
		return "前天 " + TimeFormat(_time, FMT_TIME)
	} else {
		return TimeFormat(_time, FMT_DATE)
	}
}
