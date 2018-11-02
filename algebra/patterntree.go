package algebra

// PatternTree results from parsing a groupGraphPattern
type PatternTree struct {
	Mode     string
	Children []*PatternTree
	Blocks   []*Block
}
