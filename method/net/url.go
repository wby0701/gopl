package main

import "fmt"

type values map[string][]string

func (v values) Get(key string) string{
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return "sb"
}

func (v values) Add(key, value string) {
	v[key] = append(v[key], value)
}


func main() {
	m := values{"lang" : {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	//m.Add("item", "3")

	c := make(values)
	c.Add("item", "555")
	fmt.Println(c.Get("item"))
	fmt.Println(len(c))
}