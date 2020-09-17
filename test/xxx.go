package test

import (
	"fmt"
	"reflect"
	_ "reflect"
	"strconv"
)

type person struct {
	int
	name string
}

func (p person) Dis() {
	fmt.Println(p)
}

type student struct {
	person
	age    int
	Gender int
}

func (s student) Sayhello(toname string) (string, int) {
	return s.name + "say hello to" + toname, 1
}
func (s *student) Dis() {
	fmt.Println(s)
}
func Myref() {
	s := student{person: person{int: 1, name: "aaa"}, age: 22}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Println(t)

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		val := v.Field(i)
		fmt.Println(key.Name, key.Type, "|", val)
	}
	fmt.Println("=================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	fmt.Println("=================")
	t1 := reflect.TypeOf(&s)
	v1 := reflect.ValueOf(&s)
	fmt.Println(t1)
	fmt.Println(v1)
	fmt.Println("=================")
	if k := t1.Kind(); k == reflect.Struct {
		for i := 0; i < t1.NumField(); i++ {
			key := t1.Field(i)
			val := v1.Field(i)
			fmt.Println(key.Name, key.Type, "|", val)
		}
	}
	for i := 0; i < t1.NumMethod(); i++ {
		m := t1.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	fmt.Println("==================3=================")
	fmt.Println(t.FieldByName("person"))
	fmt.Println(t.FieldByIndex([]int{0}))
	fmt.Println(t.FieldByIndex([]int{0, 0}), t.FieldByIndex([]int{0, 1}))
	m2, _ := t.MethodByName("Sayhello")
	fmt.Println(m2)
	fmt.Println("==================4=================")
	x := 123
	vx := reflect.ValueOf(&x)
	vx.Elem().SetInt(333)
	fmt.Println(x)
	fmt.Println("============")
	s11 := student{person: person{int: 1, name: "aaa"}, age: 22, Gender: 1}
	fmt.Println(s11)
	v11 := reflect.ValueOf(&s11)
	v11.Elem().FieldByName("age").CanSet()
	fmt.Println(v11.Elem().FieldByName("Gender"))
	if v11.Elem().FieldByName("age").CanSet() {
		v11.Elem().FieldByName("age").SetInt(99)
	}
	if v11.Elem().FieldByName("Gender").CanSet() {
		v11.Elem().FieldByName("Gender").SetInt(2)
	}
	fmt.Println(s11)
	fmt.Println("=====================================")

	params1 := make([]reflect.Value, 1)
	params1[0] = reflect.ValueOf("ppp")
	b := reflect.ValueOf(s11).MethodByName("Sayhello").Call(params1)
	fmt.Println(b[0], "|", b[1])
	b1 := reflect.ValueOf(&s11).MethodByName("Dis").Call([]reflect.Value{})
	fmt.Println(b1)
	fmt.Println("===================5==================")
	str := "this is string"
	fmt.Println(reflect.TypeOf(str))
	fmt.Println("======================================")
	a := new(person)
	a.name = "xxx"
	typ := reflect.TypeOf(a)
	fmt.Println(typ)
	fmt.Println(typ.Elem())
	fmt.Println(typ.NumMethod())
	fmt.Println(typ.Method(0))
	fmt.Println(typ.Name())
	fmt.Println(typ.PkgPath())
	fmt.Println(typ.Size())
	fmt.Println(typ.String())
	fmt.Println(typ.Elem().String)
	fmt.Println(typ.Elem().FieldByIndex([]int{0}))
	fmt.Println(typ.Elem().FieldByName("name"))
	fmt.Println(typ.Kind() == reflect.Ptr)
	fmt.Println(typ.Elem().Kind() == reflect.Struct)
	fmt.Println("========================================")
	fmt.Println(reflect.TypeOf(12.12).Bits())
	fmt.Println("========================================")
	cha := make(chan int)
	fmt.Println(reflect.TypeOf(cha).ChanDir())

	var fun func(x int, y ...float64) string
	var fun2 func(x int, y float64) string
	fmt.Println(reflect.TypeOf(fun).IsVariadic())
	fmt.Println(reflect.TypeOf(fun2).IsVariadic())
	fmt.Println(reflect.TypeOf(fun).In(0))
	fmt.Println(reflect.TypeOf(fun).In(1))
	fmt.Println(reflect.TypeOf(fun).NumIn())
	fmt.Println(reflect.TypeOf(fun).NumOut())
	fmt.Println(reflect.TypeOf(fun).Out(0))

	fmt.Println("=========================================")
	mp := make(map[string]int)
	mp["test1"] = 1

	arr := [1]string{"test"}
	fmt.Println(reflect.TypeOf(arr).Len())
}
func prints(i int) string {
	fmt.Println("i=", i)
	return strconv.Itoa(i)

}
