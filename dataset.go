package hdf5

import "os"

// File constants
const (
	F_ACC_RDONLY  int = 0x0000 // absence of rdwr => rd-only
	F_ACC_RDWR    int = 0x0001 // open for read and write
	F_ACC_TRUNC   int = 0x0002 // Truncate file, if it already exists, erasing all data previously stored in the file.
	F_ACC_EXCL    int = 0x0004 // Fail if file already exists.
	F_ACC_DEBUG   int = 0x0008 // print debug info
	F_ACC_CREAT   int = 0x0010 // create non-existing files
	F_ACC_DEFAULT int = 0xffff // value passed to set_elink_acc_flags to cause flags to be taken from the parent file
)

type Scope int

const (
	F_SCOPE_LOCAL  Scope = 0 // specified file handle only.
	F_SCOPE_GLOBAL Scope = 1 // entire virtual file.
)

func CreateFile(name string, flags int) (*os.File, error) {
	return nil, nil
}
