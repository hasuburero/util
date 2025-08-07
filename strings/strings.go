package strings

import ()

func TrimSlice(args []string, trim string) []string {
	for i, v := range args {
		if v == trim {
			args = append(args[:i], args[i+1:]...)
		}
	}

	return args
}
