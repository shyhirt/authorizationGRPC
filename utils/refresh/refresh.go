package refresh

import (
	"math/rand"
	"sync"
	"time"
)

type Session struct {
	Refresh string
	Expired time.Time
}
type Refresh struct {
	users             map[int64]Session
	mut               sync.Mutex
	expirationSeconds int64
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
func NewRefresh(expirationSeconds int64) *Refresh {
	return &Refresh{
		users:             make(map[int64]Session),
		mut:               sync.Mutex{},
		expirationSeconds: expirationSeconds,
	}
}

func (r *Refresh) Set(userId int64) Session {
	expired := time.Now().Add(time.Duration(r.expirationSeconds) * time.Second)
	refresh := RandStringBytesMask(32)
	session := Session{
		Refresh: refresh,
		Expired: expired,
	}
	r.mut.Lock()
	r.users[userId] = session
	r.mut.Unlock()
	return session
}

func (r *Refresh) Verify(userId int64, refresh string) bool {
	session := Session{}
	r.mut.Lock()
	session = r.users[userId]
	r.mut.Unlock()
	return refresh == session.Refresh && session.Expired.After(time.Now())
}
