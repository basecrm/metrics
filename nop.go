package metrics

import (
	"time"
)

// nopclient does nothing when called, useful in tests
type nopclient struct {
}

func (s *nopclient) Close() error {
	return nil
}

func (s *nopclient) Inc(stat string, value int64, rate float32) error {
	return nil
}

func (s *nopclient) Dec(stat string, value int64, rate float32) error {
	return nil
}

func (s *nopclient) Gauge(stat string, value int64, rate float32) error {
	return nil
}

func (s *nopclient) GaugeDelta(stat string, value int64, rate float32) error {
	return nil
}

func (s *nopclient) Timing(stat string, delta int64, rate float32) error {
	return nil
}

func (s *nopclient) UniqueString(stat string, value string, rate float32) error {
	return nil
}

func (s *nopclient) UniqueInt64(stat string, value int64, rate float32) error {
	return nil
}

func (s *nopclient) SetPrefix(prefix string) {

}

func (s *nopclient) TimingMs(stat string, d time.Duration, rate float32) error {
	return nil
}

func (s *nopclient) ReportRuntimeMetrics(prefix string, rate float32) error {
	return nil
}

func NewNop() Client {
	return &nopclient{}
}
