package main

import (
	"fmt"
	"project/model"
)

var m = model.SftLawMainBody{}

func main() {
	destroy()
}

func destroy() {
	m.Destroy(4146)
}

func d() {
	m.Id = 4146
	res := m.Delete(m)
	fmt.Println(res)
}

func save() {
	//m.Id = 4145
	m.Pid = 1
	m.Name = "测试ssss"
	m.LawCategoryCode = "10"
	res := m.Save(m)
	fmt.Println(res)
}

func get() {
	var arr [][]string
	arr = [][]string{
		{"id", "=", "4146"},
	}
	data := m.First(arr, "id")
	fmt.Println(data)
}
