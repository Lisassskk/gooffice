
package memstore ;import (_c "encoding/hex";_e "errors";_gf "fmt";_eb "gitee.com/gooffice/gooffice/common/tempstorage";_g "io";_fb "io/ioutil";_d "math/rand";_fc "sync";);

// Name returns the filename of the underlying memDataCell
func (_b *memFile )Name ()string {return _b ._ga ._bd };

// TempDir creates a name for a new temp directory using a pattern argument
func (_gad *memStorage )TempDir (pattern string )(string ,error ){return _bg (pattern ),nil };

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_a *memFile )Read (p []byte )(int ,error ){_gg :=_a ._de ;_fg :=_a ._ga ._fbe ;_gfe :=int64 (len (p ));if _gfe > _fg {_gfe =_fg ;p =p [:_gfe ];};if _gg >=_fg {return 0,_g .EOF ;};_dc :=_gg +_gfe ;if _dc >=_fg {_dc =_fg ;};_dg :=copy (p ,_a ._ga ._cd [_gg :_dc ]);_a ._de =_dc ;return _dg ,nil ;};

// Open returns tempstorage File object by name
func (_ce *memStorage )Open (path string )(_eb .File ,error ){_ced ,_bf :=_ce ._fd .Load (path );if !_bf {return nil ,_e .New (_gf .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_ga :_ced .(*memDataCell )},nil ;};

// Close is not applicable in this implementation
func (_ebg *memFile )Close ()error {return nil };

// Add reads a file from a disk and adds it to the storage
func (_ab *memStorage )Add (path string )error {_ ,_gb :=_ab ._fd .Load (path );if _gb {return nil ;};_cf ,_aa :=_fb .ReadFile (path );if _aa !=nil {return _aa ;};_ab ._fd .Store (path ,&memDataCell {_bd :path ,_cd :_cf ,_fbe :int64 (len (_cf ))});return nil ;};

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_fe *memFile )Write (p []byte )(int ,error ){_fe ._ga ._cd =append (_fe ._ga ._cd ,p ...);_fe ._ga ._fbe +=int64 (len (p ));return len (p ),nil ;};type memDataCell struct{_bd string ;_cd []byte ;_fbe int64 ;};func _ee (_af int )(string ,error ){_ca :=make ([]byte ,_af );if _ ,_aad :=_d .Read (_ca );_aad !=nil {return "",_aad ;};return _c .EncodeToString (_ca ),nil ;};

// RemoveAll removes all files according to the dir argument prefix
func (_fbef *memStorage )RemoveAll (dir string )error {_fbef ._fd .Range (func (_bb ,_ec interface{})bool {_fbef ._fd .Delete (_bb );return true });return nil ;};type memFile struct{_ga *memDataCell ;_de int64 ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_ggg :=memStorage {_fd :_fc .Map {}};_eb .SetAsStorage (&_ggg )};type memStorage struct{_fd _fc .Map };

// TempFile creates a new empty file in the storage and returns it
func (_gfef *memStorage )TempFile (dir ,pattern string )(_eb .File ,error ){_be :=dir +"\u002f"+_bg (pattern );_bfc :=&memDataCell {_bd :_be ,_cd :[]byte {}};_da :=&memFile {_ga :_bfc };_gfef ._fd .Store (_be ,_bfc );return _da ,nil ;};func _bg (_ff string )string {_cb ,_ :=_ee (6);return _ff +_cb };
