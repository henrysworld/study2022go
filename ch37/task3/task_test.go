package task3

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestTask3(t *testing.T) {
	tasks := NewTasks()
	task1 := tasks[0]
	task2 := tasks[1]

	//fmt.Println("CH===:", task1)
	//fmt.Println("CH===:", task2)
	fmt.Printf("CH===:%s===%v\n", task1.times, task1.temps)
	fmt.Printf("CH===:%s===%v\n", task2.times, task2.temps)

	taskByTime := make(map[string]struct{})

	tmpTimes1 := make([]string, 0, len(task1.times))
	tmpTemps1 := make([]float32, 0, len(task1.temps))
	tmpTimes2 := make([]string, 0, len(task2.times))
	tmpTemps2 := make([]float32, 0, len(task2.temps))
	for _, tm := range task1.times {
		taskByTime[tm] = struct{}{}
	}
	for _, tm := range task2.times {
		taskByTime[tm] = struct{}{}
	}

	//len1 := len(task1.times)
	//len2 := len(task2.times)
	//maxLen := math.Max(float64(len1), float64(len2))
	var p1, p2 int
	for p1 < len(task1.times) && p2 < len(task2.times) {
		time1 := task1.times[p1]
		time2 := task2.times[p2]
		temp1 := task1.temps[p1]
		temp2 := task2.temps[p2]

		if time1 == time2 {
			tmpTimes1 = append(tmpTimes1, time1)
			tmpTimes2 = append(tmpTimes2, time2)
			tmpTemps1 = append(tmpTemps1, temp1)
			tmpTemps2 = append(tmpTemps2, temp2)
			p1++
			p1++
		} else if time1 < time2 {
			tmpTimes1 = append(tmpTimes1, time1)
			tmpTimes2 = append(tmpTimes2, time1)
			tmpTemps1 = append(tmpTemps1, temp1)
			if (p2 - 1) < 0 {
				//如果p2为第一个，这里需要确认mcu1和mcu2的温度是否差不多。
				tmpTemps2 = append(tmpTemps2, temp2)
			} else {
				//如果p2的前一个值有值，则取一个符合线性规律的值。
				preX0 := Time2Unix(task2.times[p2-1])
				preY0 := task2.temps[p2-1]
				x1 := Time2Unix(time2)
				y1 := temp2
				x := Time2Unix(time1)

				tmpTemps2 = append(tmpTemps2, linearInterpolation(preX0, preY0, x1, y1, x))
			}

			p1++
		} else {
			tmpTimes1 = append(tmpTimes1, time2)
			tmpTimes2 = append(tmpTimes2, time2)
			tmpTemps2 = append(tmpTemps2, temp2)
			if (p1 - 1) < 0 {
				//如果p1为第一个，这里需要确认mcu1和mcu2的温度是否差不多。
				tmpTemps1 = append(tmpTemps1, temp1)
			} else {
				//如果p1的前一个值有值，则取一个符合线性规律的值。
				preX0 := Time2Unix(task1.times[p1-1])
				preY0 := task1.temps[p1-1]
				x1 := Time2Unix(time1)
				y1 := temp1
				x := Time2Unix(time2)

				tmpTemps1 = append(tmpTemps1, linearInterpolation(preX0, preY0, x1, y1, x))
			}

			p2++
		}
	}

	if p1 < len(task1.times) {
		//var tp2 int
		//if p2 >= len(task2.times) {
		//	tp2 = len(task2.times) - 1
		//}
		//这里task2肯定已经遍历完成
		tp2 := len(task2.times) - 1
		time1 := task1.times[p1]
		time2 := task2.times[tp2]
		temp1 := task1.temps[p1]
		temp2 := task2.temps[tp2]

		tmpTimes1 = append(tmpTimes1, time1)
		tmpTimes2 = append(tmpTimes2, time1)
		tmpTemps1 = append(tmpTemps1, temp1)
		if (p2 - 1) < 0 {
			//如果p2为第一个，这里需要确认mcu1和mcu2的温度是否差不多。
			tmpTemps2 = append(tmpTemps2, temp2)
		} else {
			//如果p2的前一个值有值，则取一个符合线性规律的值。
			preX0 := Time2Unix(task2.times[tp2-1])
			preY0 := task2.temps[tp2-1]
			x1 := Time2Unix(time2)
			y1 := temp2
			x := Time2Unix(time1)

			tmpTemps2 = append(tmpTemps2, linearInterpolation(preX0, preY0, x1, y1, x))
		}

		p1++
	}

	if p2 < len(task2.times) {
		//var tp1 int
		//if p1 >= len(task1.times) {
		//	tp1 = len(task1.times) - 1
		//}

		//这里task1肯定已经遍历完成
		tp1 := len(task1.times) - 1
		time1 := task1.times[tp1]
		time2 := task2.times[p2]
		temp1 := task1.temps[tp1]
		temp2 := task2.temps[p2]
		tmpTimes1 = append(tmpTimes1, time2)
		tmpTimes2 = append(tmpTimes2, time2)
		tmpTemps2 = append(tmpTemps2, temp2)
		if (p1 - 1) < 0 {
			//如果p1为第一个，这里需要确认mcu1和mcu2的温度是否差不多。
			tmpTemps1 = append(tmpTemps1, temp1)
		} else {
			//如果p1的前一个值有值，则取一个符合线性规律的值。
			preX0 := Time2Unix(task1.times[tp1-1])
			preY0 := task1.temps[tp1-1]
			x1 := Time2Unix(time1)
			y1 := temp1
			x := Time2Unix(time2)

			tmpTemps1 = append(tmpTemps1, linearInterpolation(preX0, preY0, x1, y1, x))
		}

		p2++
	}

	task1.times = tmpTimes1
	task1.temps = tmpTemps1

	task2.times = tmpTimes2
	task2.temps = tmpTemps2

	fmt.Printf("%s===%v\n", task1.times, task1.temps)
	fmt.Printf("%s===%v\n", task2.times, task2.temps)

	fmt.Println(task1)
	fmt.Println(task2)

}

