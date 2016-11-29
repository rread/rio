package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

var options struct {
	BlockSize int64
	FileSize  int64
	TestDir   string
}

func init() {
	var block, size string
	flag.StringVar(&block, "block", "64KiB", "Block size")
	flag.StringVar(&size, "size", "1GiB", "File size")
	flag.StringVar(&options.TestDir, "dir", ".", "Test directory")
	flag.Parse()

	blk, err := humanize.ParseBytes(block)
	if err != nil {
		log.Fatal(err)
	}
	options.BlockSize = int64(blk)

	fsize, err := humanize.ParseBytes(size)
	if err != nil {
		log.Fatal(err)
	}
	options.FileSize = int64(fsize)
}

func writeBytes(w io.Writer, size int64, block int64) (int64, error) {
	buf := make([]byte, block)
	var written int64
	for written < size {
		n, err := w.Write(buf)
		written += int64(n)
		if err != nil {
			return written, errors.Wrap(err, "write failed")
		}
	}
	return written, nil
}

func readBytes(r io.Reader, size int64, block int64) (int64, error) {
	buf := make([]byte, block)
	var in int64
	for in < size {
		n, err := r.Read(buf)
		in += int64(n)
		if err != nil {
			return in, errors.Wrap(err, "read failed")
		}
	}
	return in, nil
}

func result(label string, total int64, bs int64, dur time.Duration) {
	fmt.Printf("%s %v %v/io %.2fs %2.2f IO/s %2.2f MiB/s\n", label, humanize.IBytes(uint64(total)),
		humanize.IBytes(uint64(bs)),
		dur.Seconds(),
		float64(total/bs)/dur.Seconds(),
		float64(total/(1024*1024))/dur.Seconds())
}

func main() {
	fp, err := ioutil.TempFile(options.TestDir, "rio-data-")
	if err != nil {
		log.Fatal(errors.Wrap(err, "TempFile failed"))
	}
	defer os.Remove(fp.Name())

	start := time.Now()
	written, err := writeBytes(fp, options.FileSize, options.BlockSize)
	fp.Sync()
	elapsed := time.Since(start)
	result("WRITE", written, options.BlockSize, elapsed)
	if err != nil {
		fmt.Printf("write test failed %v\n", err)
	}

	flushPages()

	fp.Seek(0, 0)

	start = time.Now()
	in, err := readBytes(fp, options.FileSize, options.BlockSize)
	fp.Seek(0, 0)
	elapsed = time.Since(start)
	result("READ", in, options.BlockSize, elapsed)
	if err != nil {
		fmt.Printf("read test failed %v\n", err)
	}
}
