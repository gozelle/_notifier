package _monitor

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	timer := NewNotifier()
	timer.SetMuteDuration(3 * time.Second)
	i := 0
	for {
		i++
		if i == 2 || i == 12 {
			timer.Trigger()
			fmt.Println("触发", time.Now().Unix())
		}
		if i == 9 || i == 19 {
			timer.Reset()
			fmt.Println("重置", time.Now().Unix())
			fmt.Printf("\n\n")
		}
		fmt.Println("ReTrigger:", timer.ReTrigger(), time.Now().Unix())
		time.Sleep(1 * time.Second)
		if i > 20 {
			break
		}
	}
}
