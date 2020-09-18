package values

import (
	"io/ioutil"
	"os"
)

func NewFile(file *os.File) Value {
	return NewFileMode(file, os.O_RDONLY, 0)
}

func NewFileMode(file *os.File, flag int, perm os.FileMode) Value {
	filename := new(string)
	typ := "file"
	return Value{
		setter: func(val string) error {
			*filename = val
			var err error
			file, err = os.OpenFile(val, flag, perm)
			if err != nil {
				return NewInvalidValue(WithType(typ), WithValue(val), WithCause(err))
			}
			return nil
		},
		stringer: func() string {
			return *filename
		},
		typ: func() string {
			return typ
		},
	}
}

func NewFileContents(contents *[]byte) Value {
	filename := new(string)
	typ := "file-contents"
	return Value{
		setter: func(val string) error {
			*filename = val
			var err error
			*contents, err = ioutil.ReadFile(val)
			if err != nil {
				return NewInvalidValue(WithType(typ), WithValue(val), WithCause(err))
			}
			return nil
		},
		stringer: func() string {
			return *filename
		},
		typ: func() string {
			return typ
		},
	}
}
