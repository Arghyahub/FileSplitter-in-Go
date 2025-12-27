package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	BLOCK_SIZE := 1000
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("testing")
		log.Fatal("Atleast 2 params required:\nsplitter <NAME> <PARTS>")
		return
	}
	fileName := args[1]
	partNum := args[2]

	parts, err := strconv.Atoi(partNum)

	if err != nil {
		log.Fatal("Invalid parts, 2nd parameter should be numeric")
		return
	}

	originalFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file:\n%s", err)
		return
	}
	defer originalFile.Close()

	// Get the file size
	info, error := originalFile.Stat()
	if error != nil {
		log.Fatal("Unable to identify file size")
		return
	}

	openedFileSize := info.Size()

	partFileSize := (openedFileSize + int64(parts)) / int64(parts)
	buff := make([]byte, 1000)

	// For each part
	for i := 0; i < parts; i++ {
		newFileName := fmt.Sprintf("%s-part%d", fileName, i+1)

		newFile, err := os.OpenFile(newFileName, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatalf("Unable to create part file : %s", newFileName)
			return
		}

		// reading buffer by buffer
		for j := 0; j < int(partFileSize); j += BLOCK_SIZE {
			n, err := originalFile.Read(buff)
			if err != nil {
				log.Fatalf("Unable to read block for part : %s", newFileName)
				return
			}
			_, err = newFile.Write(buff[:n])
			if err != nil {
				log.Fatalf("Unable to write block for part : %s", newFileName)
				return
			}
		}
		newFile.Close()
	}
}
