package timerange

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrValueIsEmpty = errors.New("ERROR: value is empty.")
	ErrInvalidRange = errors.New("ERROR: values in time range are invalid.")
)

var (
	DefaultTimeLayout     = "2006/01/02"
	DefaultRangeSeperator = ".."
)

type Timerange struct {
	TimeValues     []time.Time
	TimeLayout     string
	RangeSeparator string
}

func NewTimerange() *Timerange {
	return &Timerange{
		TimeValues:     []time.Time{},
		TimeLayout:     DefaultTimeLayout,
		RangeSeparator: DefaultRangeSeperator,
	}
}

func (t *Timerange) String() string {
	return fmt.Sprint(t.TimeValues)
}

func (t *Timerange) Set(value string) error {
	if value == "" {
		return ErrValueIsEmpty
	}

	var (
		timeValues []time.Time
		err        error
	)

	if t.hasRangeSeperator(value) {
		timeValues, err = t.parseRangeIntoTimeValues(value)
	} else {
		var tt time.Time

		tt, err = t.parseTimeFromValue(value)
		timeValues = []time.Time{tt}
	}

	if err != nil {
		return err
	}

	t.TimeValues = append(t.TimeValues, timeValues...)

	return nil
}

func (t *Timerange) hasRangeSeperator(value string) bool {
	return strings.Contains(value, t.RangeSeparator)
}

func (t *Timerange) parseTimeFromValue(value string) (time.Time, error) {
	return time.Parse(t.TimeLayout, value)
}

func (t *Timerange) parseRangeIntoTimeValues(rangeValue string) (timeValues []time.Time, err error) {
	split := strings.Split(rangeValue, t.RangeSeparator)
	if len(split) != 2 {
		return nil, ErrInvalidRange
	}

	startValue, endValue := split[0], split[1]

	startDate, err := t.parseTimeFromValue(startValue)
	if err != nil {
		return nil, err
	}

	endDate, err := t.parseTimeFromValue(endValue)
	if err != nil {
		return nil, err
	}

	duration := startDate.Sub(endDate).Hours()
	if duration >= 0 {
		return nil, fmt.Errorf(
			"Expected timestamp range start date: '%s' to be before end date: '%s'",
			startDate, endDate,
		)
	}

	durationInDays := (duration / 24) * -1

	for i := 0; i <= int(durationInDays); i++ {
		timeValues = append(timeValues, startDate.Add(time.Duration(i)*time.Duration(24)*time.Hour))
	}

	return timeValues, nil
}
