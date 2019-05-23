package gox12

// X12PathFinder : Given the current location, find the path of the new segment
type X12PathFinder interface {
	FindNext(x12Path string, segment Segment) (foundPath string, found bool, err error)
}

// HeaderPathFinder : Hardcoded lookups for standard X12 structure wrappers
type HeaderPathFinder struct {
	hardMap map[string]string
}

// NewHeaderMapFinder : TODO
func NewHeaderMapFinder() *HeaderPathFinder {
	f := new(HeaderPathFinder)
	f.hardMap = map[string]string{
		"ISA": "/ISA_LOOP/ISA",
		"IEA": "/ISA_LOOP/IEA",
		"GS":  "/ISA_LOOP/GS_LOOP/GS",
		"GE":  "/ISA_LOOP/GS_LOOP/GE",
		"ST":  "/ISA_LOOP/GS_LOOP/ST_LOOP/ST",
		"SE":  "/ISA_LOOP/GS_LOOP/ST_LOOP/SE",
	}
	return f
}

// FindNext : TODO
func (finder *HeaderPathFinder) FindNext(x12Path string, segment Segment) (foundPath string, found bool, err error) {
	segID := segment.SegmentID
	p, ok := finder.hardMap[segID]
	if ok {
		return p, ok, nil
	}
	return "", false, nil
}

// FirstMatchPathFinder : TODO
type FirstMatchPathFinder struct {
	Finders []X12PathFinder
}

// NewFirstMatchPathFinder : TODO
func NewFirstMatchPathFinder(finder ...X12PathFinder) *FirstMatchPathFinder {
	f := new(FirstMatchPathFinder)
	f.Finders = make([]X12PathFinder, 0)
	for _, fn := range finder {
		f.Finders = append(f.Finders, fn)
	}
	return f
}

// FindNext : TODO
func (finder *FirstMatchPathFinder) FindNext(x12Path string, segment Segment) (foundPath string, found bool, err error) {
	for _, f2 := range finder.Finders {
		res, ok, err := f2.FindNext(x12Path, segment)
		if ok {
			return res, ok, err
		}
	}
	return "", false, nil
}

// type PathFinder func(string, Segment) (string, bool, error)

// EmptyPath : TODO
type EmptyPath struct {
	Path string
}

//func (e *EmptyPath) Run2 PathFinder {
//    return "", true, nil
//}

// this is the method signature
// need to close lookup maps
func findPath(rawpath string, seg Segment) (foundPath string, ok bool, err error) {
	return "", true, nil
}

// segMatcher is the function signature for segment matcher
// is the segment "matched"
type segMatcher func(seg Segment) bool

// segmentMatchBySegmentID matches a segment only by the segment ID
func segmentMatchBySegmentID(segmentID string) segMatcher {
	return func(seg Segment) bool {
		return seg.SegmentID == segmentID
	}
}

// segmentMatchIDByPath matches a segment by the segment ID and the ID value of the
// element at the x12path
func segmentMatchIDByPath(segmentID string, x12path string, idValue string) segMatcher {
	return func(seg Segment) bool {
		v, found, _ := seg.GetValue(x12path)
		return seg.SegmentID == segmentID && found && v == idValue
	}
}

// segmentMatchIdByPath matches a segment by the segment ID and one of the ID value of the
// element at the x12path
func segmentMatchIDListByPath(segmentID string, x12path string, idList []string) segMatcher {
	return func(seg Segment) bool {
		v, found, _ := seg.GetValue(x12path)
		x := stringInSlice(v, idList)
		return seg.SegmentID == segmentID && found && x
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
