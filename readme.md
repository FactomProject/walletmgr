Factom Wallet Manager
==========

This program exports the root seed of the current factom wallet in a portable format.  It can printed to paper and backed up securely.

Note, the seed was not intended to be a deterministic wallet, but can be used as such.

The seed serves as a pool of randomness. The randomness is pulled from to create both Factoid and Entry Credit keys. If only Factoid keys were created, then when restoring the addresses should not be problematic.  If EC and Factoid keys are interleaved, then it would make sense to restore many Factoid keys to find all of those with balances.  Rename the wallet, restore it again and create many EC keys to find the ones with balances.


## Creating the backup

Download and compile walletmgr.

Golang 1.5 or above must be installed, as walletmgr is not distributed as binaries.  It can also then be cross compiled or run in a VM.

To install run `go get -v -u github.com/FactomProject/walletmgr`

then run `walletmgr exportseed`

Neither fctwallet nor walletapp can be running at the same time as walletmgr.

Either fctwallet or walletapp must have been run and at least 1 address must have been created to backup the seed.


```
$ walletmgr exportseed
2016/01/19 17:46:56 read factom config file:  /home/brian/.factom/factomd.conf

Seed: sdrU9wonUEx47Ee1EjeXgeGtr8vrkEbd4GuQ1tfgRxp3woy2nBJQS9eqoa2WdAeiB3ga2DBEzr558GnSe4qpNBSZF6BNX5C
```

The sdxxx... value is a base58 encoded seed with a checksum.  Record this value, as it can be injected back into Factom to restore the private keys.




## Restore the backup

Currently there is no easy way to restore the key, but it can be done with enough effort.  It is a lot easier to make a digital backup, but paper backups have their benefits.


### Decode the seed

The seed is recorded in bitcoin base58check format, so it is effectively impossible for a typo to go undetected.

Currently, only a simple python script is needed to decode the address.  the base58 library is a prerequisite.

- Make sure pip is installed for your OS.  [Here](http://www.liquidweb.com/kb/how-to-install-pip-on-ubuntu-14-04-lts/) are some directions for Ubuntu.

- Install the base58 library with pip.  On ubuntu run `sudo pip install base58`.

- Next run the decoding script.  replace the example `seedToCheck` with your own.

```
import base58

seedToCheck = "sdrU9wonUEx47Ee1EjeXgeGtr8vrkEbd4GuQ1tfgRxp3woy2nBJQS9eqoa2WdAeiB3ga2DBEzr558GnSe4qpNBSZF6BNX5C"

try:
	decodedSeed = base58.b58decode_check(seedToCheck).encode("hex")
except:
	print "failed checksum"
else:
	if len(decodedSeed) != 132:
		print "too long or short"
	elif decodedSeed[:4] != "13dd":
		print "not a seed prefix"
	else:
		print "OK"
		print "secret seed is " + decodedSeed[4:]
```

it gives: `secret seed is e6d0083684c4decae94b9890b4c4ad3ca5f962f508d0641a4d5225117ebfc8ec29a0c58930d30424e0154ae6a88a3a63e32025f0e216f4d174bf81b17fff9e47`




### Make a new wallet with the seed

Install factomd and walletapp as normal

```
go get -v -u github.com/FactomProject/FactomCode/factomd
go get -v -u github.com/FactomProject/fctwallet
go get -v -u github.com/FactomProject/factom-cli
go get -v -u github.com/FactomProject/walletapp
```


Open the file ~/go/src/github.com/FactomProject/factoid/wallet/scwallet.go

Add a line forcing the seedhash to the one decoded above.

https://github.com/FactomProject/factoid/blob/666a00af6f151830cb8be97abf566239d4a11425/wallet/scwallet.go#L305

```
...

func (w *SCWallet) NewSeed(data []byte) {
	if len(data) == 0 {
		return
	} // No data, no change
	hasher := sha512.New()
	hasher.Write(data)
	seedhash := hasher.Sum(nil)
	seedhash, _ = hex.DecodeString ("e6d0083684c4decae94b9890b4c4ad3ca5f962f508d0641a4d5225117ebfc8ec29a0c58930d30424e0154ae6a88a3a63e32025f0e216f4d174bf81b17fff9e47") //added this line
	w.NextSeed = seedhash
	w.RootSeed = seedhash
	b := new(database.ByteStore)
	b.SetBytes(w.RootSeed)
...
```

also add this line `"encoding/hex"` to the import list at the top of the scwallet.go file.


Now compile the updated wallet which always creates new addresses with the seed.

`go install github.com/FactomProject/fctwallet`


Rename the existing wallet if `~/.factom/factoid_wallet_bolt.db` exists.  This will cause the wallet to make a new one when new addresses are made.

```
$ factom-cli newaddress fct recovered001
fct  =  FA3LRaa3oqC8F3wvecqo1D7WPPhG7opttoM67yCsQQLy1LkzYXUa
```

Keep making new addresses with new names, and it should create addresses you had backed up.



note: editing the same scwallet.go and creating new addresses will restore the same way when using walletapp as well.
