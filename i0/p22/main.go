package main


func main() {
	ret := generateParenthesis(3)
	for _, str := range ret {
		println(str)
	}
}

func generateParenthesis(n int) []string {
	var ret []string
	ret = generate("", n, 0, ret)
	return ret
}

func generate (str string, left, right int, ret []string) []string {
	if left == 0 && right == 0 {
		ret = append(ret, str)
		return ret
	}
	l, r := left, right
	if left != 0 {
		ret = generate(str + "(", l - 1, r + 1, ret)
	}
	if right != 0 {
		ret = generate(str + ")", l, r - 1, ret)
	}
	return ret
}
