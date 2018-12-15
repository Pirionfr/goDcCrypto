package main

import (
	"flag"
	"fmt"
	"github.com/Pirionfr/goDcCrypto/crypto"
	"os"
)

func main() {
	// Subcommands
	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)
	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)

	// decrypt subcommand flag pointers
	decryptKeyPtr := decryptCommand.String("k", "", "master key. (Required)")
	decryptMsgPtr := decryptCommand.String("m", "", "message. (Required)")

	// encrypt subcommand flag pointers
	encryptKeyPtr := encryptCommand.String("k", "", "master key. (Required)")
	encryptMsgPtr := encryptCommand.String("m", "", "message. (Required)")

	if len(os.Args) < 2 {
		fmt.Println("decrypt or encrypt subcommand is required")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "encrypt":
		encryptCommand.Parse(os.Args[2:])
	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if encryptCommand.Parsed() {
		if *encryptKeyPtr == "" {
			encryptCommand.PrintDefaults()
			os.Exit(1)
		}
		if *encryptMsgPtr == "" {
			encryptCommand.PrintDefaults()
			os.Exit(1)
		}

		res, err := crypto.EncryptString(*encryptMsgPtr, *encryptKeyPtr)
		if err != nil {
			fmt.Printf("Error while encrypting message %s with master key %s\n", *encryptMsgPtr, *decryptKeyPtr)
			os.Exit(1)
		}
		fmt.Printf("%s\n", res)
	}

	if decryptCommand.Parsed() {
		if *decryptKeyPtr == "" {
			decryptCommand.PrintDefaults()
			os.Exit(1)
		}
		if *decryptMsgPtr == "" {
			decryptCommand.PrintDefaults()
			os.Exit(1)
		}

		res, err := crypto.DecryptString(*decryptMsgPtr, *decryptKeyPtr)
		if err != nil {
			fmt.Printf("Error while decrypting message %s with master key %s\n", *encryptMsgPtr, *decryptKeyPtr)
			os.Exit(1)
		}
		fmt.Printf("%s\n", res)
	}
}
