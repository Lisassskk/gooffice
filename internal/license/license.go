package license

import (
	_be "bytes"
	_bf "compress/gzip"
	_de "crypto"
	_fd "crypto/aes"
	_cd "crypto/cipher"
	_ad "crypto/rand"
	_ab "crypto/rsa"
	_bec "crypto/sha256"
	_bfc "crypto/sha512"
	_bg "crypto/x509"
	_bff "encoding/base64"
	_cf "encoding/binary"
	_eb "encoding/hex"
	_dea "encoding/json"
	_fa "encoding/pem"
	_db "errors"
	_eg "fmt"
	_egg "gitee.com/gooffice/gooffice/common"
	_ef "gitee.com/gooffice/gooffice/common/logger"
	_g "io"
	_ff "io/ioutil"
	_e "log"
	_f "net"
	_a "net/http"
	_c "os"
	_ga "path/filepath"
	_b "regexp"
	_bc "sort"
	_fg "strings"
	_gg "sync"
	_cc "time"
)

func (_gag *LicenseKey) IsLicensed() bool {
	//if _gag == nil {
	//	return false
	//}
	//return _gag.Tier != LicenseTierUnlicensed || _gag._bfe
	return true
}

func Track(docKey string, useKey string) error { return _bd(docKey, useKey, false) }

const _ee = "\u0055\u004e\u0049OF\u0046\u0049\u0043\u0045\u005f\u0043\u0055\u0053\u0054\u004f\u004d\u0045\u0052\u005f\u004e\u0041\u004d\u0045"

