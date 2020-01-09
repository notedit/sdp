package sdp

type SDESInfo struct {
	tag           int
	cipherSuite   string
	keyParams     string
	sessionParams string
}

func NewSDESInfo(tag int, cipherSuite string, keyParams string, sessionParams string) *SDESInfo {

	sdes := &SDESInfo{
		tag:           tag,
		cipherSuite:   cipherSuite,
		keyParams:     keyParams,
		sessionParams: sessionParams,
	}

	return sdes
}

func (s *SDESInfo) Clone() *SDESInfo {
	sdes := &SDESInfo{
		tag:           s.tag,
		cipherSuite:   s.cipherSuite,
		keyParams:     s.keyParams,
		sessionParams: s.sessionParams,
	}
	return sdes
}

func (s *SDESInfo) GetTag() int {
	return s.tag
}

func (s *SDESInfo) GetCipherSuite() string {
	return s.cipherSuite
}

func (s *SDESInfo) GetKeyParams() string {
	return s.keyParams
}

func (s *SDESInfo) GetSessionParams() string {
	return s.sessionParams
}
