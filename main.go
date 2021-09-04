package main

import "fmt"

type Block struct {
  prevBlock *Block;
  hash string;
  index uint64;
}

func main() {
  block := Block{hash: "asd", index: 1}
  fmt.Printf("%v\n", block.index)
}
