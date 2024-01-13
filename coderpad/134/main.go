package main

import (
	"fmt"
	"log"
)

/*
페이스북에서 제출된 문제입니다.

0이 대다수를 차지하는 큰 배열이 있습니다.

더 공간 효율적인 자료구조인 SparseArray를 다음과 같은 인터페이스로 구현하세요.

init(arr, size): 큰 원본 배열과 사이즈를 인자로 받아 초기화 합니다.
set(i, val): 인덱스 i 를 val 값으로 업데이트 합니다.
get(i): 인덱스 i 번째 값을 반환합니다.
*/

type SparseArray interface {
	init(array []int, size int)
	set(index int, value int)
	get(index int) (value int)
}

type AlmostZeroArray struct {
	size       int
	arrayValue []int
}

func (a *AlmostZeroArray) init(array []int, size int) {
	a.size = size
	a.arrayValue = array
}

func (a *AlmostZeroArray) set(index int, value int) {
	if len(a.arrayValue)-1 < index {
		log.Fatal("given index is out of length of array")
	} else {
		a.arrayValue[index] = value
	}
}

func (a *AlmostZeroArray) get(index int) (value int) {
	if len(a.arrayValue)-1 < index {
		log.Fatal("given index is out of length of array")
		return -1
	} else {
		return a.arrayValue[index]
	}
}

func main() {
	var givenArray []int
	givenArray = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 5}

	zeroArray := &AlmostZeroArray{}

	zeroArray.init(givenArray, len(givenArray))
	fmt.Println("zeroArray:", zeroArray)

	zeroArray.set(5, 10)
	fmt.Println("zeroArray:", zeroArray)

	value := zeroArray.get(14)
	fmt.Println("index value", value)

}
