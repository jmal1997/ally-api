package ally_api

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
	funcChunks := strings.Split(func_.Name(), ".")
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

	_, err = io.Copy(buf, input)
	if err != nil {
		return nil, errors.Stack(err)
	}

	tmp := buf.Bytes()
	_, err = io.Copy(f, buf)
	if err != nil {
		return nil, errors.Stack(err)
	}
	buf = bytes.NewBuffer(tmp)

	return buf, nil
}
