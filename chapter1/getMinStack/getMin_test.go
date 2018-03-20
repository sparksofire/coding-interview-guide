package getMinStack

import (
	"math/rand"
	"sort"
	"testing"
	"time"
	"github.com/shawpo/coding-interview-guide/chapter1/getMinStack/stack1"
	"github.com/shawpo/coding-interview-guide/chapter1/getMinStack/stack2"
)

func TestStack(t *testing.T) {
	t.Log("具有getMin 功能的栈表组测试")
	var tests = []struct {
		randAmount int  // 随机数数量
		randMax     int  // 随机数最大值
	} {
			{
				30,
				100,
			},
			{
				50,
				200,
			},
			{
				500,
				1000,
			},
	}
	for i, test := range tests {
		t.Logf("正在进行第%d组测试, 共%d组", i+1, len(tests))
		stack_1 := stack1.New() // 方案一创建空栈
		stack_2 := stack2.New() // 方案二创建空栈
		rand.Seed(time.Now().UnixNano())
		t.Logf("生成一组随机数，并依次插入栈中，随机数个数为 %d，随机数最大值为 %d",
			test.randAmount, test.randMax)
		var randNums []int // 存放随机数的切片
		for i := 0; i < test.randAmount; i++ {
			randNum := rand.Intn(test.randMax)
			randNums = append(randNums, randNum)
			stack_1.Push(randNum)
			stack_2.Push(randNum)
		}
		t.Logf("随机数插入顺序为：%v", randNums)
		sort.Ints(randNums)
		t.Logf("随机数从小到大排列为：%v", randNums)
		min := randNums[0] // 随机数的最小值
		getMin1, ok1 := stack_1.GetMin()
		getMin2, ok2 := stack_2.GetMin()
		if !ok1 {
			t.Fatal("栈无数据！")
		} else if getMin1 != min {
			t.Fatalf("方案一栈GetMin()返回值 %d 与随机数最小值 %d不匹配",
				getMin1, min)
		} else {
			t.Logf("方案一栈GetMin()返回值 %d 与随机数最小值 %d匹配",
				getMin1, min)
		}

		if !ok2 {
			t.Fatal("栈无数据！")
		} else if getMin2 != min {
			t.Fatalf("方案二栈GetMin()返回值 %d 与随机数最小值 %d不匹配",
				getMin2, min)
		} else {
			t.Logf("方案二栈GetMin()返回值 %d 与随机数最小值 %d匹配",
				getMin2, min)
		}
	}

}
