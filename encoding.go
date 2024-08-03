package main

import (
  "io"
  b64 "encoding/base64"
)

func encodeFile(r io.Reader, w io.Writer) error {
  encoder := b64.NewEncoder(b64.StdEncoding, w)
  defer encoder.Close()

  _, err := io.Copy(encoder, r)
  if err != nil {
    return err
  }

  return nil
}
