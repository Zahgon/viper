package viper

type FileLookupError interface {
	error

	fileLookup()
}

type ConfigFileNotFoundError struct {
	locations []string
	name      string
}

func (e ConfigFileNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ConfigFileNotFoundError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type FileNotFoundFromSearchError struct {
	locations []string
	name      string
}

func (e FileNotFoundFromSearchError) fileLookup() { _ = "STUB: not implemented"; return }

func (e FileNotFoundFromSearchError) Error() string { _ = "STUB: not implemented"; return "" }

type FileNotFoundError struct {
	err  error
	path string
}

func (e FileNotFoundError) fileLookup() { _ = "STUB: not implemented"; return }

func (e FileNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

type ConfigFileAlreadyExistsError string

func (e ConfigFileAlreadyExistsError) Error() string { _ = "STUB: not implemented"; return "" }

type ConfigMarshalError struct {
	err error
}

func (e ConfigMarshalError) Error() string { _ = "STUB: not implemented"; return "" }

type UnsupportedConfigError string

func (str UnsupportedConfigError) Error() string { _ = "STUB: not implemented"; return "" }
