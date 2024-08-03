package main

import "golang.design/x/clipboard"

func writeToClipboard(b []byte) (<-chan struct{}, error) {
  err := clipboard.Init()
  if err != nil {
    return nil, err
  }

  return clipboard.Write(clipboard.FmtText, b), nil
}
