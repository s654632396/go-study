package main

import "fmt"

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	return 0, err
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 0, err
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 0, err
}
func (f *File) Close() (err error) {
	return err
}

type iFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (n int, err error)
	Close() (err error)
}

func main() {
	fmt.Println("vim-go")

	f := new(File)
	if f1, ok := f.(*iFile); ok {
		fmt.Println("f1 is impl iFile")
	}

}
