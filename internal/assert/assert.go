package assert

import (
	"math"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()

	if !condition {
		tb.Logf(Red(msg), v...)
		tb.FailNow()
	}
}

func Fail(tb testing.TB, msg string, v ...interface{}) {
	tb.Helper()
	tb.Logf(Red(msg), v...)
	tb.FailNow()
}

func Ok(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Logf(Red("unexpected error: %v"), err.Error())
		tb.FailNow()
	}
}

func Equals(tb testing.TB, exp, act interface{}, opts ...cmp.Option) {
	tb.Helper()
	EqualsWithMsg(tb, "", exp, act, opts...)
}

func EqualsWithMsg(tb testing.TB, message string, exp, act interface{}, opts ...cmp.Option) {
	tb.Helper()

	opts = append(opts, cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		return delta < 0.001
	}))

	diff := cmp.Diff(exp, act, opts...)

	if diff != "" {

		var output strings.Builder

		if message != "" {
			output.WriteString(Yellow(message))
			output.WriteString("\n")
		}

		output.WriteString("mismatch (")
		output.WriteString(Red("-want"))
		output.WriteString(Green(" +got"))
		output.WriteString("):\n")

		output.WriteString(colorizeDiff(diff))

		tb.Error(output.String())
		tb.FailNow()
	}
}

func colorizeDiff(diff string) string {
	var coloredDiff strings.Builder
	for _, line := range strings.Split(diff, "\n") {
		if len(line) > 0 {
			switch line[0] {
			case '-':
				coloredDiff.WriteString(Red(line))
			case '+':
				coloredDiff.WriteString(Green(line))
			default:
				coloredDiff.WriteString(line)
			}
		}
		coloredDiff.WriteString("\n")
	}

	return coloredDiff.String()
}

func Red(text string) string {
	return "\033[31m" + text + "\033[0m"
}

func Green(text string) string {
	return "\033[32m" + text + "\033[0m"
}

func Yellow(text string) string {
	return "\033[33m" + text + "\033[0m"
}
