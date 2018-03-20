package twoStacksQueue

import (
	"math/rand"
	"testing"
	"time"
)

func TestStacksQueue(t *testing.T) {
	t.Log("由两个栈组成的队列表组测试")
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
	for group, test := range tests {
		t.Logf("正在进行第%d组测试, 共%d组", group+1, len(tests))
		q := New()
		rand.Seed(time.Now().UnixNano())
		t.Logf("生成一组随机数，并依次插入队列中，随机数个数为 %d，随机数最大值为 %d",
			test.randAmount, test.randMax)
		var randNums []int // 存放随机数的切片
		for i := 0; i < test.randAmount; i++ {
			randNum := rand.Intn(test.randMax)
			randNums = append(randNums, randNum)
			q.Add(randNum)
		}
		t.Logf("随机数插入顺序为：%v", randNums)
		for i, v := range randNums {
			peekValue, peekOk := q.Peek()
			pollValue, pollOk := q.Poll()
			if !peekOk {
				t.Fatalf("peek操作异常，队列可能为空")
			}

			if !pollOk {
				t.Fatalf("poll操作异常，队列可能为空")
			}

			if !(peekOk && peekValue == v) {
				t.Fatalf("第%d组>>第%d次>peek操作，peek值%d 与预期值%d不符", group+1, i+1, peekValue, v)
			} else {
				t.Logf("第%d组>>第%d次>peek操作，peek值%d 与预期值%d相符", group+1, i+1, peekValue, v)
			}

			if !(pollOk && pollValue == v) {
				t.Fatalf("第%d组>>第%d次>pool操作，poll值%d 与预期值%d不符", group+1, i+1, pollValue, v)
			} else {
				t.Logf("第%d组>>第%d次>pool操作，poll值%d 与预期值%d相符", group+1, i+1, pollValue, v)
			}
		}
	}

}
