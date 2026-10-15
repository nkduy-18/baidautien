package main

import (
	"fmt"

	"example.com/hello/utils"
)

type GiaoDich struct {
	TieuDe string
	SoTien int
}

var thunhapList []GiaoDich
var chitieuList []GiaoDich

func nhapthunhap() {
	tieude := utils.ReadInput("Nhập tiêu đề thu nhập: ")
	sotien := utils.GetPostiveInt("Nhập số tiền :")
	thunhap := GiaoDich{
		TieuDe: tieude,
		SoTien: sotien,
	}
	thunhapList = append(thunhapList, thunhap)
}
func nhapchitieu() {
	tieude := utils.ReadInput("Nhập tiêu đề chi tiêu: ")
	sotien := utils.GetPostiveInt("Nhập số tiền:")
	chitieu := GiaoDich{
		TieuDe: tieude,
		SoTien: sotien,
	}
	chitieuList = append(chitieuList, chitieu)
}
func danhsachthunhap() {
	fmt.Println("---Danh sách thu nhập---")
	if len(thunhapList) == 0 {
		fmt.Println("Chưa có khoản thu nào.")
		return
	}
	for _, t := range thunhapList {
		fmt.Println(t)
	}
}
func danhsachchitieu() {
	fmt.Println("---Danh sách chi tiêu---")
	if len(chitieuList) == 0 {
		fmt.Println("Chưa có khoản chi nào.")
		return
	}
	for _, c := range chitieuList {
		fmt.Println(c)
	}
}
func danhsach() {
	tongthu := 0
	tongchi := 0
	for _, values := range thunhapList {
		tongthu = tongthu + values.SoTien
	}
	for _, values := range chitieuList {
		tongchi = tongchi + values.SoTien
	}
	fmt.Println("--------------")
	fmt.Println("Số dư hiện tại là:", tongthu-tongchi)
	fmt.Println("Tổng số tiền thu nhập là:", tongthu)
	fmt.Println("Tổng số tiền đã chi là:", tongchi)
	danhsachthunhap()
	danhsachchitieu()
}
func main() {
	for {
		utils.ClearScreen()
		fmt.Println("---CHƯƠNG TRÌNH QUẢN LÝ CHI TIÊU---")
		fmt.Println("1.Nhập số tiền thu nhập")
		fmt.Println("2.Nhập số tiền chi tiêu")
		fmt.Println("3.Hiển thị số dư và giao dịch")
		fmt.Println("4.Thoát chương trình")
		choice := utils.GetPostiveInt("Mời chọn chức năng:")
		switch choice {
		case 1:
			nhapthunhap()
		case 2:
			nhapchitieu()
		case 3:
			danhsach()
		case 4:
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ")
		}
		utils.ReadInput("\nNhấn phím Enter để tiếp tục...")
	}
}
