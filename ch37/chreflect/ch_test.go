package chreflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Proc struct {
	procName    string   `json:"proc_name"`
	cpuUsages   []int    `json:"cpu_usages"`
	memUsages   []int    `json:"mem_usages"`
	otherUsages []string `json:"other_usages"`
}

func TestName(t *testing.T) {
	//proc := &Proc{}
	//tProc := &Proc{}
	//
	//typ := reflect.TypeOf(proc)
	//val := reflect.ValueOf(proc)
	//if typ.Kind() == reflect.Ptr {
	//	typ = typ.Elem()
	//	val = val.Elem()
	//}
	//num := typ.NumField()
	//
	//for i := 0; i < num; i++ {
	//	typField := typ.Field(i).Type
	//	if typField.Kind() == reflect.Slice {
	//
	//	}
	//}
}

func TestM(t *testing.T) {
	companys1 := make([]*Company, 0)
	companys1 = append(companys1, &Company{
		Name: "company1",
		Employees: []*Employee{&Employee{
			Id:    1,
			Names: []string{"name1", "name2"},
		}},
	})

	companys2 := make([]*Company, 0)
	companys2 = append(companys2, &Company{
		Name: "company2",
		Employees: []*Employee{&Employee{
			Id:    2,
			Names: []string{"name21", "name22"},
		}},
	})
	p1 := &Person{
		Name:     "John",
		Age:      30,
		Emails:   []string{"john@example.com", "johndoe@example.com"},
		Companys: companys1,
	}

	p2 := &Person{
		Name:     "Jane",
		Age:      25,
		Emails:   []string{"jane@example.com", "janedoe@example.com"},
		Companys: companys2,
	}

	mergeStructs2(p1, p2)
	fmt.Println(p1)
}

func mergeStructs2(p1, p2 *Person) {

	p1Value := reflect.ValueOf(p1).Elem()
	p2Value := reflect.ValueOf(p2).Elem()

	for i := 0; i < p1Value.NumField(); i++ {
		field1 := p1Value.Field(i)
		field2 := p2Value.Field(i)

		// 判断字段是否为切片类型
		if field1.Kind() == reflect.Slice {
			field1.Set(reflect.AppendSlice(field1, field2))
		}
	}
}

func mergeStructs(p1, p2 Person) Person {
	merged := Person{
		Name: p1.Name,
		Age:  p1.Age,
	}

	p1Value := reflect.ValueOf(p1)
	p2Value := reflect.ValueOf(p2)
	mergedValue := reflect.ValueOf(&merged).Elem()

	for i := 0; i < p1Value.NumField(); i++ {
		field1 := p1Value.Field(i)
		field2 := p2Value.Field(i)
		mergedField := mergedValue.Field(i)

		// 判断字段是否为切片类型
		if field1.Kind() == reflect.Slice {
			mergedField.Set(reflect.AppendSlice(field1, field2))
		} else {
			mergedField.Set(field1)
		}
	}

	return merged
}

type Person struct {
	Name     string
	Age      int
	Emails   []string
	Companys []*Company
}

// 公司 struct
type Company struct {
	Name      string
	Employees []*Employee
}

// 员工 struct
type Employee struct {
	Id    int
	Names []string
}

func TestStruct(t *testing.T) {

}

func mergeStructs3(p1, p2 *Person) {
	p1Typ := reflect.TypeOf(p1).Elem()
	p2Typ := reflect.TypeOf(p2).Elem()
	p1Value := reflect.ValueOf(p1).Elem()
	p2Value := reflect.ValueOf(p2).Elem()

	num := p2Typ.NumField()
	for i := 0; i < num; i++ {
		field2 := p1Typ.Field(i)
		field2Typ := field2.Type

		// 判断字段是否为切片类型
		if field2Typ.Kind() == reflect.Slice {
			if field2Typ.Elem().Kind() == reflect.Ptr {
				field2ChildTyp := field2Typ.Elem()
				nc := field2ChildTyp.NumField()
				for i := 0; i < nc; i++ {
					field2ChildChild := field2ChildTyp.Field(i)
				}
			}
			//field1.Set(reflect.AppendSlice(field1, field2))
		}
	}
}
