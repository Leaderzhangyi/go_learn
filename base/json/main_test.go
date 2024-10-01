package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

// 磁盘与网络信息传输，本质是一个二进制流

var (
	s = Student{"Tom", 18, "Male"}
	c = Class{
		Id:       "1001",
		Students: []Student{s, s, s},
	}
)

func TestJson(t *testing.T) {
	bytes, err := json.Marshal(c)
	if err != nil {
		assert.NoError(t, err)
	}
	var c2 Class
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, c, c2)

}

func TestSonic(t *testing.T) {
	bytes, err := sonic.Marshal(c)
	if err != nil {
		assert.NoError(t, err)
	}
	var c2 Class
	err = sonic.Unmarshal(bytes, &c2)
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, c, c2)
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)
		var c2 Class
		_ = sonic.Unmarshal(bytes, &c2)
	}
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)
		var c2 Class
		_ = json.Unmarshal(bytes, &c2)
	}
}
