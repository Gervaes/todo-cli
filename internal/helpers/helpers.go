package helpers

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
)

// https://stackoverflow.com/a/72863273
func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetEnvVariable(key string) string {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
