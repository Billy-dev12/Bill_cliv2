package github

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckAndSetGitConfig() {

	fmt.Println("wah github mu belum config email ya haha :)")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan email anda : ")
	emailInput, _ := reader.ReadString('\n')
	emailInput = strings.TrimSpace(emailInput)

	fmt.Print("Masukan nama github mu ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	// set git config
	exec.Command("git", "config", "--global", "user.email", emailInput).Run()
	exec.Command("git", "config", "--global", "user.name", nameInput).Run()
	fmt.Println("sip mantap users config udah beres")
}
