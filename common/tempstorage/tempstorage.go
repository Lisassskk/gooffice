
package tempstorage ;import _g "io";

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _ee .TempDir (pattern )};

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _ee .RemoveAll (dir )};

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_g .Reader ;_g .Writer ;_g .Closer ;Name ()string ;};

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _ee .TempFile (dir ,pattern )};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _ee .Add (path )};var _ee storage ;

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_ee =newStorage };type storage interface{Open (_e string )(File ,error );TempFile (_ga ,_b string )(File ,error );TempDir (_ea string )(string ,error );RemoveAll (_c string )error ;Add (_d string )error ;};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _ee .Open (path )};
