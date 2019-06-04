package controller

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func SaveToFile(s string) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	f, err := os.OpenFile(filepath.Join(dir, "rompe.txt"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	currentTime := time.Now()
	lines := strings.Split(s, "\n")

	out := fmt.Sprintf("%s:\n \t%s\n\n", currentTime.Format("2006-01-02"), strings.Join(lines, "\n\t"))

	if _, err = f.WriteString(out); err != nil {
		panic(err)
	}
}
