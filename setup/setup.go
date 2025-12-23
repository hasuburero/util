package setup

import (
	"errors"
	"strings"
)

type CmdArgs struct {
	args map[string][]string
}

var (
	Key_not_found_error = errors.New("Key not found error")
)

func ParseArgs(args []string) CmdArgs {
	new_args := CmdArgs{args: make(map[string][]string)}
	for _, value := range args {
		flag := strings.Contains(value, "=")
		if !flag {
			continue
		}
		slice := strings.Split(value, "=")
		new_args.args[slice[0]] = append(new_args.args[slice[0]], slice[1])
	}

	return new_args
}

func (self CmdArgs) GetArgs(key string) ([]string, error) {
	args, exists := self.args[key]
	if !exists {
		return nil, Key_not_found_error
	}

	return args, nil
}
