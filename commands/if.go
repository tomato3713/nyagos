package commands

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/zetamatta/nyagos/shell"
)

func cmd_if(ctx context.Context, cmd *shell.Cmd) (int, error) {
	// if "xxx" == "yyy"
	args := cmd.Args
	not := false
	start := 1

	option := map[string]struct{}{}

	for len(args) >= 2 && strings.HasPrefix(args[1], "/") {
		option[strings.ToLower(args[1])] = struct{}{}
		args = args[1:]
		start++
	}

	if len(args) >= 2 && strings.EqualFold(args[1], "not") {
		not = true
		args = args[1:]
		start++
	}
	status := false
	if len(args) >= 4 && args[2] == "==" {
		if _, ok := option["/i"]; ok {
			status = strings.EqualFold(args[1], args[3])
		} else {
			status = (args[1] == args[3])
		}
		args = args[4:]
		start += 3
	} else if len(args) >= 3 && strings.EqualFold(args[1], "exist") {
		_, err := os.Stat(args[2])
		status = (err == nil)
		args = args[3:]
		start += 2
	} else if len(args) >= 3 && strings.EqualFold(args[1], "errorlevel") {
		num, num_err := strconv.Atoi(args[2])
		if num_err == nil {
			status = (shell.LastErrorLevel >= num)
		}
		start += 2
	}

	if not {
		status = !status
	}

	if len(args) > 0 {
		if status {
			subCmd, err := cmd.Clone()
			if err != nil {
				return 0, err
			}
			subCmd.Args = cmd.Args[start:]
			subCmd.RawArgs = cmd.RawArgs[start:]
			return subCmd.SpawnvpContext(ctx)
		}
	} else {
		stream, ok := ctx.Value("stream").(shell.Stream)
		if !ok {
			return 1, errors.New("not found stream")
		}
		thenBuffer := BufStream{}
		elseBuffer := BufStream{}
		elsePart := false

		save_prompt := os.Getenv("PROMPT")
		os.Setenv("PROMPT", "if>")
		nest := 1
		for {
			_, line, err := stream.ReadLine(ctx)
			if err != nil {
				break
			}
			args := shell.SplitQ(line)
			name := strings.ToLower(args[0])
			if _, ok := start_list[name]; ok {
				nest++
			} else if name == "end" {
				nest--
				if nest == 0 {
					break
				}
			} else if name == "else" {
				if nest == 1 {
					elsePart = true
					os.Setenv("PROMPT", "else>")
					continue
				}
			}
			if elsePart {
				elseBuffer.Add(line)
			} else {
				thenBuffer.Add(line)
			}
		}
		os.Setenv("PROMPT", save_prompt)

		if status {
			cmd.Loop(&thenBuffer)
		} else {
			cmd.Loop(&elseBuffer)
		}
	}
	return 0, nil
}
