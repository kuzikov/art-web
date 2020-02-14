package utils

//Params - this is stucture of all prog ascii-art
type Params struct {
	InitColorState   string
	InitBgColorState string
	FileToOutput     string
	FileToOpen       string
	StrInput         string
	Font             string
	Textpos          string
	Slice            string
	ColorOfSlice     string
	ColorOfInterval  string
	SliceBool        bool
	Interval         bool
	Char             bool
	Output           bool
	Reverse          bool
	IsCorrectInput   bool
	Color            bool
	Align            bool
	Pipe             bool
	IntArr           []int
	CharArr          []Char
}

//Char struct - spectial struct for use colors by symbols
type Char struct {
	Symbol rune
	Color  string
}
