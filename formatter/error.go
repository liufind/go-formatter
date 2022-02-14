package formatter

type fError string

func (f fError) Error() string {
	return string(f)
}
