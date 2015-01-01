package timer

import (
	"fmt"
	"time"
)

const (
	NANOSECOND int = iota
	MICROSECOND
	MILLISECOND
	SECOND
	MINUTE
	HOUR
)

type Timer struct {
	begin, end time.Time
}

func New() *Timer {
	return new(Timer)
}

func (t *Timer) Start() {
	t.begin = time.Now()
}

func (t *Timer) End() {
	t.end = time.Now()
}

func (t *Timer) Duration() time.Duration {
	return t.end.Sub(t.begin)
}

func (t *Timer) Nanoseconds() int64 {
	return t.Duration().Nanoseconds()
}

func (t *Timer) Microseconds() int64 {
	return int64(t.Duration() / time.Microsecond)
}

func (t *Timer) Milliseconds() int64 {
	return int64(t.Duration() / time.Millisecond)
}

func (t *Timer) Seconds() float64 {
	return t.Duration().Seconds()
}

func (t *Timer) Minutes() float64 {
	return t.Duration().Minutes()
}

func (t *Timer) Hours() float64 {
	return t.Duration().Hours()
}

func contains(i int, a []int) bool {
	for _, e := range a {
		if e == i {
			return true
		}
	}
	return false
}

func compute(d int64, unit time.Duration) (result, rest int64) {
	u := int64(unit)
	result = d / u
	rest = d % u
	return
}

func (t *Timer) Count() map[int]int64 {
	out := make(map[int]int64)
	n := t.Nanoseconds()
	out[HOUR], n = compute(n, time.Hour)
	out[MINUTE], n = compute(n, time.Minute)
	out[SECOND], n = compute(n, time.Second)
	out[MILLISECOND], n = compute(n, time.Millisecond)
	out[MICROSECOND], out[NANOSECOND] = compute(n, time.Microsecond)
	return out
}

func (t *Timer) String() string {
	d := t.Duration()
	switch {
	case d < 0:
		return "<not ended>"
	case d < time.Microsecond:
		return fmt.Sprintf("%d ns", t.Nanoseconds())
	case d < time.Millisecond:
		return fmt.Sprintf("%d µs", t.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%d ms", t.Milliseconds())
	case d < time.Minute:
		return fmt.Sprintf("%6.3f s", t.Seconds())
	case d < time.Hour:
		n := int64(d.Seconds())
		return fmt.Sprintf("%d'%d''", n/60, n%60)
	default:
		n := int64(d.Seconds())
		return fmt.Sprintf("%d:%02d:%02d", n/3600, (n%3600)/60, (n%3600)%60)
	}

}

type FuncTimer struct {
	t *Timer
	f func()
}

func NewFunc(f func()) *FuncTimer {
	return &FuncTimer{New(), f}
}

func (t *FuncTimer) Exec() {
	t.t.Start()
	t.f()
	t.t.End()
}

func (t *FuncTimer) Duration() time.Duration { return t.t.Duration() }
func (t *FuncTimer) Nanoseconds() int64      { return t.t.Nanoseconds() }
func (t *FuncTimer) Microseconds() int64     { return t.t.Microseconds() }
func (t *FuncTimer) Milliseconds() int64     { return t.t.Milliseconds() }
func (t *FuncTimer) Seconds() float64        { return t.t.Seconds() }
func (t *FuncTimer) Minutes() float64        { return t.t.Minutes() }
func (t *FuncTimer) Hours() float64          { return t.t.Hours() }
func (t *FuncTimer) Count() map[int]int64    { return t.t.Count() }
func (t *FuncTimer) String() string          { return t.t.String() }
