package utils

import (
	"context"
	"crypto/md5"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CurrentDate() (year, day int, err error) {
	abs, err := filepath.Abs(".")
	if err != nil {
		return
	}
	//fmt.Println(abs)

	dateDir := abs //filepath.Dir(abs)
	//fmt.Printf("dateDir: %s\n", dateDir)
	if name := filepath.Base(dateDir); regexp.MustCompile(`^\d\d`).MatchString(name) {
		day, _ = strconv.Atoi(name[:2])
	} else {
		err = fmt.Errorf("parent directory /%s/. must be format /DD/. representing the AOC day", name)
		fmt.Println(err)
		return
	}
	yearDir := filepath.Dir(dateDir)
	//fmt.Printf("yearDir: %s\n", yearDir)
	year = time.Now().Year()
	if name := filepath.Base(yearDir); regexp.MustCompile(`^\d{4}$`).MatchString(name) {
		year, _ = strconv.Atoi(name)
	}

	return
}

func GetWith(uri string, headers map[string]string) (io.ReadCloser, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["User-Agent"] = "github.com/pemoreau/advent-of-code"
	// hash url with md5
	h := md5.New()
	h.Write([]byte(u.String()))
	for k, v := range headers {
		h.Write([]byte("|"))
		h.Write([]byte(k))
		h.Write([]byte("|"))
		h.Write([]byte(v))
	}
	//id := fmt.Sprintf("%x", h.Sum(nil))
	// cached http get
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	time.Sleep(1 * time.Second)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func fetchUserInput(year, day int, session string) (string, error) {

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	rc, err := GetWith(url, map[string]string{
		"Cookie": "session=" + session,
	})
	if err != nil {
		return "", errors.New("failed to fetch aoc input")
	}
	defer rc.Close()
	b, err := io.ReadAll(rc)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read aoc input: %v", err))
	}
	return strings.TrimSpace(string(b)), nil
}

func CreateFunc(path string, fn func() (string, error)) error {
	s, err := os.Stat(path)

	if errors.Is(err, fs.ErrNotExist) {
		contents, err := fn()
		if err != nil {
			return err
		}
		if contents == "" {
			return nil
		}
		if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("failed to write: %s: is already a directory", path)
	}
	return nil
}

func DownloadInput() ([]byte, error) {
	year, day, err := CurrentDate()

	session := os.Getenv("AOC_SESSION")

	if session != "" {
		//fmt.Printf("session: %s\n", session)
		var fetch = func() (string, error) { return fetchUserInput(year, day, session) }
		if err = CreateFunc("input.txt", fetch); err != nil {
			return nil, err
		}
		return os.ReadFile("input.txt")
	} else {
		return nil, errors.New("missing AOC_SESSION")
	}
}

func tryReadfile(filename string) ([]byte, error) {
	s, err := os.Stat(filename)
	if err == nil && !s.IsDir() {
		b, err := os.ReadFile(filename)
		if err == nil {
			return b, nil
		}
	}
	return nil, err
}

func listfiles(dir string) {
	fmt.Println("listing", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func Input() string {
	inputsDir := os.Getenv("AOC_INPUTS")
	var filename = "input.txt"
	// try to open file name in current dir first
	// then in AOC_INPUTS dir
	// otherwise, download file

	//var currentDir = filepath.Dir(abs)
	//var abs, _ = filepath.Abs(filename)
	//fmt.Println("trying", abs)
	if b, err := tryReadfile(filename); err == nil {
		//abs, _ = filepath.Abs(filename)
		//fmt.Println("found:", abs)
		return string(b)
	}

	if inputsDir != "" {
		year, day, err := CurrentDate()
		if err != nil {
			panic(err)
		}

		syear := fmt.Sprintf("%4d", year)
		sday := fmt.Sprintf("%02d", day)
		filename = filepath.Join(inputsDir, syear, sday, filename)

		for range 4 {
			filename = filepath.Join("..", filename)
			//abs, _ := filepath.Abs(filename)
			//fmt.Println("trying", abs)
			//currentDir, _ = filepath.Abs(filepath.Join(currentDir, ".."))
			//listfiles(filepath.Dir(currentDir))
			if b, err := tryReadfile(filename); err == nil {
				//abs, _ = filepath.Abs(filename)
				//fmt.Println("found:", abs)
				return string(b)
			}
		}
	}

	fmt.Println("downloading input")
	if b, err := DownloadInput(); err == nil {
		return string(b)
	} else {
		panic(err)
	}
}
