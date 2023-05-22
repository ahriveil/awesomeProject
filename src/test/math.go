package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1

	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	stringNums := []string{"hello", "world"}
	for i, s := range stringNums {
		fmt.Println(i, s)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位的值是%d\n", i, x)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {

		fmt.Printf("map中index = %d的值为%d\n", key, value)
	}

	for key := range map1 {

		fmt.Println("map index =", key)
	}

	for _, value := range map1 {
		fmt.Println("map value =", value)
	}

	var i, j int

	for i = 2; i < 100; i++ {
	lo:
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break lo
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	i = 0
	for ; i < 20; i++ {
		if i%5 == 0 {
			continue
		}
		fmt.Println("i=", i)
	}

	var a int = 10

LOOP:
	for a <= 20 {
		if a == 15 {
			a++
			goto LOOP
		}
		fmt.Println("循环:", a)
		a++
	}
	fmt.Println("比较大小:", max(1, 2))
	var args = []int{1, 2, 3, 4, 5, 6}
	var result = avg(args)
	fmt.Println("平均值是:", result)

	fmt.Println("获取指针数组-----")
	getResult()
	str := "maps://public_default/works/cover/7192efa1-b49e-47cc-a997-04f3d50fb9ea.jpg"
	// (?<=/)[0-9a-zA-Z-]+(?=.*)
	rs := []rune(str)
	re := string(rs[strings.LastIndex(str, "/")+1 : strings.LastIndex(str, ".")])
	fmt.Printf("结果:%s\n", re)
	//compileRegex := regexp.MustCompile("[^/]+([0-9a-zA-Z-].*)")
	//matchArr := compileRegex.FindStringSubmatch(str)
	//fmt.Println("提取字符串内容：", matchArr)
	source := ParseWorkId("1_2_3")
	testALitter()
	fmt.Println(source)

	intersect := Intersect([]int{1, 2, 3, 4, 5}, []int{2, 5})
	fmt.Println("result:", intersect)
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	} else {
		return v2
	}
}

func avg(args []int) float32 {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return float32(sum / len(args))
}

const MAX int = 3

func getResult() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func ParseWorkId(modelId string) (source int64) {
	sa := strings.Split(modelId, "_")
	if len(sa) != 3 {
		return 0
	}
	v, err := strconv.ParseInt(sa[len(sa)-1], 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func TestSplit() (newSlice []int) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//var result = append(slice[0:0], slice[before:after]..., slice[:after]..., slice[after:]...)
	fmt.Println("result:", slice[:0])
	before := 0
	after := 4
	var result []int
	result = append(result, slice[after-1])
	result = append(result, slice[before:after-1]...)
	result = append(result, slice[after:]...)

	fmt.Println("result:", result)
	return
}

func TestSplit1() (newSlice []int) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	forms := []int{3, 6, 8}
	formsMap := map[int]int{}
	for index, data := range forms {
		formsMap[data] = index
	}
	innerMap := map[int]int{}
	for index, data := range slice {
		innerMap[data] = index
	}
	var top []int
	for _, data := range forms {
		if innerMap[data] == 0 {
			continue
		}
		top = append(top, data)
	}
	for _, data := range slice {
		if formsMap[data] == 0 {
			continue
		}

	}
	return
}

type TestObject struct {
	// count
	Count int64
}

func testALitter() {
	slice := []*TestObject{{Count: 1}, {Count: 2}}
	for i, _ := range slice {
		slice[i] = nil
	}
	fmt.Println("result:", slice)
}

func Intersect(nums1 []int, nums2 []int) []int {
	var result []int
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	for _, inner1 := range nums1 {
		for _, inner2 := range nums2 {
			if inner1 == inner2 {
				result = append(result, inner2)
			}
		}
	}
	return result
}
