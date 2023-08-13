package t

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name     string
	Age      int
	Emails   []string
	Companys []*Company
}

type Company struct {
	Name      string
	Employees []*Employee
}

type Employee struct {
	Id    int
	Names []string
}

func TestMain1(t *testing.T) {
	p1 := Person{
		Name:   "John",
		Age:    30,
		Emails: []string{"john@example.com"},
		Companys: []*Company{
			{Name: "Company1", Employees: []*Employee{{Id: 1, Names: []string{"John"}}}},
		},
	}

	p2 := Person{
		Name:   "Alice",
		Age:    25,
		Emails: []string{"alice@example.com"},
		Companys: []*Company{
			{Name: "Company2", Employees: []*Employee{{Id: 2, Names: []string{"Alice"}}}},
		},
	}

	merged := mergePersons(p1, p2)

	str, _ := json.Marshal(merged)
	fmt.Printf("%s\n", string(str))
}

func mergePersons(p1, p2 Person) Person {
	merged := Person{
		Name: p1.Name,
		Age:  p1.Age,
	}

	p1Value := reflect.ValueOf(p1)
	p2Value := reflect.ValueOf(p2)
	mergedValue := reflect.ValueOf(&merged).Elem()

	for i := 0; i < p1Value.NumField(); i++ {
		//fieldName := p1Value.Type().Field(i).Name
		fieldValue1 := p1Value.Field(i)
		fieldValue2 := p2Value.Field(i)

		if fieldValue1.Kind() == reflect.Slice && fieldValue2.Kind() == reflect.Slice && fieldValue1.Type().Elem() == reflect.TypeOf("") {
			// Merge []string slices
			mergedValue.Field(i).Set(reflect.AppendSlice(fieldValue1, fieldValue2))
		} else {
			// Copy non-slice fields
			mergedValue.Field(i).Set(fieldValue1)
		}
	}

	return merged

}
