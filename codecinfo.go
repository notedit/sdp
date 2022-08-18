package sdp

import (
	"strconv"
	"strings"
)

type CodecInfo struct {
	codec   string
	payload int
	rtx     int
	params  map[string]string
	rtcpfbs []*RTCPFeedbackInfo
	rate    int
}

func NewCodecInfo(codec string, payload int, rate int) *CodecInfo {

	codecInfo := &CodecInfo{
		codec:   codec,
		payload: payload,
		rtx:     0,
		params:  map[string]string{},
		rtcpfbs: []*RTCPFeedbackInfo{},
		rate:    rate,
	}

	return codecInfo
}

func (c *CodecInfo) Clone() *CodecInfo {

	codecInfo := &CodecInfo{
		codec:   c.codec,
		payload: c.payload,
		rtx:     c.rtx,
		params:  make(map[string]string),
		rtcpfbs: []*RTCPFeedbackInfo{},
		rate:    c.rate,
	}

	for key, param := range c.params {
		codecInfo.params[key] = param
	}

	for _, rtcpfb := range c.rtcpfbs {
		copyfb := rtcpfb.Clone()
		codecInfo.rtcpfbs = append(codecInfo.rtcpfbs, copyfb)
	}

	return codecInfo
}

func (c *CodecInfo) SetRTX(rtx int) {
	c.rtx = rtx
}

func (c *CodecInfo) GetPayload() int {
	return c.payload
}

func (c *CodecInfo) SetPayload(payload int) {
	c.payload = payload
}

func (c *CodecInfo) GetParams() map[string]string {
	return c.params
}

func (c *CodecInfo) HasParam(key string) bool {
	_, ok := c.params[key]
	return ok
}

func (c *CodecInfo) GetCodec() string {
	return c.codec
}

func (c *CodecInfo) GetRate() int {
	return c.rate
}

func (c *CodecInfo) SetRate(rate int) {
	c.rate = rate
}

func (c *CodecInfo) GetParam(key string) string {
	if c.HasParam(key) {
		return c.params[key]
	}
	return ""
}

func (c *CodecInfo) AddParam(key, value string) {

	c.params[key] = value
}

func (c *CodecInfo) AddParams(params map[string]string) {

	for key, param := range params {
		c.params[key] = param
	}
}

func (c *CodecInfo) GetRTX() int {
	return c.rtx
}

func (c *CodecInfo) HasRTX() bool {
	if c.rtx != 0 {
		return true
	}
	return false
}

func (c *CodecInfo) AddRTCPFeedback(rtcpfb *RTCPFeedbackInfo) {

	c.rtcpfbs = append(c.rtcpfbs, rtcpfb)
}

func (c *CodecInfo) GetRTCPFeedbacks() []*RTCPFeedbackInfo {

	return c.rtcpfbs
}

func CodecMapFromNames(names []string, rtx bool, rtcpfbs []*RTCPFeedbackInfo) map[string]*CodecInfo {

	codecs := map[string]*CodecInfo{}

	basePt := 96

	for _, name := range names {

		var pt int
		var rate int
		params := strings.Split(name, ";")
		encoding := strings.Split(params[0], "/")
		codecName := strings.TrimSpace(strings.ToLower(encoding[0]))
		if len(encoding) == 2 {
			rate, _ = strconv.Atoi(encoding[1])
		}
		if codecName == "pcmu" {
			pt = 0
		} else if codecName == "pcma" {
			pt = 8
		} else {
			basePt++
			pt = basePt
		}

		codec := NewCodecInfo(codecName, pt, rate)

		if rtx && codecName != "ulpfec" && codecName != "flexfec-03" && codecName != "red" {
			basePt++
			codec.SetRTX(basePt)
		}

		if rtcpfbs != nil {
			for _, rtcpfb := range rtcpfbs {
				codec.AddRTCPFeedback(rtcpfb)
			}
		}

		if len(params) > 1 {
			params = params[1:]
			for _, param := range params {
				values := strings.Split(param, "=")
				if len(values) < 2 {
					continue
				}
				codec.AddParam(values[0], values[1])
			}
		}

		codecs[codecName] = codec
	}

	return codecs
}
