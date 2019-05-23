package gox12

// PathFinder : TODO
type PathFinder func(string, Segment) (string, bool, error)

// MakeMapFinder : TODO
func MakeMapFinder() PathFinder {
	var hardMap = map[string]string{
		"ISA": "/ISA_LOOP/ISA",
		"IEA": "/ISA_LOOP/IEA",
		"GS":  "/ISA_LOOP/GS_LOOP/GS",
		"GE":  "/ISA_LOOP/GS_LOOP/GE",
		"ST":  "/ISA_LOOP/GS_LOOP/ST_LOOP/ST",
		"SE":  "/ISA_LOOP/GS_LOOP/ST_LOOP/SE",
	}
	return func(rawpath string, s Segment) (string, bool, error) {
		segID := s.SegmentID
		p, ok := hardMap[segID]
		if ok {
			return p, ok, nil
		}
		return "", false, nil
	}
}
