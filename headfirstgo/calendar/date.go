package calendar

import (
	"errors"
)

// Date is a struct with year/month/day,each of is type int
type Date struct {
	year  int
	month int
	day   int
}

// GetDaysInMonth returns days in month
// Go doesn’t allow complex types like
// slices, maps, or arrays to be constant!
func GetDaysInMonth() map[int][]int {
	return map[int][]int{
		1:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		2:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		3:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		4:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		5:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		6:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		7:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		8:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		9:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		10: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		11: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		12: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	}
}

// SetYear sets year for Date
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}

	d.year = year

	return nil
}

// SetMonth sets month for Date
func (d *Date) SetMonth(month int) error {
	_, err := getDaysInMonth(month)
	if err != nil {
		return err
	}

	d.month = month

	return nil
}

// SetDay sets day for Date
func (d *Date) SetDay(day int) error {
	days, err := getDaysInMonth(d.month)
	if err != nil {
		return err
	}

	for _, dayInMonth := range days {
		if day == dayInMonth {
			d.day = day
			return nil
		}
	}

	return errors.New("invalid day")
}

/*
By convention, a getter method's name should be the same
as the name of the field or variable it accesses.
Of course, if you want the method to be exported,
its name will need to start with a capital letter.
Getter methods don't need to modify the receiver at all,
so we could use the value receiver.
But if any method on a type takes a pointer receiver,
convention says that all should, for consistency's sake,
use all pointer receiver or value receiver.
*/

// Year returns year for Date
func (d *Date) Year() int {
	return d.year
}

// Month returns month for Date
func (d *Date) Month() int {
	return d.month
}

// Day returns day for Date
func (d *Date) Day() int {
	return d.day
}

func getDaysInMonth(month int) ([]int, error) {
	days, ok := GetDaysInMonth()[month]

	if !ok {
		return nil, errors.New("invalid month")
	}

	return days, nil
}
