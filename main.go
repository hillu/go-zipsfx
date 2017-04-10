package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"github.com/edsrzf/mmap-go"
	"io"
	"log"
	"os"
	"path/filepath"
)

func findSig(buf []byte) (result []int) {
	signature := []byte{0x50, 0x4b, 0x03, 0x04}
	var i, n int
	for {
		if n = bytes.Index(buf[i:], signature); n == -1 {
			break
		}
		r := n + i
		result = append(result, r)
		i = r + 1
	}
	return
}

func main() {
	dest := flag.String("d", ".", "destination")
	flag.Parse()
	f, err := os.Open(os.Args[0])
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	m, err := mmap.Map(f, 0, mmap.RDONLY)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var zr *zip.Reader
	for _, i := range findSig(m) {
		s := int64(len(m) - i)
		if zr, err = zip.NewReader(io.NewSectionReader(f, int64(i), s), s); err != nil {
			continue
		}
		break
	}
	if zr == nil {
		log.Print("ZIP header not found")
		os.Exit(1)
	}
	for _, f := range zr.File {
		path := filepath.Join(*dest, f.Name)
		log.Print("Extracting to " + path)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			zf, err := f.Open()
			if err != nil {
				log.Print(err)
				continue
			}
			of, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Print(err)
				continue
			}
			_, err = io.Copy(of, zf)
			zf.Close()
			of.Close()
			if err != nil {
				log.Print(err)
			}
		}
	}
	log.Print("Done.")
}
