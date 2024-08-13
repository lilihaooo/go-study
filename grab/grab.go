package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var workerSignalCh = make(chan struct{}, 10)
var cancelSignalCh = make(chan struct{})

var toMysqlJobCh = make(chan User)
var wg sync.WaitGroup

func grabCancel() {
	close(cancelSignalCh)
}

func checkIsCancel() bool {
	select {
	case <-cancelSignalCh:
		return true
	default:
		return false
	}
}

func main() {
	begin := time.Now()
	go GrabTask() // 耗时任务

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	grabCancel()
	//}()

	ToMysqlJob() // 任务全部结束后关闭通道， 遍历完缓冲释放阻塞

	elapsed := time.Since(begin)
	fmt.Println("解除阻塞")
	// 关闭全局锁

	fmt.Println(elapsed)
	time.Sleep(1 * time.Second)

}

// 下发工作
func GrabTask() {
	defer close(toMysqlJobCh) // 任务全部执行完成后， 关闭发送mysql信息的通道
	defer fmt.Println(333)
	categories := []string{
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
		"食品",
		"服装",
		"电器",
		"家居用品",
		"美容化妆",
		"电子产品",
		"运动器材",
		"书籍",
		"文具用品",
		"玩具",
		"母婴用品",
		"汽车配件",
		"珠宝配饰",
		"手工艺品",
		"旅游用品",
		"健康保健",
		"宠物用品",
		"园艺工具",
		"办公用品",
		"户外装备",
		"厨具餐具",
		"数码产品",
		"箱包皮具",
		"鞋靴",
		"音乐乐器",
		"艺术收藏品",
		"家居装饰",
		"家具",
		"家用电器",
		"清洁用品",
		"食品饮料",
		"游戏娱乐",
	}
	for _, cate := range categories {
		isCancel := checkIsCancel()
		fmt.Println(isCancel)
		if isCancel {
			return
		}
		workerSignalCh <- struct{}{}
		wg.Add(1)
		go work(cate)
	}

	// 等所以的任务全部执行完成后在return
	wg.Wait()
	return
}

// 爬取工作
func work(cate string) {
	defer func() {
		<-workerSignalCh // 释放一个工人
		wg.Done()
	}()

	users := []User{}
	for i := 0; i < 10; i++ {
		users = append(users, User{ID: uint(i), Name: cate})
	}
	for _, user := range users {
		isCancel := checkIsCancel()
		if isCancel {
			return
		}
		toMysqlJobCh <- user
		time.Sleep(1 * time.Second)
		//

	}
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToMysqlJob() {
	//for {
	for user := range toMysqlJobCh {
		//db.Save(&user)
		fmt.Println("写入Mysql：" + user.Name + "id: " + strconv.Itoa(int(user.ID)))
	}
	//}
}
