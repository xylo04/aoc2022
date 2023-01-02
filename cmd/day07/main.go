package main

import (
	"bufio"
	"fmt"
	"github.com/psanford/memfs"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var lsFileRegex, _ = regexp.Compile("(\\d+) (.*)")

type dirSize struct {
	path string
	size int
}

func main() {
	terminal, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := buildFs(string(terminal))
	dirs := findDirSizes(f)
	name, size := findDirToDelete(dirs)
	fmt.Printf("Delete dir %s to free up %d bytes\n", name, size)
}

func sumSmallDirs(dirs map[string]int64) int64 {
	var sum int64 = 0
	for _, v := range dirs {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func buildFs(terminal string) fs.FS {
	rootFS := memfs.New()
	cwd := ""
	s := bufio.NewScanner(strings.NewReader(terminal))
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "$ cd") {
			target := line[5:]
			cwd = filepath.Clean(cwd + string(filepath.Separator) + target)
		} else if strings.HasPrefix(line, "$ ls") {
			// ignore
		} else if strings.HasPrefix(line, "dir") {
			target := line[4:]
			_ = rootFS.MkdirAll(filepath.Clean(cwd + string(filepath.Separator) + target)[1:], 0777)
		} else if lsFileRegex.MatchString(line) {
			matches := lsFileRegex.FindStringSubmatch(line)
			target := matches[2]
			size, _ := strconv.ParseInt(matches[1], 10, 32)
			_ = rootFS.WriteFile(filepath.Clean(cwd + string(filepath.Separator) + target)[1:], randomBytes(int(size)), 0777)
		}
	}
	return rootFS
}

func randomBytes(size int) []byte {
	var b = make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func findDirSizes(f fs.FS) map[string]int64 {
	var rv = make(map[string]int64, 0)
	_ = fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			size := findRecursiveSize(f, path)
			rv[path] = size
		}
		return nil
	})
	return rv
}

func findRecursiveSize(f fs.FS, path string) int64 {
	size := int64(0)
	_ = fs.WalkDir(f, path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			info, _ := d.Info()
			size += info.Size()
		}
		return nil
	})
	return size
}

const totalStorage = 70000000
const requireFreeStorage = 30000000

func findDirToDelete(dirs map[string]int64) (string, int64) {
	currentFree := totalStorage - dirs["."]
	needToFree := requireFreeStorage - currentFree
	candidateName := ""
	candidateSize := int64(math.MaxInt64)
	for k, v := range dirs {
		if v >= needToFree && v < candidateSize {
			candidateSize = v
			candidateName = k
		}
	}
	return candidateName, candidateSize
}
