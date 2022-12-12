package main

import (
	"fmt"
	"log"

	"github.com/aurora-is-near/stream-backup/chunks"
	"github.com/aurora-is-near/stream-backup/messagebackup"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	backup := chunks.Chunks{
		Dir:             "mybackup",
		ChunkNamePrefix: "mystream_",
	}
	if err := backup.Open(); err != nil {
		log.Fatal(err)
	}
	defer backup.CloseReader()

	if err := backup.SeekReader(50); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		seq, data, err := backup.ReadNext()
		if err != nil {
			log.Fatal(err)
		}

		var msg messagebackup.MessageBackup
		if err = msg.UnmarshalVT(data); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d: %v\n", seq, spew.Sdump(&msg))
	}
}
