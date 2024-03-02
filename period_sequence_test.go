package synchro

import (
	"testing"
	"time"

	"github.com/Code-Hex/synchro/tz"
	"github.com/google/go-cmp/cmp"
)

func TestPeriodical_Seq(t *testing.T) {
	want := []Time[tz.UTC]{
		New[tz.UTC](2014, 2, 5, 0, 0, 0, 0),
		New[tz.UTC](2014, 2, 6, 0, 0, 0, 0),
		New[tz.UTC](2014, 2, 7, 0, 0, 0, 0),
		New[tz.UTC](2014, 2, 8, 0, 0, 0, 0),
	}

	t.Run("Time[UTC] params", func(t *testing.T) {
		period, err := NewPeriod[tz.UTC](
			New[tz.UTC](2014, 2, 5, 0, 0, 0, 0),
			New[tz.UTC](2014, 2, 8, 0, 0, 0, 0),
		)
		if err != nil {
			t.Fatal(err)
		}
		sequence := period.PeriodicDuration(24 * time.Hour).Seq()
		got := make([]Time[tz.UTC], 0, len(want))
		sequence(func(t Time[tz.UTC]) bool {
			got = append(got, t)
			return true
		})

		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("(-want, +got)\n%s", diff)
		}
	})
}
