package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/schollz/progressbar/v3"
)

var (
  encode, decode, isStr, isFile, copyToClipboard bool
  outputFile string
)

func main() {
  flag.BoolVar(&encode, "e", true, "Encode mode")
  flag.BoolVar(&decode, "d", false, "Decode Mode")
  flag.BoolVar(&isStr, "s", true, "String as input")
  flag.BoolVar(&isFile, "f", false, "File as input")

  if strings.ToLower(runtime.GOOS) != "linux" {
    flag.BoolVar(&copyToClipboard, "c", false, "Copy output to clipboard")
  }
  flag.StringVar(&outputFile, "o", "", "Output file name")

  flag.Parse()
  validateFlags()

  // get the tail
  args := flag.Args()
  if len(args) == 0 {
    fmt.Println("Missing input string/file path")
    os.Exit(1)
  }

  tail := args[0]
  inputR, err := getInputReader(tail)
  exitIfError(err)

  outputW, err := getOutputWriter()
  exitIfError(err)

  defer func() {
    if c, ok := inputR.(io.Closer); ok {
      c.Close()
    }

    if (outputW == os.Stdout) {
      // don't close stdout
      return
    }

    if c, ok := outputW.(io.Closer); ok {
      c.Close()
    }
  }()

  if decode {
    err := decodeFile(inputR, outputW)
    exitIfError(err)
  } else { // encode
    err := encodeFile(inputR, outputW)
    exitIfError(err)
  }

  // copy output to clipboard
  if copyToClipboard {
    b := outputW.(*bytes.Buffer)
    _, err := writeToClipboard(b.Bytes())
    exitIfError(err)

    fmt.Println("Output copied to clipboard")
  } else {
    fmt.Println()
  }
}

func getInputReader(tail string) (io.Reader, error) {
  if isFile {
    r, err := os.Open(tail)
    if err != nil {
      return nil, err
    }

    return r, nil
  }

  return strings.NewReader(tail), nil
}

func getOutputWriter() (io.Writer, error) {
  if len(outputFile) > 0 {
    w, err := os.Create(outputFile)
    if err != nil {
      return nil, err
    }

    pb := progressbar.DefaultBytes(-1, "Progress")
    return io.MultiWriter(w, pb), nil
  }

  var w io.Writer
  if copyToClipboard {
    w = new(bytes.Buffer)
  } else {
    w = os.Stdout
  }
  return w, nil
}

func exitIfError(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func validateFlags() {
  if (!encode && !decode) {
    fmt.Println("Mode should be either encode or decode")
    flag.Usage()
    os.Exit(1)
  }

  if (!isStr && !isFile) {
    fmt.Println("Input type must be either str or file")
    flag.Usage()
    os.Exit(1)
  }
} 
