
package diskstore ;import (_gag "gitee.com/gooffice/gooffice/common/tempstorage";_a "io/ioutil";_ga "os";_g "strings";);

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_e :=diskStorage {};_gag .SetAsStorage (&_e )};

// Open opens file from disk according to a path
func (_gf diskStorage )Open (path string )(_gag .File ,error ){return _ga .Open (path )};

// RemoveAll removes all files in the directory
func (_d diskStorage )RemoveAll (dir string )error {if _g .HasPrefix (dir ,_ga .TempDir ()){return _ga .RemoveAll (dir );};return nil ;};

// Add is not applicable in the diskstore implementation
func (_gd diskStorage )Add (path string )error {return nil };type diskStorage struct{};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_b diskStorage )TempDir (pattern string )(string ,error ){return _a .TempDir ("",pattern )};

// TempFile creates a new temp file by calling ioutil TempFile
func (_c diskStorage )TempFile (dir ,pattern string )(_gag .File ,error ){return _a .TempFile (dir ,pattern );};
