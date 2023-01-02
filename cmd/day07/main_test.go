package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

const terminal = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

var mapFS = fstest.MapFS{
	"a":       {Mode: fs.ModeDir},
	"b.txt":   {Data: randomBytes(14848514)},
	"c.dat":   {Data: randomBytes(8504156)},
	"d":       {Mode: fs.ModeDir},
	"a/e":     {Mode: fs.ModeDir},
	"a/f":     {Data: randomBytes(29116)},
	"a/g":     {Data: randomBytes(2557)},
	"a/h.lst": {Data: randomBytes(62596)},
	"a/e/i":   {Data: randomBytes(584)},
	"d/j":     {Data: randomBytes(4060174)},
	"d/d.log": {Data: randomBytes(8033020)},
	"d/d.ext": {Data: randomBytes(5626152)},
	"d/k":     {Data: randomBytes(7214296)},
}

var dirSizes = map[string]int64{
	"a/e": 584,
	"a":   94853,
	"d":   24933642,
	".":   48381165,
}

func Test_buildFs(t *testing.T) {

	type args struct {
		terminal string
	}
	tests := []struct {
		name string
		args args
		want fs.FS
	}{
		{name: "1", args: args{terminal: terminal}, want: mapFS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildFs(tt.args.terminal)
			if err := fsEqual(got, tt.want); err != nil {
				t.Errorf("%v\nwant %v\ngot %v", err, stringFs(tt.want), stringFs(got))
			}
		})
	}
}

// Only care about directory structure and file sizes, not content.
func fsEqual(got fs.FS, want fs.FS) error {
	return fs.WalkDir(want, ".", func(path string, d fs.DirEntry, err error) error {
		gotFile, err := got.Open(path)
		if err != nil {
			return err
		}
		gotStat, err := gotFile.Stat()
		if err != nil {
			return err
		}
		wantStat, err := d.Info()
		if err != nil {
			return err
		}

		if filepath.Clean(wantStat.Name()) != filepath.Clean(gotStat.Name()) ||
			wantStat.IsDir() != gotStat.IsDir() {
			return errors.New("structure mismatch at " + path)
		}
		if !wantStat.IsDir() && wantStat.Size() != gotStat.Size() {
			return errors.New("size mismatch at " + path)
		}
		return nil
	})
}

func stringFs(f fs.FS) string {
	var sb strings.Builder
	_ = fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(&sb, "%s %d\n", path, info.Size())
		return err
	})
	return sb.String()
}

func Test_findDirSizes(t *testing.T) {
	type args struct {
		f fs.FS
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{name: "1", args: args{f: mapFS}, want: dirSizes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDirSizes(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDirSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumSmallDirs(t *testing.T) {
	type args struct {
		dirs map[string]int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{dirs: dirSizes}, want: 95437},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSmallDirs(tt.args.dirs); got != tt.want {
				t.Errorf("sumSmallDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}
