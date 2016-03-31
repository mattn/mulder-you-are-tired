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
		sm := mulder{}
		for _, s := range *m.Take(2) {
			sm = append(sm, s)
		}

		ss := mulder{}
		for _, s := range sm {
			rs := []rune(s)
			rw := int(math.Ceil(float64(len(rs)) / 2))
			ss = append(ss, string(rs[:rw]))
			ss = append(ss, string(rs[rw:]))
		}
		rs := *(&ss).Take(4)
		if rand.Int()%2 == 0 {
			rs[0], rs[1], rs[2], rs[3] = rs[2], rs[1], rs[0], rs[3]
		}
		if r == "" {
			r += rs[0] + rs[1] + "、" + rs[2] + rs[3]
		} else {
			r += rs[0] + rs[1] + rs[2] + rs[3]
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
