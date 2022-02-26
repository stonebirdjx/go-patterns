// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: singleton_test.go
// @Date: 2022/2/22 22:39
// @Desc:
package singleton

import (
	"testing"
	"unsafe"
)

func TestGetHungerInstance(t *testing.T) {
	if uintptr(unsafe.Pointer(GetHungerInstance())) != uintptr(unsafe.Pointer(GetHungerInstance())) {
		t.Errorf("hunger instance is not single")
	}
}

func BenchmarkGetHungerInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if uintptr(unsafe.Pointer(GetHungerInstance())) != uintptr(unsafe.Pointer(GetHungerInstance())) {
			b.Errorf("hunger instance is not single")
		}
	}
}

func TestGetLazyInstance(t *testing.T) {
	if uintptr(unsafe.Pointer(GetLazyInstance())) != uintptr(unsafe.Pointer(GetLazyInstance())) {
		t.Errorf("lazy instance is not single")
	}
}

func BenchmarkGetLazyInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if uintptr(unsafe.Pointer(GetLazyInstance())) != uintptr(unsafe.Pointer(GetLazyInstance())) {
			b.Errorf("lazy instance is not single")
		}
	}
}
