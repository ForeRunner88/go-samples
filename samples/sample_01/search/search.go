package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个无缓冲的通道 接受匹配后的结果
	results := make(chan *Result)
	// 构造一个waitGroup处理所有的数据源
	var waitGroup sync.WaitGroup
	// 设置需要等待处理每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))
	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// 启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	// 启动一个goroutine来监控是否所有的工作已经结束
	go func() {
		// 阻塞等待
		waitGroup.Wait()
		// 用关闭通道的方式通知Display函数退出
		close(results)
	}()
	// Display函数显示返回的结果并在results关闭后返回
	Display(results)
}

// Register函数将Matcher匹配器注册到全局map变量matchers中
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
