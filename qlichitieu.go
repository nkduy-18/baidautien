package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var thunhap = make(map[string]int)
	var chitieu = make(map[string]int)
	tongthu := 0
	tongchi := 0
	//1 nhập số tiền
	for {
		fmt.Print("Nhập tiêu đề thu nhập : ")
		scanner.Scan()
		tdthu := strings.TrimSpace(scanner.Text())
		if tdthu == "" {
			break
		}
		fmt.Print("Số tiền : ")
		scanner.Scan()
		stthu, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		tongthu = tongthu + stthu
		thunhap[tdthu] = stthu
	}
	for {
		fmt.Print("Nhập tiêu đề chi tiêu : ")
		scanner.Scan()
		tdchi := strings.TrimSpace(scanner.Text())
		if tdchi == "" {
			break
		}
		fmt.Print("Số tiền : ")
		scanner.Scan()
		stchi, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		tongchi = tongchi + stchi
		chitieu[tdchi] = stchi
	}
	//2 hiển thị toàn bộ thông tin
	fmt.Println("Tổng thu nhập : ", tongthu)
	fmt.Println("Tổng chi tiêu : ", tongchi)
	fmt.Println("Số dư : ", tongthu-tongchi)
	fmt.Println("Thu nhập ", thunhap)
	fmt.Println("Chi tiêu ", chitieu)
}
