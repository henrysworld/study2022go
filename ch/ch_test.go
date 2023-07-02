package ch

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	//给nums 和nums1合并并排序
	nums := []int{1, 2, 3, 4}
	nums1 := []int{1, 2, 3, 4}
	nums3 := make([]int, 0)
	length1 := len(nums)
	length2 := len(nums1)
	var lenx int
	var max int
	if length1 < length2 {
		lenx = length1
		max = length2
	} else {
		lenx = length2
		max = length1
	}

	for i := 0; i < lenx; i++ {
		if nums[i] < nums1[i] {
			nums3 = append(nums3, nums[i])
		} else if nums[i] > nums1[i] {
			nums3 = append(nums3, nums1[i])
		} else {
			nums3 = append(nums3, nums[i])
			nums3 = append(nums3, nums1[i])
		}
	}

	i := lenx
	for i < max {
		nums3 = append(nums3, nums[i])
		i++
	}

	//nums3 = append(nums, nums1...)

	//sort.Ints(nums3)
	sort.Ints(nums3)
	fmt.Printf("%v", nums3)

}

func TestFor(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 1024)
			runtime.Stack(buf, false)
			//fmt.Printf("Trace:\n%s\n", buf)
		}
	}()

	filterEventMp := make([]map[string]string, 0)
	m1 := make(map[string]string, 3)
	m1["event_name"] = "e1"
	m1["proc_name"] = "p1"
	m1["message"] = "m1"
	m1["other"] = "o1"

	m2 := make(map[string]string, 3)
	m2["event_name"] = "e2"
	m2["proc_name"] = "p2"
	m2["message"] = "m2"
	m2["other"] = "o2"

	filterEventMp = append(filterEventMp, m1)
	filterEventMp = append(filterEventMp, m2)

	//s := make([]int, 0, 0)
	//s = append(s, 1)
	ch := make(chan struct{}, 3)
	close(ch)
	ch <- struct{}{}

	evt := make([]*EventTbl, 0)
	e1 := &EventTbl{
		EventId:   1,
		EventName: "e1",
		ProcName:  "p1",
		Message:   "m1",
	}

	e2 := &EventTbl{
		EventId:   2,
		EventName: "e2",
		ProcName:  "p2",
		Message:   "m2",
	}

	evt = append(evt, e1)
	evt = append(evt, e2)

}

// please help me
func isFilterEvent(ctx context.Context, evt []*EventTbl, filterEvtMp []map[string]string) bool {
	for _, v := range evt {
		for _, f := range filterEvtMp {
			isFilter := true
			for key, val := range f {
				if key == "event_name" {
					if v.EventName != val {
						isFilter = false
						break
					}
				} else if key == "proc_name" {
					if v.ProcName != val {
						isFilter = false
						break
					}
				} else if key == "message" {
					if v.Message != val {
						isFilter = false
						break
					}
				} else {
					isFilter = false
					break
				}
			}

			if isFilter {
				return true
			}
		}

	}

	return false
}

type EventTbl struct {
	EventId   int32  `json:"event_id"`
	EventName string `json:"event_name"`
	ProcName  string `json:"proc_name"`
	Message   string `json:"message"`
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

//

type Article struct {
	Title       string `json:"title"`
	StoryTitle  string `json:"story_title"`
	NumComments int    `json:"num_comments"`
}

func sortByNumCommentsAndTitle(slice []Article) {
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].NumComments < slice[j].NumComments {
			return true
		} else if slice[i].NumComments > slice[j].NumComments {
			return false
		}

		// 若 NumComments 相等，则根据 Title 进行升序排序
		return slice[i].Title < slice[j].Title
	})
}

func TestArticle(t *testing.T) {
	slice := []Article{
		{Title: "Title C", StoryTitle: "Story C", NumComments: 2},
		{Title: "Title A", StoryTitle: "Story A", NumComments: 1},
		{Title: "Title B", StoryTitle: "Story B", NumComments: 2},
		{Title: "Title D", StoryTitle: "Story D", NumComments: 3},
	}

	sortByNumCommentsAndTitle1(slice)

	for _, article := range slice {
		fmt.Printf("Title: %s, NumComments: %d\n", article.Title, article.NumComments)
	}
}

func sortByNumCommentsAndTitle1(article []Article) {
	sort.Slice(article, func(i, j int) bool {
		if article[i].NumComments < article[j].NumComments {
			return true
		} else if article[i].NumComments == article[j].NumComments {
			if article[i].NumComments <= article[j].NumComments {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})
}

//// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
//func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	// NOTE:
//	// Do not move the code below to a goroutine.
//	// The `ConsumeClaim` itself is called within a goroutine, see:
//	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
//	for {
//		select {
//		case message := <-claim.Messages():
//
//			consumer.token <- true
//			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
//			go func() {
//				<-consumer.token
//				handler()
//			}()
//			session.MarkMessage(message, "")
//
//		// Should return when `session.Context()` is done.
//		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
//		// https://github.com/Shopify/sarama/issues/1192
//		case <-session.Context().Done():
//			return nil
//		}
//	}
//}

func TestMap(t *testing.T) {

}

func testmap1(mp sync.Map) {

}
