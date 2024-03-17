package date

// DefaultFormat is the default date format
var DefaultFormat = "Y-m-d H:i:s"

// DefaultMapping is the default mapping of date format attributes to Go time layout
// currently it is 1:1 to the PHP date format attributes, but that might change in the future
// therefore feel free to copy this over or override it
var DefaultMapping = map[string]string{
	// day of the month, 2 digits with leading zeros: 01 to 31
	"d": "02",
	// a textual representation of a day, three letters: Mon through Sun
	"D": "Mon",
	// day of the month without leading zeros: 1 to 31
	"j": "2",
	// a full textual representation of the day of the week: Sunday through Saturday
	"l": "Monday",
	// iso-8601 numeric representation of the day of the week: 1 (for Monday) through 7 (for Sunday) - custom handling needed
	"N": "",
	// english ordinal suffix for the day of the month, 2 characters: st, nd, rd, or th - custom handling needed
	"S": "",
	// numeric representation of the day of the week: 0 (for Sunday) through 6 (for Saturday) - custom handling needed
	"w": "",
	// the day of the year (starting from 0): 0 through 365 - custom handling needed
	"z": "",
	// iso-8601 week number of year, weeks starting on Monday: example: 42 (the 42nd week in the year)
	"W": "02",
	// a full textual representation of a month, such as January or March: January through December
	"F": "January",
	// numeric representation of a month, with leading zeros: 01 through 12
	"m": "01",
	// a short textual representation of a month, three letters: Jan through Dec
	"M": "Jan",
	// numeric representation of a month, without leading zeros: 1 through 12
	"n": "1",
	// number of days in the given month: 28 through 31 - custom handling needed
	"t": "",
	// whether it's a leap year: 1 if it is a leap year, 0 otherwise - custom handling needed
	"L": "",
	// iso-8601 week-numbering year. This has the same value as Y, except that if the ISO week number (W) belongs to the previous or next year, that year is used instead.
	"o": "2006",
	// an expanded full numeric representation of a year, at least 4 digits, with - for years BCE, and + for years CE - custom handling needed
	"X": "",
	// an expanded full numeric representation if required, or a standard full numeral representation if possible - custom handling needed
	"x": "",
	// a full numeric representation of a year, at least 4 digits
	"Y": "2006",
	// a two-digit representation of a year
	"y": "06",
	// lowercase ante meridiem and post meridiem: am or pm
	"a": "pm",
	// uppercase ante meridiem and post meridiem: AM or PM
	"A": "PM",
	// swatch internet time: 000 through 999 - custom handling needed
	"B": "",
	// 12-hour format of an hour without leading zeros: 1 through 12
	"g": "3",
	// 24-hour format of an hour without leading zeros: 0 through 23
	"G": "15",
	// 12-hour format of an hour with leading zeros: 01 through 12
	"h": "03",
	// 24-hour format of an hour with leading zeros: 00 through 23
	"H": "15",
	// minutes with leading zeros: 00 to 59
	"i": "04",
	// seconds with leading zeros: 00 through 59
	"s": "05",
	// microseconds - note that date() will always generate 000000 since it takes an int parameter, whereas DateTime::format() does support microseconds if DateTime was created with microseconds.
	"u": "000000",
	// milliseconds - same note applies as for u.
	"v": "000",
	// timezone identifier: examples: UTC, GMT, Atlantic/Azores
	"e": "MST",
	// whether or not the date is in daylight saving time: 1 if Daylight Saving Time, 0 otherwise - custom handling needed
	"I": "",
	// difference to Greenwich time (GMT) without colon between hours and minutes: example: +0200
	"O": "-0700",
	// difference to Greenwich time (GMT) with colon between hours and minutes: example: +02:00
	"P": "-07:00",
	// the same as P, but returns Z instead of +00:00 (available as of PHP 8.0.0): examples: Z or +02:00 - custom handling needed
	"p": "",
	// timezone abbreviation, if known; otherwise the GMT offset: examples: EST, MDT
	"T": "MST",
	// timezone offset in seconds. The offset for timezones west of UTC is always negative, and for those east of UTC is always positive - custom handling needed
	"Z": "",
	// iso 8601 date: 2004-02-12T15:19:21+00:00
	"c": "2006-01-02T15:04:05-07:00",
	// rfc 2822 formatted date: example: Thu, 21 Dec 2000 16:01:07 +0200
	"r": "Mon, 02 Jan 2006 15:04:05 -0700",
	// seconds since the Unix Epoch (January 1 1970 00:00:00 GMT) - custom handling needed
	"U": "",
}
