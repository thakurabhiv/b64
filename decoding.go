package main

import (
  "io"
  b64 "encoding/base64"
)

func decodeFile(r io.Reader, w io.Writer) error {
  decoder := b64.NewDecoder(b64.StdEncoding, r)
  _, err := io.Copy(w, decoder)
  if err != nil {
    return err
  }

  return nil
}
