package cmd

import (
	"os"
	"io/ioutil"
)

func captureStdio(f func()) (stdout string, stderr string, err error) {
	rStdout, wStdout, err := os.Pipe()
	defer wStdout.Close()
	if err != nil {
		return
	}
	rStderr, wStderr, err := os.Pipe()
	defer wStderr.Close()
	if err != nil {
		return
	}

	orgStdout := os.Stdout
	orgStderr := os.Stderr
	os.Stdout = wStdout
	os.Stderr = wStderr

	f()

	os.Stdout = orgStdout
	wStdout.Close()
	os.Stderr = orgStderr
	wStderr.Close()

	stdoutBuf, err := ioutil.ReadAll(rStdout)
	if err != nil {
		return
	}
	stderrBuf, err := ioutil.ReadAll(rStderr)
	if err != nil {
		return
	}

	return string(stdoutBuf), string(stderrBuf), nil
}
