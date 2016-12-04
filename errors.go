package sigfox

type RequestBodyError string

func (c RequestBodyError) Error() string {
	return string(c)
}

type ContentTypeError string

func (c ContentTypeError) Error() string {
	return string(c)
}

type MethodError string

func (c MethodError) Error() string {
	return string(c)
}