func _bd(_egaa string, _afab string, _aagg bool) error {
	if _bda == nil {
		return _db.New("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079")
	}
	if !_bda._bfe || len(_bda._abg) == 0 {
		return nil
	}
	if len(_egaa) == 0 && !_aagg {
		return _db.New("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	_dfb.Lock()
	defer _dfb.Unlock()
	if _gebc == nil {
		_gebc = map[string]struct{}{}
	}
	if _gfc == nil {
		_gfc = map[string]int{}
	}
	_cbae := 0
	if !_aagg {
		_, _feb := _gebc[_egaa]
		if !_feb {
			_gebc[_egaa] = struct{}{}
			_cbae++
		}
		if _cbae == 0 {
			return nil
		}
		_gfc[_afab]++
	}
	_fbd := _cc.Now()
	_fafb, _bedg := _afag.loadState(_bda._abg)
	if _bedg != nil {
		_ef.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _bedg)
		return _bedg
	}
	if _fafb.Usage == nil {
		_fafb.Usage = map[string]int{}
	}
	for _dbcf, _adb := range _gfc {
		_fafb.Usage[_dbcf] += _adb
	}
	_gfc = nil
	const _cgf = 24 * _cc.Hour
	const _gaa = 3 * 24 * _cc.Hour
	if len(_fafb.Instance) == 0 || _fbd.Sub(_fafb.LastReported) > _cgf || (_fafb.LimitDocs && _fafb.RemainingDocs <= _fafb.Docs+int64(_cbae)) || _aagg {
		_ceb, _gba := _c.Hostname()
		if _gba != nil {
			return _gba
		}
		_aecb := _fafb.Docs
		_gce, _baba, _gba := _dgbg()
		if _gba != nil {
			return _gba
		}
		_bc.Strings(_baba)
		_bc.Strings(_gce)
		_bbe, _gba := _fff()
		if _gba != nil {
			return _gba
		}
		_da := false
		for _, _eff := range _baba {
			if _eff == _bbe.String() {
				_da = true
			}
		}
		if !_da {
			_baba = append(_baba, _bbe.String())
		}
		_eag := _aga()
		_eag._dbg = _bda._abg
		_aecb += int64(_cbae)
		_gfce := meteredUsageCheckinForm{Instance: _fafb.Instance, Next: _fafb.Next, UsageNumber: int(_aecb), NumFailed: _fafb.NumErrors, Hostname: _ceb, LocalIP: _fg.Join(_baba, "\u002c\u0020"), MacAddress: _fg.Join(_gce, "\u002c\u0020"), Package: "\u0075n\u0069\u006f\u0066\u0066\u0069\u0063e", PackageVersion: _egg.Version, Usage: _fafb.Usage}
		if len(_gce) == 0 {
			_gfce.MacAddress = "\u006e\u006f\u006e\u0065"
		}
		_ddd := int64(0)
		_faff := _fafb.NumErrors
		_dcf := _fbd
		_fgf := 0
		_gff := _fafb.LimitDocs
		_bba, _gba := _eag.checkinUsage(_gfce)
		if _gba != nil {
			if _fbd.Sub(_fafb.LastReported) > _gaa {
				return _db.New("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e")
			}
			_ddd = _aecb
			_faff++
			_dcf = _fafb.LastReported
		} else {
			_gff = _bba.LimitDocs
			_fgf = _bba.RemainingDocs
			_faff = 0
		}
		if len(_bba.Instance) == 0 {
			_bba.Instance = _gfce.Instance
		}
		if len(_bba.Next) == 0 {
			_bba.Next = _gfce.Next
		}
		_gba = _afag.updateState(_eag._dbg, _bba.Instance, _bba.Next, int(_ddd), _gff, _fgf, int(_faff), _dcf, nil)
		if _gba != nil {
			return _gba
		}
		if !_bba.Success {
			return _eg.Errorf("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s", _bba.Message)
		}
	} else {
		_bedg = _afag.updateState(_bda._abg, _fafb.Instance, _fafb.Next, int(_fafb.Docs)+_cbae, _fafb.LimitDocs, int(_fafb.RemainingDocs), int(_fafb.NumErrors), _fafb.LastReported, _fafb.Usage)
		if _bedg != nil {
			return _bedg
		}
	}
	return nil
}

type stateLoader interface {
	loadState(_faf string) (reportState, error)
	updateState(_dbc, _bbb, _cec string, _fed int, _gagf bool, _fdg int, _bbbg int, _bgb _cc.Time, _cdf map[string]int) error
}

func (_geg *LicenseKey) getExpiryDateToCompare() _cc.Time {
	if _geg.Trial {
		return _cc.Now().UTC()
	}
	return _egg.ReleasedAt
}

type meteredStatusForm struct{}

func _fdf(_gegbe, _edf []byte) ([]byte, error) {
	_ddda := make([]byte, _bff.URLEncoding.DecodedLen(len(_edf)))
	_ffd, _dbeg := _bff.URLEncoding.Decode(_ddda, _edf)
	if _dbeg != nil {
		return nil, _dbeg
	}
	_ddda = _ddda[:_ffd]
	_babb, _dbeg := _fd.NewCipher(_gegbe)
	if _dbeg != nil {
		return nil, _dbeg
	}
	if len(_ddda) < _fd.BlockSize {
		return nil, _db.New("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074")
	}
	_agce := _ddda[:_fd.BlockSize]
	_ddda = _ddda[_fd.BlockSize:]
	_dfff := _cd.NewCFBDecrypter(_babb, _agce)
	_dfff.XORKeyStream(_ddda, _ddda)
	return _ddda, nil
}

var _bda = MakeUnlicensedKey()

var _fac = _cc.Date(2019, 6, 6, 0, 0, 0, 0, _cc.UTC)

type reportState struct {
	Instance      string         `json:"inst"`
	Next          string         `json:"n"`
	Docs          int64          `json:"d"`
	NumErrors     int64          `json:"e"`
	LimitDocs     bool           `json:"ld"`
	RemainingDocs int64          `json:"rd"`
	LastReported  _cc.Time       `json:"lr"`
	LastWritten   _cc.Time       `json:"lw"`
	Usage         map[string]int `json:"u"`
}

var _gga *_ab.PublicKey

const (
	_ec  = "\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
	_cfd = "\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d"
)

func (_fbf *meteredClient) getStatus() (meteredStatusResp, error) {
	var _bef meteredStatusResp
	_deee := _fbf._dga + "\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073"
	var _bgf meteredStatusForm
	_dgb, _ea := _dea.Marshal(_bgf)
	if _ea != nil {
		return _bef, _ea
	}
	_bfgg, _ea := _cgg(_dgb)
	if _ea != nil {
		return _bef, _ea
	}
	_eab, _ea := _a.NewRequest("\u0050\u004f\u0053\u0054", _deee, _bfgg)
	if _ea != nil {
		return _bef, _ea
	}
	_eab.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_eab.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_eab.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_eab.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _fbf._dbg)
	_dba, _ea := _fbf._fgc.Do(_eab)
	if _ea != nil {
		return _bef, _ea
	}
	defer _dba.Body.Close()
	if _dba.StatusCode != 200 {
		return _bef, _eg.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _dba.StatusCode)
	}
	_gcf, _ea := _fce(_dba)
	if _ea != nil {
		return _bef, _ea
	}
	_ea = _dea.Unmarshal(_gcf, &_bef)
	if _ea != nil {
		return _bef, _ea
	}
	return _bef, nil
}

