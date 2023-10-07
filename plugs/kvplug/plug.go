package kvplug

type KVHandler interface {
	GenerateKey(cname, name string) string
	GenerateValue()
}