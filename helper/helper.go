package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const session = "53616c7465645f5f3a1e3d838561b6884099854592725eb54916e1d36b96fa364ed725bd55f62b425743db9c4c2dcb21b4f3d7860db06ddb07097e2125842a63"

func DownloadInput() {
	fInfo, _ := os.Stat("input.txt")
	if fInfo != nil {
		return
	}

	p, _ := os.Getwd()
	d := filepath.Base(p)
	dNum, _ := strconv.Atoi(strings.TrimPrefix(d, "day"))
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", dNum), nil)
	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", session
	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "neilo40's golang client")
	c := &http.Client{Timeout: 60 * time.Second}
	r, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	bodyBytes, _ := io.ReadAll(r.Body)
	os.WriteFile("input.txt", bodyBytes, 0644)
}
