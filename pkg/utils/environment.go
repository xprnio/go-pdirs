package utils

type Environment interface {
	LookupEnv(string) (string, bool)
}