func (_bbd *LicenseKey) TypeToString() string {
	if _bbd._bfe {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	if _bbd.Tier == LicenseTierUnlicensed {
		return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	}
	if _bbd.Tier == LicenseTierCommunity {
		return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065"
	}
	if _bbd.Tier == LicenseTierIndividual || _bbd.Tier == "\u0069\u006e\u0064i\u0065" {
		return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	}
	return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073"
}

type meteredClient struct {
	_dga string
	_dbg string
	_fgc *_a.Client
}

func GenRefId(prefix string) (string, error) {
	var _cbge _be.Buffer
	_cbge.WriteString(prefix)
	_gad := make([]byte, 8+16)
	_acee := _cc.Now().UTC().UnixNano()
	_cf.BigEndian.PutUint64(_gad, uint64(_acee))
	_, _gagd := _ad.Read(_gad[8:])
	if _gagd != nil {
		return "", _gagd
	}
	_cbge.WriteString(_eb.EncodeToString(_gad))
	return _cbge.String(), nil
}

func (_acf *LicenseKey) isExpired() bool {
	return _acf.getExpiryDateToCompare().After(_acf.ExpiresAt)
}

type LegacyLicense struct {
	Name        string
	Signature   string `json:",omitempty"`
	Expiration  _cc.Time
	LicenseType LegacyLicenseType
}

func _ag(_bb string, _bfff string) ([]byte, error) {
	var (
		_adc int
		_ada string
	)
	for _, _ada = range []string{"\u000a\u002b\u000a", "\u000d\u000a\u002b\r\u000a", "\u0020\u002b\u0020"} {
		if _adc = _fg.Index(_bfff, _ada); _adc != -1 {
			break
		}
	}
	if _adc == -1 {
		return nil, _eg.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072")
	}
	_def := _bfff[:_adc]
	_fab := _adc + len(_ada)
	_ggc := _bfff[_fab:]
	if _def == "" || _ggc == "" {
		return nil, _eg.Errorf("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065")
	}
	_ca, _ece := _bff.StdEncoding.DecodeString(_def)
	if _ece != nil {
		return nil, _eg.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c")
	}
	_bfa, _ece := _bff.StdEncoding.DecodeString(_ggc)
	if _ece != nil {
		return nil, _eg.Errorf("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065")
	}
	_cdg, _ := _fa.Decode([]byte(_bb))
	if _cdg == nil {
		return nil, _eg.Errorf("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_afa, _ece := _bg.ParsePKIXPublicKey(_cdg.Bytes)
	if _ece != nil {
		return nil, _ece
	}
	_ac := _afa.(*_ab.PublicKey)
	if _ac == nil {
		return nil, _eg.Errorf("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_fc := _bfc.New()
	_fc.Write(_ca)
	_bcb := _fc.Sum(nil)
	_ece = _ab.VerifyPKCS1v15(_ac, _de.SHA512, _bcb, _bfa)
	if _ece != nil {
		return nil, _ece
	}
	return _ca, nil
}

func init() {
	_afd, _agd := _eb.DecodeString(_cb)
	if _agd != nil {
		_e.Fatalf("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073", _agd)
	}
	_dgdd, _agd := _bg.ParsePKIXPublicKey(_afd)
	if _agd != nil {
		_e.Fatalf("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073", _agd)
	}
	_gga = _dgdd.(*_ab.PublicKey)
}

var _gfc map[string]int

func (_ggg defaultStateHolder) updateState(_baf, _dgddd, _caaf string, _bedf int, _dfe bool, _fcd int, _gd int, _dead _cc.Time, _cece map[string]int) error {
	_dfee := _agc()
	if len(_dfee) == 0 {
		return _db.New("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_ffc := _ga.Join(_dfee, "\u002eu\u006e\u0069\u0064\u006f\u0063")
	_geb := _c.MkdirAll(_ffc, 0777)
	if _geb != nil {
		return _geb
	}
	if len(_baf) < 20 {
		return _db.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_afg := []byte(_baf)
	_cbg := _bfc.Sum512_256(_afg[:20])
	_ecgf := _eb.EncodeToString(_cbg[:])
	_cba := _ga.Join(_ffc, _ecgf)
	var _aac reportState
	_aac.Docs = int64(_bedf)
	_aac.NumErrors = int64(_gd)
	_aac.LimitDocs = _dfe
	_aac.RemainingDocs = int64(_fcd)
	_aac.LastWritten = _cc.Now().UTC()
	_aac.LastReported = _dead
	_aac.Instance = _dgddd
	_aac.Next = _caaf
	_aac.Usage = _cece
	_bedb, _geb := _dea.Marshal(_aac)
	if _geb != nil {
		return _geb
	}
	const _aeg = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_bedb, _geb = _aca([]byte(_aeg), _bedb)
	if _geb != nil {
		return _geb
	}
	_geb = _ff.WriteFile(_cba, _bedb, 0600)
	if _geb != nil {
		return _geb
	}
	return nil
}

const (
	LicenseTierUnlicensed = "\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	LicenseTierCommunity  = "\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y"
	LicenseTierIndividual = "\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c"
	LicenseTierBusiness   = "\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073"
)

func _agc() string {
	_febb := _c.Getenv("\u0048\u004f\u004d\u0045")
	if len(_febb) == 0 {
		_febb, _ = _c.UserHomeDir()
	}
	return _febb
}

func _dgbg() ([]string, []string, error) {
	_fge, _ccc := _f.Interfaces()
	if _ccc != nil {
		return nil, nil, _ccc
	}
	var _acbe []string
	var _bbc []string
	for _, _dbee := range _fge {
		if _dbee.Flags&_f.FlagUp == 0 || _be.Equal(_dbee.HardwareAddr, nil) {
			continue
		}
		_cda, _cee := _dbee.Addrs()
		if _cee != nil {
			return nil, nil, _cee
		}
		_caea := 0
		for _, _dcd := range _cda {
			var _acff _f.IP
			switch _dbd := _dcd.(type) {
			case *_f.IPNet:
				_acff = _dbd.IP
			case *_f.IPAddr:
				_acff = _dbd.IP
			}
			if _acff.IsLoopback() {
				continue
			}
			if _acff.To4() == nil {
				continue
			}
			_bbc = append(_bbc, _acff.String())
			_caea++
		}
		_aaa := _dbee.HardwareAddr.String()
		if _aaa != "" && _caea > 0 {
			_acbe = append(_acbe, _aaa)
		}
	}
	return _acbe, _bbc, nil
}

const _aff = "\u0055\u004e\u0049\u004fFF\u0049\u0043\u0045\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041T\u0048"

func _cfg(_gf string, _cae string, _dg string) (string, error) {
	_dca := _fg.Index(_dg, _gf)
	if _dca == -1 {
		return "", _eg.Errorf("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_dff := _fg.Index(_dg, _cae)
	if _dff == -1 {
		return "", _eg.Errorf("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_aa := _dca + len(_gf) + 1
	return _dg[_aa : _dff-1], nil
}

func _aga() *meteredClient {
	_aag := meteredClient{_dga: "h\u0074\u0074\u0070\u0073\u003a\u002f/\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069d\u006f\u0063\u002ei\u006f/\u0061\u0070\u0069", _fgc: &_a.Client{Timeout: 30 * _cc.Second}}
	if _abb := _c.Getenv("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L"); _fg.HasPrefix(_abb, "\u0068\u0074\u0074\u0070") {
		_aag._dga = _abb
	}
	return &_aag
}

type LicenseKey struct {
	LicenseId    string   `json:"license_id"`
	CustomerId   string   `json:"customer_id"`
	CustomerName string   `json:"customer_name"`
	Tier         string   `json:"tier"`
	CreatedAt    _cc.Time `json:"-"`
	CreatedAtInt int64    `json:"created_at"`
	ExpiresAt    _cc.Time `json:"-"`
	ExpiresAtInt int64    `json:"expires_at"`
	CreatedBy    string   `json:"created_by"`
	CreatorName  string   `json:"creator_name"`
	CreatorEmail string   `json:"creator_email"`
	UniPDF       bool     `json:"unipdf"`
	UniOffice    bool     `json:"unioffice"`
	UniHTML      bool     `json:"unihtml"`
	Trial        bool     `json:"trial"`
	_bfe         bool
	_abg         string
}

func (_beb LegacyLicense) Verify(pubKey *_ab.PublicKey) error {
	_bbf := _beb
	_bbf.Signature = ""
	_ce := _be.Buffer{}
	_afdb := _dea.NewEncoder(&_ce)
	if _cea := _afdb.Encode(_bbf); _cea != nil {
		return _cea
	}
	_ccg, _fag := _eb.DecodeString(_beb.Signature)
	if _fag != nil {
		return _fag
	}
	_fb := _bec.Sum256(_ce.Bytes())
	_fag = _ab.VerifyPKCS1v15(pubKey, _de.SHA256, _fb[:], _ccg)
	return _fag
}

func (_fgd defaultStateHolder) loadState(_gca string) (reportState, error) {
	_dbb := _agc()
	if len(_dbb) == 0 {
		return reportState{}, _db.New("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_ccf := _ga.Join(_dbb, "\u002eu\u006e\u0069\u0064\u006f\u0063")
	_dcaf := _c.MkdirAll(_ccf, 0777)
	if _dcaf != nil {
		return reportState{}, _dcaf
	}
	if len(_gca) < 20 {
		return reportState{}, _db.New("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079")
	}
	_bac := []byte(_gca)
	_gb := _bfc.Sum512_256(_bac[:20])
	_cbf := _eb.EncodeToString(_gb[:])
	_fcc := _ga.Join(_ccf, _cbf)
	_gegb, _dcaf := _ff.ReadFile(_fcc)
	if _dcaf != nil {
		if _c.IsNotExist(_dcaf) {
			return reportState{}, nil
		}
		_ef.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dcaf)
		return reportState{}, _db.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	const _dec = "\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057"
	_gegb, _dcaf = _fdf([]byte(_dec), _gegb)
	if _dcaf != nil {
		return reportState{}, _dcaf
	}
	var _gea reportState
	_dcaf = _dea.Unmarshal(_gegb, &_gea)
	if _dcaf != nil {
		_ef.Log.Error("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076", _dcaf)
		return reportState{}, _db.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061")
	}
	return _gea, nil
}

func SetMeteredKey(apiKey string) error {
	if len(apiKey) == 0 {
		_ef.Log.Error("\u004d\u0065\u0074\u0065\u0072e\u0064\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0041\u0050\u0049 \u004b\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079")
		_ef.Log.Error("\u002d\u0020\u0047\u0072\u0061\u0062\u0020\u006f\u006e\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0072\u0065\u0065\u0020\u0054\u0069\u0065\u0072\u0020\u0061t\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063\u006c\u006fud\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _eg.Errorf("\u006de\u0074\u0065\u0072e\u0064\u0020\u006ci\u0063en\u0073\u0065\u0020\u0061\u0070\u0069\u0020k\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u003a\u0020\u0063\u0072\u0065\u0061\u0074\u0065 o\u006ee\u0020\u0061\u0074\u0020\u0068\u0074t\u0070\u0073\u003a\u002f\u002fc\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063.\u0069\u006f")
	}
	if _bda != nil && (_bda._bfe || _bda.Tier != LicenseTierUnlicensed) {
		_ef.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u0061\u006e\u006eo\u0074 \u0073\u0065\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0074\u0077\u0069c\u0065\u0020\u002d\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u006a\u0075\u0073\u0074\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069z\u0065\u0020\u006f\u006e\u0063\u0065")
		return _db.New("\u006c\u0069\u0063en\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0061\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0073\u0065\u0074")
	}
	_caa := _aga()
	_caa._dbg = apiKey
	_egd, _agf := _caa.getStatus()
	if _agf != nil {
		return _agf
	}
	if !_egd.Valid {
		return _db.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_bfgc := &LicenseKey{_bfe: true, _abg: apiKey}
	_bda = _bfgc
	return nil
}

type meteredUsageCheckinForm struct {
	Instance       string         `json:"inst"`
	Next           string         `json:"next"`
	UsageNumber    int            `json:"usage_number"`
	NumFailed      int64          `json:"num_failed"`
	Hostname       string         `json:"hostname"`
	LocalIP        string         `json:"local_ip"`
	MacAddress     string         `json:"mac_address"`
	Package        string         `json:"package"`
	PackageVersion string         `json:"package_version"`
	Usage          map[string]int `json:"u"`
}

func _eca(_ggf string) (LicenseKey, error) {
	var _bcf LicenseKey
	_aed, _dgd := _cfg(_ec, _cfd, _ggf)
	if _dgd != nil {
		return _bcf, _dgd
	}
	_bab, _dgd := _ag(_agaa, _aed)
	if _dgd != nil {
		return _bcf, _dgd
	}
	_dgd = _dea.Unmarshal(_bab, &_bcf)
	if _dgd != nil {
		return _bcf, _dgd
	}
	_bcf.CreatedAt = _cc.Unix(_bcf.CreatedAtInt, 0)
	if _bcf.ExpiresAtInt > 0 {
		_ded := _cc.Unix(_bcf.ExpiresAtInt, 0)
		_bcf.ExpiresAt = _ded
	}
	return _bcf, nil
}

func (_dd *meteredClient) checkinUsage(_bag meteredUsageCheckinForm) (meteredUsageCheckinResp, error) {
	_bag.Package = "\u0075n\u0069\u006f\u0066\u0066\u0069\u0063e"
	_bag.PackageVersion = _egg.Version
	var _aec meteredUsageCheckinResp
	_eabg := _dd._dga + "\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e"
	_fbc, _cfe := _dea.Marshal(_bag)
	if _cfe != nil {
		return _aec, _cfe
	}
	_gfa, _cfe := _cgg(_fbc)
	if _cfe != nil {
		return _aec, _cfe
	}
	_dfd, _cfe := _a.NewRequest("\u0050\u004f\u0053\u0054", _eabg, _gfa)
	if _cfe != nil {
		return _aec, _cfe
	}
	_dfd.Header.Add("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e")
	_dfd.Header.Add("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_dfd.Header.Add("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0067\u007a\u0069\u0070")
	_dfd.Header.Add("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y", _dd._dbg)
	_fe, _cfe := _dd._fgc.Do(_dfd)
	if _cfe != nil {
		return _aec, _cfe
	}
	defer _fe.Body.Close()
	if _fe.StatusCode != 200 {
		return _aec, _eg.Errorf("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064", _fe.StatusCode)
	}
	_eabb, _cfe := _fce(_fe)
	if _cfe != nil {
		return _aec, _cfe
	}
	_cfe = _dea.Unmarshal(_eabb, &_aec)
	if _cfe != nil {
		return _aec, _cfe
	}
	return _aec, nil
}

func SetLegacyLicenseKey(s string) error {
	_bcef := _b.MustCompile("\u005c\u0073")
	s = _bcef.ReplaceAllString(s, "")
	var _bcc _g.Reader
	_bcc = _fg.NewReader(s)
	_bcc = _bff.NewDecoder(_bff.RawURLEncoding, _bcc)
	_bcc, _gab := _bf.NewReader(_bcc)
	if _gab != nil {
		return _gab
	}
	_eabd := _dea.NewDecoder(_bcc)
	_gbg := &LegacyLicense{}
	if _efd := _eabd.Decode(_gbg); _efd != nil {
		return _efd
	}
	if _ecd := _gbg.Verify(_gga); _ecd != nil {
		return _db.New("\u006c\u0069\u0063en\u0073\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006e\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gbg.Expiration.Before(_egg.ReleasedAt) {
		return _db.New("\u006ci\u0063e\u006e\u0073\u0065\u0020\u0065\u0078\u0070\u0069\u0072\u0065\u0064")
	}
	_bdd := _cc.Now().UTC()
	_gfg := LicenseKey{}
	_gfg.CreatedAt = _bdd
	_gfg.CustomerId = "\u004c\u0065\u0067\u0061\u0063\u0079"
	_gfg.CustomerName = _gbg.Name
	_gfg.Tier = LicenseTierBusiness
	_gfg.ExpiresAt = _gbg.Expiration
	_gfg.CreatorName = "\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u0073\u0075p\u0070\u006f\u0072\u0074"
	_gfg.CreatorEmail = "\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0040\u0075\u006e\u0069\u0064o\u0063\u002e\u0069\u006f"
	_gfg.UniOffice = true
	_bda = &_gfg
	return nil
}

func (_bed *LicenseKey) ToString() string {
	if _bed._bfe {
		return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e"
	}
	_fde := _eg.Sprintf("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _bed.LicenseId)
	_fde += _eg.Sprintf("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a", _bed.CustomerId)
	_fde += _eg.Sprintf("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a", _bed.CustomerName)
	_fde += _eg.Sprintf("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n", _bed.Tier)
	_fde += _eg.Sprintf("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _egg.UtcTimeFormat(_bed.CreatedAt))
	if _bed.ExpiresAt.IsZero() {
		_fde += "\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a"
	} else {
		_fde += _eg.Sprintf("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a", _egg.UtcTimeFormat(_bed.ExpiresAt))
	}
	_fde += _eg.Sprintf("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a", _bed.CreatorName, _bed.CreatorEmail)
	return _fde
}

var _dfb = &_gg.Mutex{}

func SetLicenseKey(content string, customerName string) error {
	if _bgd {
		return nil
	}
	_feg, _ecb := _eca(content)
	if _ecb != nil {
		_ef.Log.Error("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076", _ecb)
		return _ecb
	}
	if !_fg.EqualFold(_feg.CustomerName, customerName) {
		_ef.Log.Error("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027", customerName, _feg.CustomerName)
		return _eg.Errorf("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'", customerName, _feg.CustomerName)
	}
	_ecb = _feg.Validate()
	if _ecb != nil {
		_ef.Log.Error("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076", _ecb)
		return _ecb
	}
	_bda = &_feg
	return nil
}

func _cgg(_fgff []byte) (_g.Reader, error) {
	_bad := new(_be.Buffer)
	_abd := _bf.NewWriter(_bad)
	_abd.Write(_fgff)
	_efce := _abd.Close()
	if _efce != nil {
		return nil, _efce
	}
	return _bad, nil
}

func _fff() (_f.IP, error) {
	_bbed, _becg := _f.Dial("\u0075\u0064\u0070", "\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030")
	if _becg != nil {
		return nil, _becg
	}
	defer _bbed.Close()
	_bce := _bbed.LocalAddr().(*_f.UDPAddr)
	return _bce.IP, nil
}

type defaultStateHolder struct{}

type LegacyLicenseType byte

var _gebc map[string]struct{}

var _fca = _cc.Date(2020, 1, 1, 0, 0, 0, 0, _cc.UTC)

var _ge = _cc.Date(2010, 1, 1, 0, 0, 0, 0, _cc.UTC)

func GetLicenseKey() *LicenseKey {
	if _bda == nil {
		return nil
	}
	_fbe := *_bda
	return &_fbe
}

func _bbbd(_adg *_a.Response) (_g.ReadCloser, error) {
	var _fdc error
	var _fec _g.ReadCloser
	switch _fg.ToLower(_adg.Header.Get("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")) {
	case "\u0067\u007a\u0069\u0070":
		_fec, _fdc = _bf.NewReader(_adg.Body)
		if _fdc != nil {
			return _fec, _fdc
		}
		defer _fec.Close()
	default:
		_fec = _adg.Body
	}
	return _fec, nil
}

func TrackUse(useKey string) {
	if _bda == nil {
		return
	}
	if !_bda._bfe || len(_bda._abg) == 0 {
		return
	}
	if len(useKey) == 0 {
		return
	}
	_dfb.Lock()
	defer _dfb.Unlock()
	if _gfc == nil {
		_gfc = map[string]int{}
	}
	_gfc[useKey]++
}

func GetMeteredState() (MeteredStatus, error) {
	if _bda == nil {
		return MeteredStatus{}, _db.New("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	if !_bda._bfe || len(_bda._abg) == 0 {
		return MeteredStatus{}, _db.New("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074")
	}
	_dgc, _eac := _afag.loadState(_bda._abg)
	if _eac != nil {
		_ef.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eac)
		return MeteredStatus{}, _eac
	}
	if _dgc.Docs > 0 {
		_fga := _bd("", "", true)
		if _fga != nil {
			return MeteredStatus{}, _fga
		}
	}
	_dfb.Lock()
	defer _dfb.Unlock()
	_adcc := _aga()
	_adcc._dbg = _bda._abg
	_cac, _eac := _adcc.getStatus()
	if _eac != nil {
		return MeteredStatus{}, _eac
	}
	if !_cac.Valid {
		return MeteredStatus{}, _db.New("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064")
	}
	_afe := MeteredStatus{OK: true, Credits: _cac.OrgCredits, Used: _cac.OrgUsed}
	return _afe, nil
}

func (_gc *LicenseKey) Validate() error {
	if _gc._bfe {
		return nil
	}
	if len(_gc.LicenseId) < 10 {
		return _eg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064")
	}
	if len(_gc.CustomerId) < 10 {
		return _eg.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064")
	}
	if len(_gc.CustomerName) < 1 {
		return _eg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065")
	}
	if _ge.After(_gc.CreatedAt) {
		return _eg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if _gc.ExpiresAt.IsZero() {
		_acb := _gc.CreatedAt.AddDate(1, 0, 0)
		if _fca.After(_acb) {
			_acb = _fca
		}
		_gc.ExpiresAt = _acb
	}
	if _gc.CreatedAt.After(_gc.ExpiresAt) {
		return _eg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074")
	}
	if _gc.isExpired() {
		return _eg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064")
	}
	if len(_gc.CreatorName) < 1 {
		return _eg.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065")
	}
	if len(_gc.CreatorEmail) < 1 {
		return _eg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c")
	}
	if _gc.CreatedAt.After(_fac) {
		if !_gc.UniOffice {
			return _eg.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u003a\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u006b\u0065\u0079\u0020i\u0073\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		}
	}
	return nil
}

const _cb = "\u0033\u0030\u0035\u0063\u0033\u0030\u0030\u00640\u0036\u0030\u0039\u0032\u0061\u0038\u00364\u0038\u0038\u0036\u0066\u0037\u0030d\u0030\u0031\u0030\u0031\u0030\u00310\u0035\u0030\u0030\u0030\u0033\u0034\u0062\u0030\u0030\u0033\u0030\u00348\u0030\u0032\u0034\u0031\u0030\u0030\u0062\u0038\u0037\u0065\u0061\u0066\u0062\u0036\u0063\u0030\u0037\u0034\u0039\u0039\u0065\u0062\u00397\u0063\u0063\u0039\u0064\u0033\u0035\u0036\u0035\u0065\u0063\u00663\u0031\u0036\u0038\u0031\u0039\u0036\u0033\u0030\u0031\u0039\u0030\u0037c\u0038\u0034\u0031\u0061\u0064\u0064c6\u0036\u0035\u0030\u0038\u0036\u0062\u0062\u0033\u0065\u0064\u0038\u0065\u0062\u0031\u0032\u0064\u0039\u0064\u0061\u0032\u0036\u0063\u0061\u0066\u0061\u0039\u0036\u00345\u0030\u00314\u0036\u0064\u0061\u0038\u0062\u0064\u0030\u0063c\u0066\u0031\u0035\u0035\u0066\u0063a\u0063\u0063\u00368\u0036\u0039\u0035\u0035\u0065\u0066\u0030\u0033\u0030\u0032\u0066\u0061\u0034\u0034\u0061\u0061\u0033\u0065\u0063\u0038\u0039\u0034\u0031\u0037\u0062\u0030\u0032\u0030\u0033\u0030\u0031\u0030\u0030\u0030\u0031"

const _agaa = "\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n"

func init() {
	_eef := _c.Getenv(_aff)
	_fba := _c.Getenv(_ee)
	if len(_eef) == 0 || len(_fba) == 0 {
		return
	}
	_defa, _cfc := _ff.ReadFile(_eef)
	if _cfc != nil {
		_ef.Log.Error("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076", _cfc)
		return
	}
	_cfc = SetLicenseKey(string(_defa), _fba)
	if _cfc != nil {
		_ef.Log.Error("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076", _cfc)
		return
	}
}

var _afag stateLoader = defaultStateHolder{}

func MakeUnlicensedKey() *LicenseKey {
	_ecg := LicenseKey{}
	_ecg.CustomerName = "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064"
	_ecg.Tier = LicenseTierUnlicensed
	_ecg.CreatedAt = _cc.Now().UTC()
	_ecg.CreatedAtInt = _ecg.CreatedAt.Unix()
	return &_ecg
}

type meteredUsageCheckinResp struct {
	Instance      string `json:"inst"`
	Next          string `json:"next"`
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	RemainingDocs int    `json:"rd"`
	LimitDocs     bool   `json:"ld"`
}

func _df(_dee string, _ba []byte) (string, error) {
	_bfg, _ := _fa.Decode([]byte(_dee))
	if _bfg == nil {
		return "", _eg.Errorf("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	_dc, _dbf := _bg.ParsePKCS1PrivateKey(_bfg.Bytes)
	if _dbf != nil {
		return "", _dbf
	}
	_af := _bfc.New()
	_af.Write(_ba)
	_ega := _af.Sum(nil)
	_dcb, _dbf := _ab.SignPKCS1v15(_ad.Reader, _dc, _de.SHA512, _ega)
	if _dbf != nil {
		return "", _dbf
	}
	_ae := _bff.StdEncoding.EncodeToString(_ba)
	_ae += "\u000a\u002b\u000a"
	_ae += _bff.StdEncoding.EncodeToString(_dcb)
	return _ae, nil
}

var _bgd = false

func _aca(_dgbd, _gcc []byte) ([]byte, error) {
	_aaf, _edd := _fd.NewCipher(_dgbd)
	if _edd != nil {
		return nil, _edd
	}
	_aafd := make([]byte, _fd.BlockSize+len(_gcc))
	_dbgc := _aafd[:_fd.BlockSize]
	if _, _acae := _g.ReadFull(_ad.Reader, _dbgc); _acae != nil {
		return nil, _acae
	}
	_bcbe := _cd.NewCFBEncrypter(_aaf, _dbgc)
	_bcbe.XORKeyStream(_aafd[_fd.BlockSize:], _gcc)
	_cad := make([]byte, _bff.URLEncoding.EncodedLen(len(_aafd)))
	_bff.URLEncoding.Encode(_cad, _aafd)
	return _cad, nil
}

type meteredStatusResp struct {
	Valid        bool  `json:"valid"`
	OrgCredits   int64 `json:"org_credits"`
	OrgUsed      int64 `json:"org_used"`
	OrgRemaining int64 `json:"org_remaining"`
}

func _fce(_ccb *_a.Response) ([]byte, error) {
	var _gbe []byte
	_eda, _efg := _bbbd(_ccb)
	if _efg != nil {
		return _gbe, _efg
	}
	return _ff.ReadAll(_eda)
}

type MeteredStatus struct {
	OK      bool
	Credits int64
	Used    int64
}
