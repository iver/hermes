package lib

// Content interface
type Content interface {
	SetVariables(map[string]interface{}) error
	SetValue(string,string) error
}