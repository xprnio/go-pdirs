package internal

import "os"

type environ struct{}

func Environ() *environ {
	return &environ{}
}

func (_ *environ) LookupEnv(env string) (string, bool) {
	return os.LookupEnv(env)
}
