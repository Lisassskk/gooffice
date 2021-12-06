
package logger ;import (_c "fmt";_ea "io";_fd "os";_fg "path/filepath";_e "runtime";);

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _ea .Writer ;};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Warning logs warning message.
func (_dc ConsoleLogger )Warning (format string ,args ...interface{}){if _dc .LogLevel >=LogLevelWarning {_cb :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_dc .output (_fd .Stdout ,_cb ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};func (_cd WriterLogger )logToWriter (_egg _ea .Writer ,_gb string ,_ccg string ,_cge ...interface{}){_gga (_egg ,_gb ,_ccg ,_cge );};

// Debug logs debug message.
func (_df ConsoleLogger )Debug (format string ,args ...interface{}){if _df .LogLevel >=LogLevelDebug {_fb :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_df .output (_fd .Stdout ,_fb ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Debug logs debug message.
func (_aag WriterLogger )Debug (format string ,args ...interface{}){if _aag .LogLevel >=LogLevelDebug {_gg :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_aag .logToWriter (_aag .Output ,_gg ,format ,args ...);};};func (_aab ConsoleLogger )output (_ab _ea .Writer ,_g string ,_cbf string ,_dge ...interface{}){_gga (_ab ,_g ,_cbf ,_dge ...);};

// Info logs info message.
func (_cc ConsoleLogger )Info (format string ,args ...interface{}){if _cc .LogLevel >=LogLevelInfo {_dg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cc .output (_fd .Stdout ,_dg ,format ,args ...);};};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Error logs error message.
func (_aa ConsoleLogger )Error (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelError {_db :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_aa .output (_fd .Stdout ,_db ,format ,args ...);};};

// Notice logs notice message.
func (_fa ConsoleLogger )Notice (format string ,args ...interface{}){if _fa .LogLevel >=LogLevelNotice {_ffb :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fa .output (_fd .Stdout ,_ffb ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_eegc ConsoleLogger )IsLogLevel (level LogLevel )bool {return _eegc .LogLevel >=level };

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Trace logs trace message.
func (_fe ConsoleLogger )Trace (format string ,args ...interface{}){if _fe .LogLevel >=LogLevelTrace {_agc :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fe .output (_fd .Stdout ,_agc ,format ,args ...);};};

// DummyLogger does nothing.
type DummyLogger struct{};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_b string ,_ee ...interface{});Warning (_eeg string ,_d ...interface{});Notice (_bb string ,_a ...interface{});Info (_ae string ,_ag ...interface{});Debug (_ed string ,_bbf ...interface{});Trace (_ff string ,_eef ...interface{});IsLogLevel (_bd LogLevel )bool ;};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};var Log Logger =DummyLogger {};func _gga (_cbb _ea .Writer ,_fba string ,_egb string ,_cgee ...interface{}){_ ,_ccc ,_gad ,_agb :=_e .Caller (3);if !_agb {_ccc ="\u003f\u003f\u003f";_gad =0;}else {_ccc =_fg .Base (_ccc );};_ef :=_c .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_fba ,_ccc ,_gad )+_egb +"\u000a";_c .Fprintf (_cbb ,_ef ,_cgee ...);};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fab WriterLogger )IsLogLevel (level LogLevel )bool {return _fab .LogLevel >=level };

// Error logs error message.
func (_de WriterLogger )Error (format string ,args ...interface{}){if _de .LogLevel >=LogLevelError {_dga :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_de .logToWriter (_de .Output ,_dga ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Info logs info message.
func (_bg WriterLogger )Info (format string ,args ...interface{}){if _bg .LogLevel >=LogLevelInfo {_bbg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_bg .logToWriter (_bg .Output ,_bbg ,format ,args ...);};};

// Notice logs notice message.
func (_gc WriterLogger )Notice (format string ,args ...interface{}){if _gc .LogLevel >=LogLevelNotice {_eg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_gc .logToWriter (_gc .Output ,_eg ,format ,args ...);};};

// Trace logs trace message.
func (_cg WriterLogger )Trace (format string ,args ...interface{}){if _cg .LogLevel >=LogLevelTrace {_gf :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_cg .logToWriter (_cg .Output ,_gf ,format ,args ...);};};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _ea .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Warning logs warning message.
func (_ec WriterLogger )Warning (format string ,args ...interface{}){if _ec .LogLevel >=LogLevelWarning {_gaa :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ec .logToWriter (_ec .Output ,_gaa ,format ,args ...);};};
