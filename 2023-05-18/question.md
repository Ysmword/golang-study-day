下面这段代码能否通过编译，不能的话原因是什么；如果通过，输出什么。

func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}

答：通过不了，因为new返回的是[\]int指针，而不是[\]int