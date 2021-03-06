package main

import (
	"io"
	"os"
	"strings"
	
)

type spaceEraser struct {
	r io.Reader
}

func (s spaceEraser) Read(b []byte) (int, error){
	x, err := s.r.Read(b)
	var str []byte
	
	for i := 0 ; i<=19; i++{
		if b[i] != 32 {
			str = append(str,b[i])
		}		
	}
	copy(b,str)	
	return x, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