func TaskMCU1MCU2Merge(task1, task2 Task) {

	tmpTimes1 := make([]string, 0, len(task1.times))
	tmpTemps1 := make([]float32, 0, len(task1.temps))
	tmpTimes2 := make([]string, 0, len(task2.times))
	tmpTemps2 := make([]float32, 0, len(task2.temps))

	//len1 := len(task1.times)
	//len2 := len(task2.times)
	//maxLen := math.Max(float64(len1), float64(len2))
	var p1, p2 int
	for p1 < len(task1.times) && p2 < len(task2.times) {
		time1 := task1.times[p1]
		time2 := task2.times[p2]
		temp1 := task1.temps[p1]
		temp2 := task2.temps[p2]

		if time1 == time2 {
			tmpTimes1 = append(tmpTimes1, time1)
			tmpTimes2 = append(tmpTimes2, time2)
			tmpTemps1 = append(tmpTemps1, temp1)
			tmpTemps2 = append(tmpTemps2, temp2)
			p1++
			p1++
		} else if time1 < time2 {
			tmpTimes1 = append(tmpTimes1, time1)
			tmpTimes2 = append(tmpTimes2, time1)
			tmpTemps1 = append(tmpTemps1, temp1)
			//这里需要对tmpTemps2进行温度补偿，这里先补偿一个-1，在最后在进行线性补偿
			tmpTemps2 = append(tmpTemps2, -1)
			//if (p2 - 1) < 0 {
			//	//如果p2为第一个。
			//	tmpTemps2 = append(tmpTemps2, temp2)
			//} else {
			//	//如果p2的前一个值有值，则取一个符合线性规律的值。
			//	preX0 := Time2Unix(task2.times[p2-1])
			//	preY0 := task2.temps[p2-1]
			//	x1 := Time2Unix(time2)
			//	y1 := temp2
			//	x := Time2Unix(time1)
			//
			//	tmpTemps2 = append(tmpTemps2, linearInterpolation(preX0, preY0, x1, y1, x))
			//}

			p1++
		} else {
			tmpTimes1 = append(tmpTimes1, time2)
			tmpTimes2 = append(tmpTimes2, time2)
			tmpTemps2 = append(tmpTemps2, temp2)
			//这里需要对tmpTemps1进行温度补偿，这里先补偿一个-1，在最后在进行线性补偿
			tmpTemps1 = append(tmpTemps1, -1)
			//if (p1 - 1) < 0 {
			//	//如果p1为第一个。
			//	tmpTemps1 = append(tmpTemps1, temp1)
			//} else {
			//	//如果p1的前一个值有值，则取一个符合线性规律的值。
			//	preX0 := Time2Unix(task1.times[p1-1])
			//	preY0 := task1.temps[p1-1]
			//	x1 := Time2Unix(time1)
			//	y1 := temp1
			//	x := Time2Unix(time2)
			//
			//	tmpTemps1 = append(tmpTemps1, linearInterpolation(preX0, preY0, x1, y1, x))
			//}

			p2++
		}
	}

	if p1 < len(task1.times) {
		//var tp2 int
		//if p2 >= len(task2.times) {
		//	tp2 = len(task2.times) - 1
		//}
		//这里task2肯定已经遍历完成
		//tp2 := len(task2.times) - 1
		time1 := task1.times[p1]
		//time2 := task2.times[tp2]
		temp1 := task1.temps[p1]
		//temp2 := task2.temps[tp2]

		tmpTimes1 = append(tmpTimes1, time1)
		tmpTimes2 = append(tmpTimes2, time1)
		tmpTemps1 = append(tmpTemps1, temp1)
		//这里需要对tmpTemps2进行温度补偿，这里先补偿一个-1，在最后在进行线性补偿
		tmpTemps2 = append(tmpTemps2, -1)
		//if (p2 - 1) < 0 {
		//	//如果p2为第一个。
		//	tmpTemps2 = append(tmpTemps2, temp2)
		//} else {
		//	//如果p2的前一个值有值，则取一个符合线性规律的值。
		//	preX0 := Time2Unix(task2.times[tp2-1])
		//	preY0 := task2.temps[tp2-1]
		//	x1 := Time2Unix(time2)
		//	y1 := temp2
		//	x := Time2Unix(time1)
		//
		//	tmpTemps2 = append(tmpTemps2, linearInterpolation(preX0, preY0, x1, y1, x))
		//}

		p1++
	}

	if p2 < len(task2.times) {
		//var tp1 int
		//if p1 >= len(task1.times) {
		//	tp1 = len(task1.times) - 1
		//}

		//这里task1肯定已经遍历完成
		//tp1 := len(task1.times) - 1
		//time1 := task1.times[tp1]
		time2 := task2.times[p2]
		//temp1 := task1.temps[tp1]
		temp2 := task2.temps[p2]
		tmpTimes1 = append(tmpTimes1, time2)
		tmpTimes2 = append(tmpTimes2, time2)
		tmpTemps2 = append(tmpTemps2, temp2)
		//这里需要对tmpTemps1进行温度补偿，这里先补偿一个-1，在最后在进行线性补偿
		tmpTemps1 = append(tmpTemps1, -1)
		//if (p1 - 1) < 0 {
		//	//如果p1为第一个。
		//	tmpTemps1 = append(tmpTemps1, temp1)
		//} else {
		//	//如果p1的前一个值有值，则取一个符合线性规律的值。
		//	preX0 := Time2Unix(task1.times[tp1-1])
		//	preY0 := task1.temps[tp1-1]
		//	x1 := Time2Unix(time1)
		//	y1 := temp1
		//	x := Time2Unix(time2)
		//
		//	tmpTemps1 = append(tmpTemps1, linearInterpolation(preX0, preY0, x1, y1, x))
		//}

		p2++
	}

	task1.times = tmpTimes1
	task1.temps = tmpTemps1

	task2.times = tmpTimes2
	task2.temps = tmpTemps2

	fmt.Printf("%s===%v\n", task1.times, task1.temps)
	fmt.Printf("%s===%v\n", task2.times, task2.temps)
}

