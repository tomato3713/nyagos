package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var pwdHelp = cmdHelp{
	usage: "pwd -N (N:digit) -L -P",
	short: "Print the current working drive and directory.",
	long: `Print the current working drive and directory.
    Options:
        -N: N is digit. Print the N-previous directory.
        -L: Use PWD from environment, even if it contains symlinks. (default)
        -P: avoid symlinks.`,
}

func cmdPwd(ctx context.Context, cmd Param) (int, error) {
	physical := false
	if len(cmd.Args()) >= 2 {
		if cmd.Arg(1) == "-P" || cmd.Arg(1) == "-p" {
			physical = true
		} else if cmd.Arg(1) == "-L" || cmd.Arg(1) == "-l" {
			physical = false
		} else if i, err := strconv.ParseInt(cmd.Arg(1), 10, 0); err == nil && i < 0 {
			i += int64(len(cdHistory))
			if i < 0 {
				return errnoNoHistory, fmt.Errorf("pwd %s: too old history", cmd.Arg(1))
			}
			fmt.Fprintln(cmd.Out(), cdHistory[i])
			return 0, nil
		} else if cmd.Arg(1) == "-h" || cmd.Arg(1) == "/h" {
			fmt.Fprintf(cmd.Out(), pwdHelp.shortHelp())
			return 0, nil
		} else if cmd.Arg(1) == "-help" || cmd.Arg(1) == "/help" {
			fmt.Fprintf(cmd.Out(), pwdHelp.longHelp())
			return 0, nil
		}
	}
	wd, _ := os.Getwd()
	if physical {
		if _wd, err := filepath.EvalSymlinks(wd); err == nil {
			wd = _wd
		}
	}
	fmt.Fprintln(cmd.Out(), wd)
	return 0, nil
}
