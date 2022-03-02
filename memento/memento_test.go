// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: memento_test.go
// @Date: 2022/3/2 9:45
// @Desc: 备忘录模式:测试方法
package memento

import "testing"

func TestDemo(t *testing.T) {
	in := &InputText{}
	snapshots := []*Snapshot{}

	tests := []struct {
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
		{
			input: "world",
			want:  "",
		},
		{
			input: ":list",
			want:  "helloworld",
		},
		{
			input: ":undo",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			switch tt.input {
			case ":list":
				if tt.want != in.GetText() {
					t.Fatal()
				}
			case ":undo":
				in.Restore(snapshots[len(snapshots)-1])
				snapshots = snapshots[:len(snapshots)-1]
			default:
				snapshots = append(snapshots, in.Snapshot())
				in.Append(tt.input)
			}
		})
	}
}
