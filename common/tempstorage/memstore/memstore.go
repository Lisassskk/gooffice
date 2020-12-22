// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore

import (
	_e "encoding/hex"
	_df "errors"
	_da "fmt"
	_c "io"
	_d "io/ioutil"
	_ad "math/rand"
	_ae "sync"

	_ea "gitee.com/yu_sheng/gooffice/common/tempstorage"
)

func _eg(_dc int) (string, error) {
	_bd := make([]byte, _dc)
	if _, _cc := _ad.Read(_bd); _cc != nil {
		return "", _cc
	}
	return _e.EncodeToString(_bd), nil
}

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_f *memFile) Read(p []byte) (int, error) {
	_ab := _f._ee
	_g := _f._ag._bf
	_ce := int64(len(p))
	if _ce > _g {
		_ce = _g
		p = p[:_ce]
	}
	if _ab >= _g {
		return 0, _c.EOF
	}
	_ff := _ab + _ce
	if _ff >= _g {
		_ff = _g
	}
	_fa := copy(p, _f._ag._de[_ab:_ff])
	_f._ee = _ff
	return _fa, nil
}

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage()          { _agb := memStorage{_b: _ae.Map{}}; _ea.SetAsStorage(&_agb) }
func _cge(_ge string) string { _dec, _ := _eg(6); return _ge + _dec }

type memStorage struct{ _b _ae.Map }
type memDataCell struct {
	_bc string
	_de []byte
	_bf int64
}

// TempFile creates a new empty file in the storage and returns it
func (_dd *memStorage) TempFile(dir, pattern string) (_ea.File, error) {
	_eb := dir + "\u002f" + _cge(pattern)
	_faa := &memDataCell{_bc: _eb, _de: []byte{}}
	_ga := &memFile{_ag: _faa}
	_dd._b.Store(_eb, _faa)
	return _ga, nil
}

// TempDir creates a name for a new temp directory using a pattern argument
func (_ca *memStorage) TempDir(pattern string) (string, error) { return _cge(pattern), nil }

type memFile struct {
	_ag *memDataCell
	_ee int64
}

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_gb *memFile) Write(p []byte) (int, error) {
	_gb._ag._de = append(_gb._ag._de, p...)
	_gb._ag._bf += int64(len(p))
	return len(p), nil
}

// Close is not applicable in this implementation
func (_fe *memFile) Close() error { return nil }

// Open returns tempstorage File object by name
func (_ac *memStorage) Open(path string) (_ea.File, error) {
	_fef, _ec := _ac._b.Load(path)
	if !_ec {
		return nil, _df.New(_da.Sprintf("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", path))
	}
	return &memFile{_ag: _fef.(*memDataCell)}, nil
}

// RemoveAll removes all files according to the dir argument prefix
func (_agg *memStorage) RemoveAll(dir string) error {
	_agg._b.Range(func(_gf, _db interface{}) bool { _agg._b.Delete(_gf); return true })
	return nil
}

// Name returns the filename of the underlying memDataCell
func (_cd *memFile) Name() string { return _cd._ag._bc }

// Add reads a file from a disk and adds it to the storage
func (_cg *memStorage) Add(path string) error {
	_, _agf := _cg._b.Load(path)
	if _agf {
		return nil
	}
	_bg, _dag := _d.ReadFile(path)
	if _dag != nil {
		return _dag
	}
	_cg._b.Store(path, &memDataCell{_bc: path, _de: _bg, _bf: int64(len(_bg))})
	return nil
}
