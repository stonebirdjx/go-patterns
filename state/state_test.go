// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: state_test.go
// @Date: 2022/2/28 21:40
// @Desc: 状态模式测试方法
package state

import "testing"

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	if m.GetStateName() != "LeaderApproveState" {
		t.Fatal()
	}

	m.Approval()
	if m.GetStateName() != "FinanceApproveState" {
		t.Fatal()
	}

	m.Reject()
	if m.GetStateName() != "LeaderApproveState" {
		t.Fatal()
	}

	m.Approval()
	if m.GetStateName() != "FinanceApproveState" {
		t.Fatal()
	}
}
