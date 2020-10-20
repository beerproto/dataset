package commands


type Output string

const (
	CsvExt  = ".csv"
	JsonExt = ".json"

	TTY  = Output("tty")
	FILE = Output("file")
)

