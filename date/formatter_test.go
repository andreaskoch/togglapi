package date

import (
	"testing"
	"time"
)

func Test_NewISO8601Formatter_FormatterIsReturned(t *testing.T) {
	// act
	formatter := NewISO8601Formatter()

	// assert
	if formatter == nil {
		t.Fail()
		t.Logf("NewISO8601Formatter() should not return nil")
	}
}

func Test_GetDateString_DatesAreISO8061Formatted(t *testing.T) {
	// arrange
	berlinTimeZone, _ := time.LoadLocation("Europe/Berlin")
	newYorkTimeZone, _ := time.LoadLocation("America/New_York")

	inputs := []struct {
		Date           time.Time
		ExpectedResult string
	}{
		{
			Date:           time.Date(2015, 3, 27, 7, 42, 35, 0, time.UTC),
			ExpectedResult: "2015-03-27T07:42:35+00:00",
		},

		{
			Date:           time.Date(2015, 1, 1, 0, 0, 1, 0, time.UTC),
			ExpectedResult: "2015-01-01T00:00:01+00:00",
		},

		{
			Date:           time.Date(2015, 12, 31, 23, 59, 59, 0, berlinTimeZone),
			ExpectedResult: "2015-12-31T23:59:59+01:00",
		},

		{
			Date:           time.Date(2015, 1, 1, 0, 0, 1, 0, berlinTimeZone),
			ExpectedResult: "2015-01-01T00:00:01+01:00",
		},

		{
			Date:           time.Date(2015, 1, 1, 0, 0, 1, 0, newYorkTimeZone),
			ExpectedResult: "2015-01-01T00:00:01-05:00",
		},
	}

	dateFormatter := iso80601Formatter{}

	for _, input := range inputs {

		// act
		result := dateFormatter.GetDateString(input.Date)

		// assert
		if result != input.ExpectedResult {
			t.Fail()
			t.Logf("GetDateString(%q) should have returned %q but returned %q instead.", input.Date, input.ExpectedResult, result)
		}
	}

}

func Test_GetDate_InvalidISO8601DatesGiven_ErrorIsReturned(t *testing.T) {
	// arrange
	inputs := []string{
		"Mon Jan _2 15:04:05 2006",
		"02 Jan 06 15:04 MST",
		"Jan _2 15:04:05",
		"2015-31-31T00:00:01+01:00",
		"2015-12-31T25:00:01+01:00",
		"2015-12-31T23:00:01 UTC",
		"2015-12-31T23:00:01",
	}

	dateFormatter := iso80601Formatter{}

	for _, input := range inputs {

		// act
		_, err := dateFormatter.GetDate(input)

		// assert
		if err == nil {
			t.Fail()
			t.Logf("GetDate(%q) should have returned an error because it is not a valid ISO8601 date.", input)
		}
	}

}

func Test_GetDate_ValidISO8601DatesGiven_DatesAreParsed(t *testing.T) {
	// arrange
	berlinTimeZone, _ := time.LoadLocation("Europe/Berlin")
	newYorkTimeZone, _ := time.LoadLocation("America/New_York")

	inputs := []struct {
		DateString     string
		ExpectedResult time.Time
	}{
		{
			DateString:     "2015-03-27T07:42:35+00:00",
			ExpectedResult: time.Date(2015, 3, 27, 7, 42, 35, 0, time.UTC),
		},

		{
			DateString:     "2015-01-01T00:00:01+00:00",
			ExpectedResult: time.Date(2015, 1, 1, 0, 0, 1, 0, time.UTC),
		},

		{
			DateString:     "2015-12-31T23:59:59+01:00",
			ExpectedResult: time.Date(2015, 12, 31, 23, 59, 59, 0, berlinTimeZone),
		},

		{
			DateString:     "2015-01-01T00:00:01+01:00",
			ExpectedResult: time.Date(2015, 1, 1, 0, 0, 1, 0, berlinTimeZone),
		},

		{
			DateString:     "2015-01-01T00:00:01-05:00",
			ExpectedResult: time.Date(2015, 1, 1, 0, 0, 1, 0, newYorkTimeZone),
		},
	}

	dateFormatter := iso80601Formatter{}

	for _, input := range inputs {

		// act
		result, _ := dateFormatter.GetDate(input.DateString)

		// assert
		if result.Format(time.RFC3339) != input.ExpectedResult.Format(time.RFC3339) {
			t.Fail()
			t.Logf("GetDate(%q) should have returned %q but returned %q instead.", input.DateString, input.ExpectedResult, result)
		}
	}

}
