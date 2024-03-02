package synchro

import "github.com/nobishino/gocoro/iter"

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
