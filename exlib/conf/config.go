package conf


type Config interface {
	GetValue(mark, name string) string
}