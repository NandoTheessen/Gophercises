package reader 

import (
	"io/ioutil"
	"fmt"
)

// ProcessFile takes a filename as argument and returns a bytearray with the 
// contents of that file or an error
func ProcessFile(filename string) (contents []byte, err error) {
	contents, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return contents, nil
}