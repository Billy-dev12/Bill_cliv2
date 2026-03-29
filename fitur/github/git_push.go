package github

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitPush() {

	cmd := exec.Command("git", "config", "--global", "user.email")
	output, err := cmd.Output()
	email := strings.TrimSpace(string(output))
	if err != nil || email == "" {
		CheckAndSetGitConfig()
	} else {
		fmt.Println("email mu sudah ada di config git, lanjut push aja :)")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Masukan Link Github Mu :")
	linkRepo, _ := reader.ReadString('\n')
	linkRepo = strings.TrimSpace(linkRepo)

	fmt.Print("Masukan Pesan Commit (Contoh : Update pertama)")
	commitMassage, _ := reader.ReadString('\n')
	commitMassage = strings.TrimSpace(commitMassage)

	if linkRepo == "" {
		fmt.Println("Link repo nya jangan kosong ya")
		return
	}

	if commitMassage == "" {
		commitMassage = "Commit by bil cli"
	}

	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", commitMassage).Run()
	exec.Command("git", "remote", "add", "origin", linkRepo).Run()
	exec.Command("git", "branch", "-M", "main").Run()

	cmdPush := exec.Command("git", "push", "-u", "origin", "main")
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr

	err = cmdPush.Run()

	if err != nil {
		fmt.Println("Gagal melakukan push ke GitHub:", err)
		return
	}

	fmt.Println("Berhasil melakukan push ke GitHub!")
}
