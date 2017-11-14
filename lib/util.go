package lib

import (
	"bytes"
	"github.com/jackmanlabs/errors"
	"io"
	"os"
	"runtime"
	"strings"
)

func dump(input io.Reader) (io.Reader, error) {

	pc, _, _, _ := runtime.Caller(2)
	func_ := runtime.FuncForPC(pc)
	funcChunks := strings.Split(func_.Name(), "/")
	funcName := funcChunks[len(funcChunks)-1]

	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, input)
	if err != nil {
		return nil, errors.Stack(err)
	}

	f, err := os.Create(funcName + ".xml")
	if err != nil {
		return nil, errors.Stack(err)
	}
	defer f.Close()

	tee := io.TeeReader(input, f)

	_, err = io.Copy(buf, tee)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return buf, nil
}
