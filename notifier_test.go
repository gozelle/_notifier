package _notifier

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	notifier := NewNotifier()
	notifier.SetRemindDuration(3 * time.Second)
	notifier.SetCallbacks(
		func(state *State) {
			t.Log("告警了")
		},
		func(state *State) {
			t.Log("提醒了")
		},
		func(state *State) {
			t.Log("修复了")
		})
	i := 0
	for {
		i++
		t.Logf("第 %d 秒", i)
		notifier.Check(func() bool {
			return i != 9 && i != 19
		})
		time.Sleep(1 * time.Second)
		if i > 20 {
			break
		}
	}
}
