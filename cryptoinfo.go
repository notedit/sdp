package sdp

type CryptoInfo struct {
	tag           int
	cipherSuite   string
	keyParams     string
	sessionParams string
}

func NewCryptoInfo(tag int, cipherSuite string, keyParams string, sessionParams string) *CryptoInfo {

	sdes := &CryptoInfo{
		tag:           tag,
		cipherSuite:   cipherSuite,
		keyParams:     keyParams,
		sessionParams: sessionParams,
	}

	return sdes
}

func (s *CryptoInfo) Clone() *CryptoInfo {
	sdes := &CryptoInfo{
		tag:           s.tag,
		cipherSuite:   s.cipherSuite,
		keyParams:     s.keyParams,
		sessionParams: s.sessionParams,
	}
	return sdes
}

func (s *CryptoInfo) GetTag() int {
	return s.tag
}

func (s *CryptoInfo) GetCipherSuite() string {
	return s.cipherSuite
}

func (s *CryptoInfo) GetKeyParams() string {
	return s.keyParams
}

func (s *CryptoInfo) GetSessionParams() string {
	return s.sessionParams
}
