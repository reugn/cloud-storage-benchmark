package cli

import "github.com/schollz/progressbar/v3"

type ProgressBar interface {
	Add(int)
	Clear()
}

type DefaultProgressBar struct {
	delegate *progressbar.ProgressBar
}

var _ ProgressBar = (*DefaultProgressBar)(nil)

func NewDefaultProgressBar(max int, description string) *DefaultProgressBar {
	return &DefaultProgressBar{
		delegate: progressbar.Default(int64(max), description),
	}
}

func (p *DefaultProgressBar) Add(i int) {
	_ = p.delegate.Add(i)
}

func (p *DefaultProgressBar) Clear() {
	_ = p.delegate.Clear()
}

type NoOpProgressBar struct {
}

var _ ProgressBar = (*NoOpProgressBar)(nil)

func NewNoProgressBar() *NoOpProgressBar {
	return &NoOpProgressBar{}
}

func (p *NoOpProgressBar) Add(_ int) {
	// no-op
}

func (p *NoOpProgressBar) Clear() {
	// no-op
}
