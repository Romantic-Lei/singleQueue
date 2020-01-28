package main
import (
	"fmt"
	"os"
	"errors"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array [5]int // 数组 =》 模拟队列
	front int // 表示指向队首,但不含队首
	rear int // 表示指向队尾
}

// 提供一个添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	// 设计时将 rear 设计成队列尾，且含队列的最后一个元素
	if this.rear == this.maxSize -1 {
		return errors.New("Queue full")
	}
	this.rear++ // rear 后移
	this.array[this.rear] = val
	return 
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判读是否为空
	if this.rear == this.front {
		return -1, errors.New("Queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, nil
}

// 显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("当前队列为：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\t", i, this.array[i])
	}
	fmt.Println()
}

// 主函数
func main() {
	// 先创建一个队列, 并给其赋值
	queue := &Queue{
		maxSize : 5,
		front : -1,
		rear : -1,
	}

	var key string 
	var val int 
	for {
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 从队列中获取数据")
		fmt.Println("3. 输入show 获取队列的信息")
		fmt.Println("4. 输入exit 退出程序")

		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("输入需要入队列的数据")
				fmt.Scanln(&val)
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("加入队列成功")
				}
			case "get": 
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("取出的值为：", val)
				}
			case "show": 
				queue.ShowQueue()
			case "exit" :
				os.Exit(0)
		}
	}
}