package sdp

type SourceGroupInfo struct {
	semantics string
	ssrcs     []uint32
}

func NewSourceGroupInfo(semantics string, ssrcs []uint32) *SourceGroupInfo {

	info := &SourceGroupInfo{
		semantics: semantics,
		ssrcs:     make([]uint32, len(ssrcs)),
	}

	copy(info.ssrcs, ssrcs)
	return info
}

func (s *SourceGroupInfo) Clone() *SourceGroupInfo {

	cloned := &SourceGroupInfo{
		semantics: s.semantics,
		ssrcs:     make([]uint32, len(s.ssrcs)),
	}

	copy(cloned.ssrcs, s.ssrcs)
	return cloned
}

func (s *SourceGroupInfo) GetSemantics() string {

	return s.semantics
}

func (s *SourceGroupInfo) GetSSRCs() []uint32 {

	return s.ssrcs
}
