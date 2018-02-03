package main

import (
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"os"

	"github.com/Bren2010/proquint"
	"github.com/akamensky/base58"
)

var base58flag = flag.Bool("base58", false, "Decode and Encode with base58")
var crc32flag = flag.Bool("crc32", false, "Prepend and use a CRC32 checksum")

func encode(data string) (err error) {
	// Encoding
	if len(data)%2 != 0 {
		data = data + `-`
	}
	dataBytes := []byte(data)
	// Decode with base58 first
	if *base58flag {
		dataBytes, err = base58.Decode(data)
		if err != nil {
			return errors.New("Unable to decode with base58")
		}
	}
	// Prefix a CRC32 checksum
	if *crc32flag {
		crc := crc32.ChecksumIEEE(dataBytes)
		crcBytes := []byte{
			byte(crc >> 24),
			byte(crc >> 16),
			byte(crc >> 8),
			byte(crc),
		}
		dataBytes = append(crcBytes, dataBytes...)
	}
	// Encode with proquint
	fmt.Println(proquint.Encode(dataBytes))
	return nil
}

func decode(data string) error {
	var crc uint32
	// Decoding
	// First strip proquint
	dataBytes := proquint.Decode(data)
	data = string(dataBytes)
	// Check CRC
	if *crc32flag {
		crc = uint32(dataBytes[0])<<24 | uint32(dataBytes[1])<<16 | uint32(dataBytes[2])<<8 | uint32(dataBytes[3])
		dataBytes = dataBytes[4:]
	}
	// Encode base58 data
	if *base58flag {
		data = base58.Encode(dataBytes)
	}

	if *crc32flag {
		crc2 := crc32.ChecksumIEEE(dataBytes)
		if crc != crc2 {
			fmt.Fprintf(os.Stderr, "Checksum fails\n")
			os.Exit(1)
		}
	}
	fmt.Println(data)
	return nil
}

func main() {
	flag.Parse()
	var err error
	data := ""
	if len(flag.Args()) > 0 {
		data = flag.Arg(0)
	} else {
		length, err := fmt.Scanln(&data)
		if err != nil {
			panic(err)
		}
		if length == 0 {
			return
		}
	}

	if ok, _ := proquint.IsProquint(data); ok {
		err = decode(data)
	} else {
		err = encode(data)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