func TestCompensateTemperature(t *testing.T) {
	//task := &Task{
	//	times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
	//	temps: []float32{52.96, 42.96, -1, 48.96},
	//}

	//task := &Task{
	//	times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
	//	temps: []float32{-1, 42.96, -1, 48.96},
	//}

	//task := &Task{
	//	times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
	//	temps: []float32{-1, 42.96, -1, -1},
	//}

	task := &Task{
		times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
		temps: []float32{-1, -1, -1, -1},
	}

	//task := &Task{
	//	times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
	//	temps: []float32{42.96, -1, -1, -1},
	//}

	//task := &Task{
	//	times: []string{"2023-08-08 23:37:16", "2023-08-08 23:38:16", "2023-08-08 23:39:06", "2023-08-08 23:39:56"},
	//	temps: []float32{-1, -1, -1, 48.96},
	//}

	CompensateTemperature(task)

	fmt.Printf("%s===%v\n", task.times, task.temps)
	//fmt.Printf("%s===%v\n", task2.times, task2.temps)
}

func CompensateTemperature(task *Task) {
	start := -1
	end := -1

	//var x0, x1, x string
	//var y0, y1 float32
	for i := 0; i < len(task.times); i++ {
		if task.temps[i] != -1 {
			//补偿前半部分全部为-1
			//说明第一次找到第一个不为-1的元素，说明0-end之间的元素全部为-1
			//因此需要将end之前的元素全部补偿为end
			if end == -1 {
				end = i
				for start >= 0 {
					task.temps[start] = task.temps[end]
					start--
				}

				continue
			}

			//标记temp第一个不为-1的元素
			end = i

			//说明还没有遇到temp为-1的元素，无需补偿
			if start == -1 {
				continue
			}

			//start != 0 and end != -1 补偿中间部分
			for start < end {
				x0 := Time2Unix(task.times[start-1])
				y0 := task.temps[start-1]
				x1 := Time2Unix(task.times[end])
				y1 := task.temps[end]
				x := Time2Unix(task.times[start])
				task.temps[start] = linearInterpolation(x0, y0, x1, y1, x)
				start++
			}
			continue
		}
		start = i
	}

	//补偿后半部分全部为-1，或者全部为-1
	if start > end {
		if end == -1 {
			for start >= 0 {
				task.temps[start] = 0
				start--
			}
		} else {
			for start > end {
				task.temps[start] = task.temps[end]
				start--
			}
		}
	}

}

