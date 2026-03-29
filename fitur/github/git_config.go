package github

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func getConfigFile() string {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".config", "bill")

	// Bikin foldernya otomatis kalau belum ada
	os.MkdirAll(configDir, os.ModePerm)

	return filepath.Join(configDir, "config.json")
}

// 2. Fungsi untuk MENYIMPAN Token ke dalam file
func SaveToken(token string) error {
	filePath := getConfigFile()

	// Menulis isi token ke file config.json
	err := os.WriteFile(filePath, []byte(token), 0600) // 0600 artinya cuma user ini yang boleh baca/tulis
	return err
}

// 3. Fungsi untuk MEMBACA Token dari file
func GetSavedToken() string {
	filePath := getConfigFile()

	// Membaca file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "" // Kembalikan kosong kalau file belum ada atau gagal dibaca
	}

	return strings.TrimSpace(string(data))
}
