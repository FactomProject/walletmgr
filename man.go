// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

const usage = `
walletmgr [subcommand]   
        
    exportseed              Export seed for future deterministic wallet re-creation
    help                    Print this help message
`

// man returns an usage error string for the specified sub command.
func man(s string) {
	m := map[string]string{
		"exportseed":    "walletmgr exportseed",
		"help":           usage,
		"default":        "More Help can be found by typing:\n\n  walletmgr help",
	}

	if m[s] != "" {
		errorln(m[s])
		return
	}

	errorln(m["default"])
}
