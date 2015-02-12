// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//+build ignore

// Release is a tool for building the NDK tarballs hosted on dl.google.com.
//
// The Go toolchain only needs the gcc compiler and headers, which are ~10MB.
// The entire NDK is ~400MB. Building smaller toolchain binaries reduces the
// run time of gomobile init significantly.
package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const ndkVersion = "ndk-r10d"

type version struct {
	os   string
	arch string
}

var hosts = []version{
	// TODO: windows
	/*
		{"darwin", "x86"},
		{"darwin", "x86_64"},
		{"linux", "x86"},
	*/
	{"linux", "x86_64"},
}

var tmpdir string

func main() {
	var err error
	tmpdir, err = ioutil.TempDir("", "gomobile-release-")
	if err != nil {
		log.Fatal(err)
	}

	for _, host := range hosts {
		if err := mkpkg(host); err != nil {
			log.Fatal(err)
		}
	}
}

func mkpkg(host version) error {
	ndkName := "android-" + ndkVersion + "-" + host.os + "-" + host.arch + "."
	if host.os == "windows" {
		ndkName += "exe"
	} else {
		ndkName += "bin"
	}
	url := "http://dl.google.com/android/ndk/" + ndkName
	log.Printf("%s\n", url)
	binPath := tmpdir + "/" + ndkName
	if err := fetch(binPath, url); err != nil {
		log.Fatal(err)
	}
	if err := inflate(binPath); err != nil {
		return err
	}
	// The NDK is unpacked into tmpdir/android-ndk-r10d.
	// Move the files we want into tmpdir/linux-x86_64/android-ndk-r10d.
	// We preserve the same file layout to make the full NDK interchangable
	// with the cut down file.
	usr := "android-" + ndkVersion + "/platforms/android-15/arch-arm/usr"
	gcc := "android-" + ndkVersion + "/toolchains/arm-linux-androideabi-4.8/prebuilt/" + host.os + "-" + host.arch
	dst := tmpdir + "/" + host.os + "-" + host.arch
	if err := os.MkdirAll(dst+"/"+usr, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(dst+"/"+gcc, 0755); err != nil {
		return err
	}
	if err := move(dst+"/"+usr, tmpdir+"/"+usr, "include", "lib"); err != nil {
		return err
	}
	if err := move(dst+"/"+gcc, tmpdir+"/"+gcc, "bin", "lib", "libexec"); err != nil {
		return err
	}

	// Build the tarball.
	dst += "/"
	f, err := os.Create("gomobile-ndk-r10d-" + host.os + "-" + host.arch + ".tgz")
	if err != nil {
		return err
	}
	tw := tar.NewWriter(gzip.NewWriter(bufio.NewWriter(f)))
	err = filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		err = tw.WriteHeader(&tar.Header{
			Name: path[len(dst):],
			Size: info.Size(),
		})
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, f)
		f.Close()
		return err
	})
	err2 := tw.Close()
	if err != nil {
		return err
	}
	return err2
}

func fetch(dst, url string) error {
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	err2 := resp.Body.Close()
	err3 := f.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return err3
}

func inflate(path string) error {
	p7zip := "7z"
	if runtime.GOOS == "darwin" {
		p7zip = "/Applications/Keka.app/Contents/Resources/keka7z"
	}
	cmd := exec.Command(p7zip, "x", path)
	cmd.Dir = tmpdir
	out, err := cmd.CombinedOutput()
	if err != nil {
		os.Stderr.Write(out)
		return err
	}
	return nil
}

func move(dst, src string, names ...string) error {
	for _, name := range names {
		if err := os.Rename(src+"/"+name, dst+"/"+name); err != nil {
			return err
		}
	}
	return nil
}
