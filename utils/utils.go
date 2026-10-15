package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPostiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		values, err := strconv.Atoi(input)
		if err == nil && values > 0 {
			return values
		}
		fmt.Println("Gia tri khong hop le")
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/c", "cls")

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error clearing screen: ", err)
	}
}