func CompensateTemperature1(task *Task) {
	var start, end int = -1, -1

	for i := 0; i < len(task.times); i++ {
		if task.temps[i] != -1 {
			if end == -1 {
				end = i
				for j := 0; j < end; j++ {
					task.temps[j] = task.temps[end]
				}
			} else {
				linearInterpolate(start, end, i, task)
				start = -1
			}
			end = i
		} else if start == -1 {
			start = i
		}
	}

	// Handle trailing "-1" temperatures
	if start != -1 {
		var fillValue float32
		if end != -1 {
			fillValue = task.temps[end]
		}
		for i := start; i < len(task.temps); i++ {
			task.temps[i] = fillValue
		}
	}
}

func linearInterpolate(start, end, i int, task *Task) {
	if start == -1 || end == -1 || start >= end {
		return
	}

	x0 := Time2Unix(task.times[start-1])
	y0 := task.temps[start-1]
	x1 := Time2Unix(task.times[end])
	y1 := task.temps[end]

	for j := start; j < i; j++ {
		x := Time2Unix(task.times[j])
		task.temps[j] = linearInterpolation(x0, y0, x1, y1, x)
	}
}

func TasksMerge(tasks []*Task) {
	allTimes := make([]string, 0)
	allTemps := make([]float32, 0)
	for i := 0; i < len(tasks); i++ {
		allTimes = append(allTimes, tasks[i].times...)
		allTemps = append(allTemps, tasks[i].temps...)
	}

	tmpTask := Task{
		times: allTimes,
		temps: allTemps,
	}

	sort.SliceStable(tmpTask.times, func(i, j int) bool {
		return tmpTask.times[i] < tmpTask.times[j]
	})

	for i := len(allTimes) - 1; i > 0; i-- {
		if allTimes[i] == allTimes[i-1] {
			//allTimes = append()
		}
	}
}

