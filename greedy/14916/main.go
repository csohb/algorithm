package main

import "fmt"

//춘향이는 편의점 카운터에서 일한다.
//손님이 2원짜리와 5원짜리로만 거스름돈을 달라고 한다.
//2원짜리 동전과 5원짜리 동전은 무한정 많이 가지고 있다.
//동전의 개수가 최소가 되도록 거슬러 주어야 한다. 거스름돈이 n인 경우, 최소 동전의 개수가 몇 개인지 알려주는 프로그램을 작성하시오.
//예를 들어, 거스름돈이 15원이면 5원짜리 3개를,
//거스름돈이 14원이면 5원짜리 2개와 2원짜리 2개로 총 4개를,
//거스름돈이 13원이면 5원짜리 1개와 2원짜리 4개로 총 5개를 주어야 동전의 개수가 최소가 된다.

func returnMoney(customerMoney int) int {
	// 거스름돈이 5원 이상의 돈이면
	var coin5Count int
	var coin2Count int
	var isAvailable bool

	for customerMoney >= 0 {
		if customerMoney%5 == 0 {
			coin5Count = customerMoney / 5
			isAvailable = true
			break
		} else {
			customerMoney -= 2
			coin2Count++
		}
	}
	if isAvailable {
		return coin5Count + coin2Count
	} else {
		return -1
	}
}

func main() {
	count := returnMoney(3)
	fmt.Println("count:", count)
}
