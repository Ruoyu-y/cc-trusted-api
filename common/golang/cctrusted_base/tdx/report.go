package tdx

import (
	"encoding/binary"
	"errors"

	"github.com/cc-api/cc-trusted-api/common/golang/cctrusted_base"
)

var _ cctrusted_base.Report = (*TdxReport)(nil)

type TdxReport struct {
	Quote *TdxQuote
}

func NewTdxReportFromBytes(b []byte) (*TdxReport, error) {
	t := &TdxReport{}
	err := t.InitFromBytes(b)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// IMRGroup implements cctrusted_base.Report.
func (t *TdxReport) IMRGroup() cctrusted_base.IMRGroup {
	group := cctrusted_base.IMRGroup{}
	group.MaxIndex = 3
	group.Group = make([]cctrusted_base.TcgDigest, 4)
	r0 := t.Quote.Body.Rtmr0
	group.Group[0] = NewRTMR(r0)
	r1 := t.Quote.Body.Rtmr1
	group.Group[1] = NewRTMR(r1)
	r2 := t.Quote.Body.Rtmr2
	group.Group[2] = NewRTMR(r2)
	r3 := t.Quote.Body.Rtmr3
	group.Group[3] = NewRTMR(r3)
	return group
}

// InitFromBytes implements cctrusted_base.Report.
func (t *TdxReport) InitFromBytes(b []byte) (err error) {
	t.Quote, err = NewTdxQuote(b)
	return err
}

// Dump implements cctrusted_base.Report
func (t *TdxReport) Dump(format cctrusted_base.QuoteDumpFormat) {
	t.Quote.Dump(format, "")
}

// Marshal can marshal the TdxReport structure into bytes
func (t *TdxReport) Marshal() ([]byte, error) {
	rawBytes := append(t.Quote.Header.raw.Binary, t.Quote.Body.raw.Binary...)
	sig_ecdsa, ok := t.Quote.Signature.(*TdxQuoteSignatureEcdsa256)
	if !ok {
		return []byte{}, errors.New("Invalid TDX Quote Signature.")
	}
	sig_size := make([]byte, 4)
	binary.LittleEndian.PutUint32(sig_size, uint32(len(sig_ecdsa.raw.Binary)))
	rawBytes = append(rawBytes, sig_size...)
	rawBytes = append(rawBytes, sig_ecdsa.raw.Binary...)
	return rawBytes, nil
}
