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

func findSig(buf []byte) int64 {
	// var signature = []byte{0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x08, 0x00}
	var signature = []byte{175, 180, 252, 251, 235, 255, 247, 255}
	for i := range signature {
		signature[i] = signature[i] ^ 0xff
	}
	return int64(bytes.Index(buf, signature))
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
	i := findSig(m)
	if i == -1 {
		log.Print("ZIP header not found")
		os.Exit(1)
	}
	s := int64(len(m)) - i
	zr, err := zip.NewReader(io.NewSectionReader(f, i, s), s)
	if err != nil {
		log.Print(err)
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
