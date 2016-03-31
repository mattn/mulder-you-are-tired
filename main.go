package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type mulder []string

func NewMulder() *mulder {
	return &mulder{"モルダー", "あなた", "疲れてる", "のよ"}
}

func (m *mulder) Shuffle() *mulder {
	n := len(*m)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	}
	return m
}

func (m *mulder) WorkUntil(f func()) {
	for len(*m) > 0 {
		f()
	}
}

func (m *mulder) Take(n int) *mulder {
	rn := n
	if len(*m) < rn {
		rn = len(*m)
	}
	r := (*m)[:rn]
	*m = (*m)[rn:]
	for len(r) < n {
		r = append(r, "")
	}
	return &r
}

func (m *mulder) Tired() (r string) {
	m.Shuffle().WorkUntil(func() {
		for _, s := range *m.Take(2) {
			t := ""
			if rand.Int()%2 == 0 {
				sm := mulder{}
				rs := []rune(s)
				for _, rc := range rs {
					sm = append(sm, string(rc))
				}
				rw := int(math.Ceil(float64(len(rs)) / 2))
				a := *sm.Take(rw)
				b := *sm.Take(rw)
				if rw == 1 {
					t += b[0] + a[0]
				} else {
					t += b[0] + a[1] + a[0] + b[1]
				}
			} else {
				t += s
			}
			if r == "" {
				t += "、"
			}
			r += t
		}
	})
	return r
}

func main() {
	i := 1
	for {
		s := NewMulder().Tired()
		fmt.Printf("%d: %s\n", i, s)
		if s == "モルダー、あなた疲れてるのよ" {
			break
		}
		i++
	}
}
