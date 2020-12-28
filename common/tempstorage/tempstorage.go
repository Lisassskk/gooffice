
package tempstorage ;import _dd "io";type storage interface{Open (_f string )(File ,error );TempFile (_e ,_dg string )(File ,error );TempDir (_a string )(string ,error );RemoveAll (_fc string )error ;Add (_c string )error ;};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _cd .Open (path )};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _cd .TempDir (pattern )};

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_dd .Reader ;_dd .Writer ;_dd .Closer ;Name ()string ;};var _cd storage ;

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_cd =newStorage };

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _cd .RemoveAll (dir )};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _cd .Add (path )};

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _cd .TempFile (dir ,pattern )};