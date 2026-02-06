package test

import (
	"sync"
	"testing"
	"time"

	test_tools "github.com/ziv/sonder/tools/test-tools"
)

var (
	instance *Instance
	once     sync.Once
)

type Instance struct {
	ID   string
	Name string
}

func init1() *Instance {
	once.Do(func() {
		instance = &Instance{
			ID:   "111",
			Name: "111",
		}
	})
	return instance
}

func init2() *Instance {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		instance = &Instance{
			ID:   "111",
			Name: "111",
		}
	})
	return instance
}

func BenchmarkInit1(b *testing.B) {
	defer test_tools.TimeTrack(time.Now())
	for i := 0; i < b.N; i++ {
		init1()
	}
}

func BenchmarkInit2(b *testing.B) {
	defer test_tools.TimeTrack(time.Now())
	for i := 0; i < b.N; i++ {
		init2()
	}
}
