package analyzer

import (
	"fmt"
)


type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}


type ParsingError struct {
	Path    string
	Message string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("erreur de parsing dans %s: %s", e.Path, e.Message)
}


func NewFileNotFoundError(path string, err error) *FileNotFoundError {
	return &FileNotFoundError{
		Path: path,
		Err:  err,
	}
}

func NewParsingError(path, message string) *ParsingError {
	return &ParsingError{
		Path:    path,
		Message: message,
	}
} 