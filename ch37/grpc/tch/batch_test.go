package tch

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// MockAPI 是一个模拟API调用的结构体
type MockAPI struct {
	Data     []int
	PageSize int
}

// BatchFunc 是一个批处理函数，处理每一批次的数据
type BatchFunc func([]int)

// NewMockAPI 创建一个新的 MockAPI
func NewMockAPI(data []int, pageSize int) *MockAPI {
	return &MockAPI{
		Data:     data,
		PageSize: pageSize,
	}
}

// FindInBatches 使用指定的批处理函数来处理数据
func (api *MockAPI) FindInBatches(f BatchFunc) {
	pageCount := (len(api.Data) + api.PageSize - 1) / api.PageSize

	for i := 0; i < pageCount; i++ {
		start := i * api.PageSize
		end := start + api.PageSize
		if end > len(api.Data) {
			end = len(api.Data)
		}
		f(api.Data[start:end])
	}
}

//1.将底片中的精修照片删掉。
//2.你从底片中选图片,将选中的图片保存在另一个文件夹中。
//3.匹配底片中的文件名和Raf文件名相同的文件，将Raf文件保存在另一个bak目录。

func TestDel(t *testing.T) {
	folder1 := "/Users/henrychen/Desktop/test/src"
	folder2 := "/Users/henrychen/Desktop/test/retouched_images"

	files, err := ioutil.ReadDir(folder1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// 检查文件是否存在于 folder2，无视文件扩展名
		filename := file.Name()
		ext := filepath.Ext(filename)
		nameWithoutExt := filename[0 : len(filename)-len(ext)]

		matches, err := filepath.Glob(filepath.Join(folder2, nameWithoutExt+".*"))
		if err != nil {
			log.Printf("failed to match files: %v", err)
			continue
		}

		// 如果在 folder2 中找到了匹配的文件，就删除 folder1 中的这个文件
		if len(matches) > 0 {
			target := filepath.Join(folder1, file.Name())
			err := os.Remove(target)
			if err != nil {
				log.Printf("failed to delete file: %v", err)
			} else {
				log.Printf("deleted file: %s", target)
			}
		}
	}
}

func TestMove(t *testing.T) {
	folder1 := "/Users/henrychen/Desktop/test/jpg"
	folder2 := "/Users/henrychen/Desktop/test/raf"
	folder3 := "/Users/henrychen/Desktop/test/bak"

	// 检查 folder3 是否存在，如果不存在则创建
	if _, err := os.Stat(folder3); os.IsNotExist(err) {
		os.Mkdir(folder3, 0755)
	}

	files, err := ioutil.ReadDir(folder1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// 检查文件是否存在于 folder2，无视文件扩展名
		filename := file.Name()
		ext := filepath.Ext(filename)
		nameWithoutExt := filename[0 : len(filename)-len(ext)]

		matches, err := filepath.Glob(filepath.Join(folder2, nameWithoutExt+".*"))
		if err != nil {
			log.Printf("failed to match files: %v", err)
			continue
		}

		// 复制所有匹配的文件到 folder3
		for _, match := range matches {
			dest := filepath.Join(folder3, filepath.Base(match))
			err := copyFile(match, dest)
			if err != nil {
				log.Printf("failed to copy file: %v", err)
			} else {
				log.Printf("copied file: %s", dest)
			}
		}
	}
}

//func main() {
//	folder1 := "path/to/folder1"
//	folder2 := "path/to/folder2"
//	folder3 := "path/to/folder3"
//
//	// 检查 folder3 是否存在，如果不存在则创建
//	if _, err := os.Stat(folder3); os.IsNotExist(err) {
//		os.Mkdir(folder3, 0755)
//	}
//
//	files, err := ioutil.ReadDir(folder1)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, file := range files {
//		// 检查文件是否存在于 folder2，无视文件扩展名
//		filename := file.Name()
//		ext := filepath.Ext(filename)
//		nameWithoutExt := filename[0 : len(filename)-len(ext)]
//
//		matches, err := filepath.Glob(filepath.Join(folder2, nameWithoutExt+".*"))
//		if err != nil {
//			log.Printf("failed to match files: %v", err)
//			continue
//		}
//
//		// 复制所有匹配的文件到 folder3
//		for _, match := range matches {
//			dest := filepath.Join(folder3, filepath.Base(match))
//			err := copyFile(match, dest)
//			if err != nil {
//				log.Printf("failed to copy file: %v", err)
//			} else {
//				log.Printf("copied file: %s", dest)
//			}
//		}
//	}
//}

// copyFile 拷贝文件的函数
func copyFile(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
