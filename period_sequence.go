package synchro

import (
	"time"

	"github.com/nobishino/gocoro/iter"
)

// Seq returns iterator for the periodical sequence.
func (p periodical[T]) Seq() iter.Seq[Time[T]] {
	return func(yield func(Time[T]) bool) {
		defer p.markDone()
		for current := range p {
			if !yield(current) {
				break
			}
		}
	}
}

// Sequence returns iterator that yields Time[T] values at regular intervals
// between the start and end times of the Period[T]. The interval is specified
// by the next function argument.
//
// If start < end, the process will increase from start to end. In other words,
// when the current value exceeds end, the iteration is terminated.
func (p Period[T]) Sequence(next func(Time[T]) Time[T]) iter.Seq[Time[T]] {
	compare := isNotAfter[T]
	// p.start > p.end
	if p.from.After(p.to) {
		compare = isNotBefore[T]
	}

	return func(yield func(Time[T]) bool) {
		for current := p.from; compare(current, p.to); current = next(current) {
			if !yield(current) {
				break
			}
		}
	}
}

// SequenceDuration is a wrapper for the Sequence function.
// The interval is specified by the time.Duration argument.
func (p Period[T]) SequenceDuration(d time.Duration) iter.Seq[Time[T]] {
	return p.Sequence(func(t Time[T]) Time[T] {
		return t.Add(d)
	})
}
