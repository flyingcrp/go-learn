package event

import (
	"context"
	"log/slog"
	"sync"
)

type Event struct {
	Type    string // "user.registered", "user.updated"
	Payload any    // 携带的数据
}

type Bus struct {
	mu   sync.RWMutex
	subs []chan Event
}

func NewBus() *Bus {
	return &Bus{}
}

// Publish 非阻塞地向所有订阅者发送事件。如果订阅者的 channel 满了就丢弃。
func (b *Bus) Publish(ctx context.Context, e Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subs {
		select {
		case ch <- e:
			// 发送成功
		case <-ctx.Done():
			return
		default:
			// channel 满了，丢弃（生产环境可以加 metrics）
			slog.Warn("event bus: subscriber queue full, dropped", "type", e.Type)
		}
	}
}

// Subscribe 返回一个只读 channel，消费者用 for-select 循环读。
// ctx 取消时，自动清理订阅（关闭 channel + 从列表移除）。
func (b *Bus) Subscribe(ctx context.Context, bufSize int) <-chan Event {
	ch := make(chan Event, bufSize)

	b.mu.Lock()
	b.subs = append(b.subs, ch)
	b.mu.Unlock()

	// ctx 取消时，移除这个订阅
	go func() {
		<-ctx.Done()
		b.mu.Lock()
		for i, sub := range b.subs {
			if sub == ch {
				b.subs = append(b.subs[:i], b.subs[i+1:]...)
				break
			}
		}
		b.mu.Unlock()
		close(ch)
	}()

	return ch
}
