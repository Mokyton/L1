package main

import "fmt"

func main() {
	set1 := []any{1, 2, 3, 4}
	set2 := []any{5, 2, 3, 6}
	res := crossroad1(set1, set2)
	fmt.Println(res)
	set3 := []any{1, 2, 2, 2, 3, 3, 4, 5, 6}
	set4 := []any{7, 2, 7, 7, 7, 2, 3, 3, 3}
	res = crossroad2(set3, set4)
	fmt.Println(res)
}

func crossroad1(set1, set2 []any) []any { // получаем 2 множества, объединяем их в один слайс
	total := append(set1, set2...)
	set := make(map[any]struct{})
	res := make([]any, 0, len(total)/2) // заранее даем капасити чуть больше, чтобы выйграть время во время реалокации памяти

	for i := 0; i < len(total); i++ { // добавляем елементы в мапу если в мапе уже есть этот элемент значит этот элемент пересекается в множествах
		_, ok := set[total[i]]
		if !ok {
			set[total[i]] = struct{}{}
			continue
		}
		res = append(res, total[i]) //  добавляем пересечение в финальное множество
	}
	return res // минус подхода мы должны быть 100%, что нам подадут именно множества
}

func crossroad2(set1, set2 []any) []any { //  более безопасный вариант, но требует больше памяти и работает более медленно
	set1 = deleteDuplicates(set1)
	set2 = deleteDuplicates(set2)
	total := append(set1, set2...)
	set := make(map[any]int)
	res := make([]any, 0, len(total)/2)
	for i := 0; i < len(total); i++ {
		set[total[i]] += 1
	}
	for k, v := range set {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func deleteDuplicates(d []any) []any {
	var res []any
	set := make(map[any]struct{})
	for _, v := range d {
		set[v] = struct{}{}
	}
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}
