package _monitor

import (
	"time"
)

func NewNotifier() *Notifier {
	return &Notifier{}
}

// Notifier 监控通知器
type Notifier struct {
	triggeredAt    *time.Time
	reTriggerAt    *time.Time
	remindDuration time.Duration
}

// SetRemindDuration 设置提醒周期，在此时间段内，调用 Remind() 将得到 false
func (p *Notifier) SetRemindDuration(d time.Duration) *Notifier {
	p.remindDuration = d
	return p
}

// FromNow 获取首次触发到现在的时间间隔
func (p *Notifier) FromNow() time.Duration {
	if p.triggeredAt == nil {
		return 0
	}
	return time.Since(*p.triggeredAt)
}

// Sent 标记通知已发送状态
func (p *Notifier) Sent() {
	now := time.Now()
	p.triggeredAt = &now
}

// Empty 判断通知器是否为空状态（未触发过的状态）
func (p *Notifier) Empty() bool {
	return p.triggeredAt != nil
}

// Clear 重置通知器状态
func (p *Notifier) Clear() {
	p.triggeredAt = nil
	p.reTriggerAt = nil
}

// Remind 获取是否到达发送提醒的时间点
// 获取再次触发后，静默周期将从头计算。
// 后续的触发动作有必要保证触发成功，否则将在下个周期获得重新触发的机会
func (p *Notifier) Remind() bool {
	if p.triggeredAt == nil {
		return false
	}
	if p.reTriggerAt == nil {
		p.reTriggerAt = p.triggeredAt
	}
	if p.remindDuration == 0 || p.reTriggerAt == nil {
		return false
	}
	if uint64(time.Since(*p.reTriggerAt)) > uint64(p.remindDuration) {
		now := time.Now()
		p.reTriggerAt = &now
		return true
	}
	return false
}