func TestSort(t *testing.T) {
	allTimes := []string{"2021-01-01", "2021-01-02", "2021-01-02", "2021-01-02", "2021-01-03"}
	allTemps := []float32{26.5, 28.0, 28.1, 28.2, 27.2}
	index := 0 // 新的不重复元素的位置

	for i := 1; i < len(allTimes); i++ {
		if allTimes[i] != allTimes[index] {
			index++
			allTimes[index] = allTimes[i]
			allTemps[index] = allTemps[i]
		} else {
			allTemps[index] = allTemps[i]
		}
	}

	allTimes = allTimes[:index+1]
	allTemps = allTemps[:index+1]

	fmt.Println(allTimes)
	fmt.Println(allTemps)

}

func TestSort1(t *testing.T) {
	allTimes := []string{"2021-01-01", "2021-01-02", "2021-01-02", "2021-01-02", "2021-01-03"}
	allTemps := []float32{26.5, 28.0, -1, -1, 27.2}
	index := 0 // 新的不重复元素的位置

	for i := 1; i < len(allTimes); i++ {
		if allTimes[i] != allTimes[index] {
			index++
			allTimes[index] = allTimes[i]
			allTemps[index] = allTemps[i]
		} else {
			// 如果times重复，但temps中的值不为-1，则更新index指向的值
			if allTemps[i] != -1 {
				allTemps[index] = allTemps[i]
			}
			// 如果times重复且temps中的值为-1，则跳过此元素，不增加index
		}
	}

	allTimes = allTimes[:index+1]
	allTemps = allTemps[:index+1]

	fmt.Println(allTimes)
	fmt.Println(allTemps)

}

//12.11 12.14 12.22 12.28 12.30        12.31 12.32 12.33 12.34 12.35
//12.11 12.14 12.22 12.28 12.27        12.28 12.30 12.33 12.34 12.35
//12.11 12.14 12.22 12.28 12.27 12.30        12.28 12.30 12.31 12.33 12.34 12.35

func (t Task) Len() int {
	return len(t.times)
}

func (t Task) Less(i, j int) bool {
	return t.times[i] < t.times[j]
}

func (t Task) Swap(i, j int) {
	t.times[i], t.times[j] = t.times[j], t.times[i]
	t.temps[i], t.temps[j] = t.temps[j], t.temps[i]
}

func Time2Unix(d string) int64 {
	layout := "2006-01-02 15:04:05"
	tm, err := time.Parse(layout, d)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return tm.Unix()
}

func linearInterpolation(x0 int64, y0 float32, x1 int64, y1 float32, x int64) float32 {
	//if (x1 - x0) == 0 {
	//	return y0
	//}
	return roundFloat32(y0 + float32(x-x0)*(y1-y0)/float32(x1-x0))
}

func roundFloat32(f float32) float32 {
	return float32(int(f*100+0.5)) / 100.0
}

//func TestTask4(t *testing.T) {
//	tasksA := NewTasks()
//	tasksB := NewTasks()
//
//	m := map[string][]*Task{
//		"FSDA": tasksA,
//		"FSDB": tasksB,
//	}
//
//	socTempByTime := make(map[string]map[string]*SocTemp)
//
//	for _, task := range tasksA {
//		for i := 0; i < len(task.SocTemps); i++ {
//			soc := task.SocTemps[i]
//			socTempByTime[soc.RpTime]["FSDA"] = soc
//		}
//	}
//	for _, task := range tasksB {
//		for i := 0; i < len(task.SocTemps); i++ {
//			soc := task.SocTemps[i]
//			socTempByTime[soc.RpTime]["FSDB"] = soc
//		}
//	}
//
//	for k, ms := range socTempByTime {
//		ms
//	}
//}

type SocTemp struct {
	RpTime string
	Temp   float32
}

type Task struct {
	SocTemps []*SocTemp
	times    []string  `json:"times"`
	temps    []float32 `json:"temps"`
}

