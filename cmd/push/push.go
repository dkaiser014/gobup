/*
*
* File and Dir functions are code snippets from
* https://blog.depa.do/post/copy-files-and-directories-in-go
*
 */

package push

import (
	"bup/pkg/getbupdir"
	"bup/pkg/getrootdir"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// File copies a single file from the root_directory
// to the backup_directory
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		log.Fatal("Error:", err)
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		log.Fatal("Error:", err)
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		log.Fatal("Error:", err)
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		log.Fatal("Error", err)
		return err
	}

	return os.Chmod(dst, srcinfo.Mode())
}

// Dir copies the folder inside root_directory
// to the backup_directory
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		log.Fatal("Error:", err)
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		log.Fatal("Error:", err)
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		log.Fatal("Error:", err)
		return err
	}

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				log.Fatal("Error:", err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				log.Fatal("Error:", err)
			}
		}
	}

	return nil
}

// Push copies files and directories from
// the root_directory to the backup_directory
func Push() string {
	rootDir := getrootdir.GetRootDir()
	bupDir := getbupdir.GetBupDir()
	Dir(rootDir, bupDir)

	return ""
}
