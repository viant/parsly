package parsly

//Location represents a location
type Location struct {
	Path string
	Offset int
}
//NewLocation creates a location
func NewLocation(path string, offset int) *Location {
	return &Location{Path:path, Offset:offset}
}