// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: strategy_test.go
// @Date: 2022/2/28 21:00
// @Desc: 策略模式测试方法
package strategy

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func ExamplePayByBank() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}
