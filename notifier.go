package _monitor

import (
	"time"
)

func NewNotifier() *Notifier {
	return &Notifier{}
}

// Notifier 监控通知器
type Notifier struct {
	triggeredAt  *time.Time
	reTriggerAt  *time.Time
	muteDuration time.Duration
}

// SetMuteDuration 设置沉默时间，在此时间段内，调用 ReTrigger() 将得到 false
func (p *Notifier) SetMuteDuration(muteDuration time.Duration) {
	p.muteDuration = muteDuration
}

// Trigger 触发通知
func (p *Notifier) Trigger() {
	now := time.Now()
	p.triggeredAt = &now
}

// Reset 重置通知器
func (p *Notifier) Reset() {
	p.triggeredAt = nil
	p.reTriggerAt = nil
}

// ReTrigger 获取是否再次触发
// 获取再次触发后，静默周期将从头计算。
// 后续的触发动作有必要保证触发成功，否则将在下个周期获得重新触发的机会
func (p *Notifier) ReTrigger() bool {
	if p.triggeredAt == nil {
		return false
	}
	if p.reTriggerAt == nil {
		p.reTriggerAt = p.triggeredAt
	}
	if p.muteDuration == 0 || p.reTriggerAt == nil {
		return false
	}
	if uint64(time.Since(*p.reTriggerAt)) > uint64(p.muteDuration) {
		now := time.Now()
		p.reTriggerAt = &now
		return true
	}
	return false
}
