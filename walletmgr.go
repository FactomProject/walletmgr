// Copyright 2015 Factom Foundation

// Minimal seed export functionality
package main

import (
	"fmt"
	"github.com/btcsuitereleases/btcutil/base58"
	"github.com/FactomProject/factoid/wallet"
)

func main() {
    /* 
        Create and initialize a new SCWallet instance, and call GetSeed() at least once on it, 
        to have a random RootSeed set up (for testing purposes)
    */
    testWallet := new(wallet.SCWallet)
    testWallet.Init()
    testWallet.GetSeed()

    /*
        prefixing the 64 byte seeds with 0x13dd and then passing the result
        through the standard bitcoin base58check provides a range of addresses from 
        sdLGjhUDxGpiBEPRhTwysRYmxNQD6V48Aa84oVzfHvy6suim6qB6m3MCp8aHu1k1CNVLJdB8N9HtGR4NZTtFfp3mj591eA3
        to
        sdumH5tdzKD9zoBYceaCGX7ESMALunabDCPXHfuDjAWyu76vyaVZvvkrSjNn7ECRqzRUMuF1QgQ335DYVG9AF5agZpbqcQf
        the "sd" prefix is used to denote this as a seed value
    */
    firstPrefixByte := byte(0x13)
    secondPrefixByte := []byte{0xdd}

    /*
        The "version #" byte which is passed as a second parameter to .CheckEncode is appended to the 
        beginning of the byte-slice during the function operation. Therefore only the secondPrefixByte needs to
        be appended to the RootSeed bytes to craft the function's first parameter
    */
    base58EncodedResult := base58.CheckEncode(append(secondPrefixByte, testWallet.RootSeed...), firstPrefixByte)
    fmt.Printf("base58EncodedResult: %+v\n", base58EncodedResult)

}