func NewTasks() []*Task {
	var tasks []*Task
	for i := 0; i < 10; i++ {
		tasks = append(tasks, NewTask())
	}

	//for _, task := range tasks {
	//	fmt.Printf("%s===%v\n", task.times, task.temps)
	//	data, _ := json.Marshal(task)
	//
	//	fmt.Printf("%s", string(data))
	//}

	return tasks
}

func compare() {

}

func generateTimes() []string {
	startTime := time.Date(2023, 8, 11, 12, 9, 3, 0, time.UTC)
	var times []string

	// Decide the length of times slice (between 5 and 10)
	length := rand.Intn(6) + 5

	for i := 0; i < length; i++ {
		// Add a random duration (between 0 and 10 seconds) to make some times identical
		duration := time.Duration(rand.Intn(11)) * time.Second
		times = append(times, startTime.Add(duration).Format("2006-01-02 15:04:05"))
		startTime = startTime.Add(1 * time.Minute) // Increase by 1 minute for the next time
	}

	return times
}

func generateTemps(length int) []float32 {
	var temps []float32
	for i := 0; i < length; i++ {
		temps = append(temps, generateRandomFloat32())
	}
	return temps
}

func NewTask() *Task {
	rand.Seed(time.Now().UnixNano())

	times := generateTimes()
	temps := generateTemps(len(times))
	var socList []*SocTemp
	for i := 0; i < len(times); i++ {
		socList = append(socList, &SocTemp{
			RpTime: times[i],
			Temp:   temps[i],
		})
	}
	t := Task{
		SocTemps: socList,
		times:    times,
		temps:    temps,
	}

	return &t
}
func TestName(t *testing.T) {
	fmt.Println(generateRandomFloat32())
}

func generateRandomFloat32() float32 {
	rand.Seed(time.Now().UnixNano())

	integerPart := rand.Intn(101) // 0-100之间的整数部分
	decimalPart := rand.Float32() // 0-1之间的小数部分

	// 将整数部分和小数部分相加，并保留两位小数
	return float32(integerPart) + float32(int(decimalPart*100+0.5))/100
}

// SampleMax 按照每10个点取一个最大值采样。
func SampleMax(data []float32) []float32 {
	if len(data) == 0 {
		return []float32{}
	}

	const stepSize = 10
	result := make([]float32, 0, len(data)/stepSize)

	for i := 0; i < len(data); i += stepSize {
		maxValue := data[i]
		for j := 1; j < stepSize && i+j < len(data); j++ {
			if data[i+j] > maxValue {
				maxValue = data[i+j]
			}
		}
		result = append(result, maxValue)
	}
	return result
}

func SampleMax3(data []float32) []float32 {
	if len(data) == 0 {
		return []float32{}
	}

	var stepSize = 0
	result := make([]float32, 0, len(data))
	var maxTemp float32
	for i, datum := range data {
		if stepSize < 9 {
			stepSize++
			if maxTemp < datum {
				maxTemp = datum
			}
			if i == len(data)-1 {
				result = append(result, maxTemp)
			}
			continue
		}

		result = append(result, maxTemp)
		stepSize = 0
		maxTemp = 0
	}

	return result
}
func TestName66(t *testing.T) {
	data := []float32{1, 5, 9, 13, 4, 2, 7, 3, 5, 1, 5, 7, 12, 11, 7, 2, 3, 1, 4, 5, 8, 6, 3, 2, 4, 8, 10, 7, 3, 9, 1, 5, 6, 2, 7, 4, 6, 8, 2, 5, 3, 1, 4, 7, 6, 2, 3, 9, 5, 8, 1, 5, 7, 2, 6, 3, 9, 4, 5, 6, 2, 1, 4, 8, 3, 7, 5, 2, 6, 1, 4, 5, 7, 2, 3, 6, 9, 8, 7, 5, 4, 3, 2, 1}
	sampledData := SampleMax(data)
	sampledData3 := SampleMax3(data)
	fmt.Println(sampledData)  // This will print the max value for each 10-element segment.
	fmt.Println(sampledData3) // This will print the max value for each 10-element segment.
}
