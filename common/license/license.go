// Package license helps manage commercial licenses and check if they are valid for the version of unidoc used.
package license

import (
	_db "bytes"
	_da "compress/gzip"
	_d "crypto"
	_cf "crypto/rand"
	_cg "crypto/rsa"
	_fe "crypto/sha256"
	_ac "crypto/sha512"
	_e "crypto/x509"
	_f "encoding/base64"
	_b "encoding/hex"
	_df "encoding/json"
	_gg "encoding/pem"
	_aa "errors"
	_gd "fmt"
	_ab "io"
	_g "log"
	_ad "regexp"
	_adf "strings"
	_c "time"

	_gc "gitee.com/yu_sheng/gooffice/common"
)

const (
	LicenseTierUnlicensed = "unlicensed"
	LicenseTierCommunity  = "\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y"
	LicenseTierIndividual = "\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	LicenseTierBusiness   = "\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073"
)

var _ddf = false

func _aac(_cb string, _ec string) ([]byte, error) {
	_bf := "\u000a\u002b\u000a"
	_ggf := "\u000d\u000a\u002b\r\u000a"
	_gf := _adf.Index(_ec, _bf)
	if _gf == -1 {
		_gf = _adf.Index(_ec, _ggf)
		if _gf == -1 {
			return nil, _gd.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072")
		}
	}
	_gfe := _ec[:_gf]
	_ee := _gf + len(_bf)
	_cga := _ec[_ee:]
	if _gfe == "" || _cga == "" {
		return nil, _gd.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065")
	}
	_gfc, _cbe := _f.StdEncoding.DecodeString(_gfe)
	if _cbe != nil {
		return nil, _gd.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c")
	}
	_bc, _cbe := _f.StdEncoding.DecodeString(_cga)
	if _cbe != nil {
		return nil, _gd.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_bd, _ := _gg.Decode([]byte(_cb))
	if _bd == nil {
		return nil, _gd.Errorf("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_dac, _cbe := _e.ParsePKIXPublicKey(_bd.Bytes)
	if _cbe != nil {
		return nil, _cbe
	}
	_ge := _dac.(*_cg.PublicKey)
	if _ge == nil {
		return nil, _gd.Errorf("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_be := _ac.New()
	_be.Write(_gfc)
	_bec := _be.Sum(nil)
	_cbe = _cg.VerifyPKCS1v15(_ge, _d.SHA512, _bec, _bc)
	if _cbe != nil {
		return nil, _cbe
	}
	return _gfc, nil
}

// TypeToString returns a string representation of the license type.
func (_dge *LicenseKey) TypeToString() string {
	if _dge.Tier == LicenseTierUnlicensed {
		return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	}
	if _dge.Tier == LicenseTierCommunity {
		return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	}
	if _dge.Tier == LicenseTierIndividual || _dge.Tier == "\u0069\u006e\u0064i\u0065" {
		return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	}
	return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073"
}

type LegacyLicense struct {
	Name        string
	Signature   string `json:",omitempty"`
	Expiration  _c.Time
	LicenseType LegacyLicenseType
}

// LegacyLicenseType is the type of license
type LegacyLicenseType byte

const (
	_ff = "\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
	_cc = "\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
)

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey(s string) error {
	_bb := _ad.MustCompile("\u005c\u0073")
	s = _bb.ReplaceAllString(s, "")
	var _ba _ab.Reader
	_ba = _adf.NewReader(s)
	_ba = _f.NewDecoder(_f.RawURLEncoding, _ba)
	_ba, _aacd := _da.NewReader(_ba)
	if _aacd != nil {
		return _aacd
	}
	_dfab := _df.NewDecoder(_ba)
	_bfa := &LegacyLicense{}
	if _gcd := _dfab.Decode(_bfa); _gcd != nil {
		return _gcd
	}
	if _gcb := _bfa.Verify(_ffa); _gcb != nil {
		return _aa.New("\u006c\u0069\u0063en\u0073\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006e\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _bfa.Expiration.Before(_gc.ReleasedAt) {
		return _aa.New("\u006ci\u0063e\u006e\u0073\u0065\u0020\u0065\u0078\u0070\u0069\u0072\u0065\u0064")
	}
	_bgd := _c.Now().UTC()
	_fag := LicenseKey{}
	_fag.CreatedAt = _bgd
	_fag.CustomerId = "\u004c\u0065\u0067\u0061\u0063\u0079"
	_fag.CustomerName = _bfa.Name
	_fag.Tier = LicenseTierBusiness
	_fag.ExpiresAt = _bfa.Expiration
	_fag.CreatorName = "\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u0073\u0075p\u0070\u006f\u0072\u0074"
	_fag.CreatorEmail = "\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0040\u0075\u006e\u0069\u0064o\u0063\u002e\u0069\u006f"
	_fag.GoOffice = true
	_gccb = &_fag
	return nil
}

// LicenseKey represents a license key for GoOffice.
type LicenseKey struct {
	LicenseId    string  `json:"license_id"`
	CustomerId   string  `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	Tier         string  `json:"tier"`
	CreatedAt    _c.Time `json:"-"`
	CreatedAtInt int64   `json:"created_at"`
	ExpiresAt    _c.Time `json:"-"`
	ExpiresAtInt int64   `json:"expires_at"`
	CreatorName  string  `json:"creator_name"`
	CreatorEmail string  `json:"creator_email"`
	UniPDF       bool    `json:"unipdf"`
	GoOffice     bool    `json:"gooffice"`
	Trial        bool    `json:"trial"`
}

var _ffa *_cg.PublicKey

// SetLicenseKey sets and validates the license key.
func SetLicenseKey(content string, customerName string) error {
	if _ddf {
		return nil
	}
	_fb, _ffe := _de(content)
	if _ffe != nil {
		return _ffe
	}
	if _adf.ToLower(_fb.CustomerName) != _adf.ToLower(customerName) {
		return _gd.Errorf("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'", customerName, _fb.CustomerName)
	}
	_ffe = _fb.Validate()
	if _ffe != nil {
		return _ffe
	}
	_gccb = &_fb
	return nil
}
func GetLicenseKey() *LicenseKey {
	if _gccb == nil {
		return nil
	}
	_ag := *_gccb
	return &_ag
}

func _de(_ccd string) (LicenseKey, error) {
	var _gge LicenseKey
	_gcf, _dd := _ea(_ff, _cc, _ccd)
	if _dd != nil {
		return _gge, _dd
	}
	_acg, _dd := _aac(_dfa, _gcf)
	if _dd != nil {
		return _gge, _dd
	}
	_dd = _df.Unmarshal(_acg, &_gge)
	if _dd != nil {
		return _gge, _dd
	}
	_gge.CreatedAt = _c.Unix(_gge.CreatedAtInt, 0)
	_gge.ExpiresAt = _c.Unix(_gge.ExpiresAtInt, 0)
	return _gge, nil
}

// MakeUnlicensedKey returns an unlicensed key.
func MakeUnlicensedKey() *LicenseKey {
	_eed := LicenseKey{}
	_eed.CustomerName = "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	_eed.Tier = LicenseTierUnlicensed
	_eed.CreatedAt = _c.Now().UTC()
	_eed.CreatedAtInt = _eed.CreatedAt.Unix()
	return &_eed
}

// Verify verifies a license by checking the license content and signature
// against a public key.
func (_eef LegacyLicense) Verify(pubKey *_cg.PublicKey) error {
	_cgd := _eef
	_cgd.Signature = ""
	_fff := _db.Buffer{}
	_ggga := _df.NewEncoder(&_fff)
	if _ef := _ggga.Encode(_cgd); _ef != nil {
		return _ef
	}
	_aaa, _gfg := _b.DecodeString(_eef.Signature)
	if _gfg != nil {
		return _gfg
	}
	_eg := _fe.Sum256(_fff.Bytes())
	_gfg = _cg.VerifyPKCS1v15(pubKey, _d.SHA256, _eg[:], _aaa)
	return _gfg
}
func _ca(_ffd string, _fg []byte) (string, error) {
	_af, _ := _gg.Decode([]byte(_ffd))
	if _af == nil {
		return "", _gd.Errorf("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	_gbc, _fea := _e.ParsePKCS1PrivateKey(_af.Bytes)
	if _fea != nil {
		return "", _fea
	}
	_cag := _ac.New()
	_cag.Write(_fg)
	_dg := _cag.Sum(nil)
	_acf, _fea := _cg.SignPKCS1v15(_cf.Reader, _gbc, _d.SHA512, _dg)
	if _fea != nil {
		return "", _fea
	}
	_ae := _f.StdEncoding.EncodeToString(_fg)
	_ae += "\u000a\u002b\u000a"
	_ae += _f.StdEncoding.EncodeToString(_acf)
	return _ae, nil
}
func (_fc *LicenseKey) isExpired() bool { return _fc.getExpiryDateToCompare().After(_fc.ExpiresAt) }

var _fee = _c.Date(2010, 1, 1, 0, 0, 0, 0, _c.UTC)

func init() {
	_aef, _gcc := _b.DecodeString(_bg)
	if _gcc != nil {
		_g.Fatalf("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073", _gcc)
	}
	_bcb, _gcc := _e.ParsePKIXPublicKey(_aef)
	if _gcc != nil {
		_g.Fatalf("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073", _gcc)
	}
	_ffa = _bcb.(*_cg.PublicKey)
}
func (_afg *LicenseKey) getExpiryDateToCompare() _c.Time {
	if _afg.Trial {
		return _c.Now().UTC()
	}
	return _gc.ReleasedAt
}

// Validate returns an error if the licenseis invalid, nil otherwise.
func (_cef *LicenseKey) Validate() error {
	if len(_cef.LicenseId) < 10 {
		return _gd.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064")
	}
	if len(_cef.CustomerId) < 10 {
		return _gd.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064")
	}
	if len(_cef.CustomerName) < 1 {
		return _gd.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065")
	}
	if _fee.After(_cef.CreatedAt) {
		return _gd.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _cef.CreatedAt.After(_cef.ExpiresAt) {
		return _gd.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074")
	}
	if _cef.isExpired() {
		return _gd.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064")
	}
	if len(_cef.CreatorName) < 1 {
		return _gd.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065")
	}
	if len(_cef.CreatorEmail) < 1 {
		return _gd.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c")
	}
	if !_cef.GoOffice {
		return _gd.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0054\u0068\u0069s\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u006be\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020f\u006f\u0072\u0020\u0055\u006e\u0069O\u0066\u0066\u0069c\u0065\u002e")
	}
	return nil
}

var _gccb = MakeUnlicensedKey()

const _dfa = "\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n"

func _ea(_abe string, _cd string, _aeb string) (string, error) {
	_eaa := _adf.Index(_aeb, _abe)
	if _eaa == -1 {
		return "", _gd.Errorf("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ce := _adf.Index(_aeb, _cd)
	if _ce == -1 {
		return "", _gd.Errorf("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ga := _eaa + len(_abe) + 1
	return _aeb[_ga : _ce-1], nil
}

const _bg = "\u0033\u0030\u0035\u0063\u0033\u0030\u0030\u00640\u0036\u0030\u0039\u0032\u0061\u0038\u00364\u0038\u0038\u0036\u0066\u0037\u0030d\u0030\u0031\u0030\u0031\u0030\u00310\u0035\u0030\u0030\u0030\u0033\u0034\u0062\u0030\u0030\u0033\u0030\u00348\u0030\u0032\u0034\u0031\u0030\u0030\u0062\u0038\u0037\u0065\u0061\u0066\u0062\u0036\u0063\u0030\u0037\u0034\u0039\u0039\u0065\u0062\u00397\u0063\u0063\u0039\u0064\u0033\u0035\u0036\u0035\u0065\u0063\u00663\u0031\u0036\u0038\u0031\u0039\u0036\u0033\u0030\u0031\u0039\u0030\u0037c\u0038\u0034\u0031\u0061\u0064\u0064c6\u0036\u0035\u0030\u0038\u0036\u0062\u0062\u0033\u0065\u0064\u0038\u0065\u0062\u0031\u0032\u0064\u0039\u0064\u0061\u0032\u0036\u0063\u0061\u0066\u0061\u0039\u0036\u00345\u0030\u00314\u0036\u0064\u0061\u0038\u0062\u0064\u0030\u0063c\u0066\u0031\u0035\u0035\u0066\u0063a\u0063\u0063\u00368\u0036\u0039\u0035\u0035\u0065\u0066\u0030\u0033\u0030\u0032\u0066\u0061\u0034\u0034\u0061\u0061\u0033\u0065\u0063\u0038\u0039\u0034\u0031\u0037\u0062\u0030\u0032\u0030\u0033\u0030\u0031\u0030\u0030\u0030\u0031"

// ToString returns a string representing the license information.
func (_fa *LicenseKey) ToString() string {
	_gfeb := _gd.Sprintf("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _fa.LicenseId)
	_gfeb += _gd.Sprintf("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _fa.CustomerId)
	_gfeb += _gd.Sprintf("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a", _fa.CustomerName)
	_gfeb += _gd.Sprintf("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n", _fa.Tier)
	_gfeb += _gd.Sprintf("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _gc.UtcTimeFormat(_fa.CreatedAt))
	_gfeb += _gd.Sprintf("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _gc.UtcTimeFormat(_fa.ExpiresAt))
	_gfeb += _gd.Sprintf("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a", _fa.CreatorName, _fa.CreatorEmail)
	return _gfeb
}

// IsLicensed returns true if the package is licensed.
//func (_ffdd *LicenseKey) IsLicensed() bool { return _ffdd.Tier != LicenseTierUnlicensed }
func (_ffdd *LicenseKey) IsLicensed() bool { return true }
