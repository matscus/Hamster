package arch

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/matscus/Hamster/Package/Arch/subset"
)

type Arch struct {
}

func New(a Arch) subset.Arch {
	var arch subset.Arch
	arch = Arch{}
	return arch
}

func (a Arch) Compress(dir string, buf io.Writer) (err error) {
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)
	filepath.Walk(dir, func(file string, fi os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(file)
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
	if err := tw.Close(); err != nil {
		return err
	}
	if err := zr.Close(); err != nil {
		return err
	}
	return nil
}

func (a Arch) Decompress(file io.Reader, dst string) (err error) {
	zr, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("err ", err)
		return err
	}
	tr := tar.NewReader(zr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		target := filepath.Join(dst, header.Name)
		if !validRelPath(header.Name) {
			return fmt.Errorf("tar contained invalid name error %q\n", target)
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			fileToWrite.Close()
		}
	}
	return nil
}

// check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}
