
package license ;import _ga "gitee.com/gooffice/gooffice/internal/license";

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _ga .MakeUnlicensedKey ()};

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense =_ga .LegacyLicense ;

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {return _ga .SetLegacyLicenseKey (s )};

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _ga .GetLicenseKey ()};

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _ga .SetLicenseKey (content ,customerName );};

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _ga .SetMeteredKey (apiKey )};const (LicenseTierUnlicensed =_ga .LicenseTierUnlicensed ;LicenseTierCommunity =_ga .LicenseTierCommunity ;LicenseTierIndividual =_ga .LicenseTierIndividual ;LicenseTierBusiness =_ga .LicenseTierBusiness ;);

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_ga .MeteredStatus ,error ){return _ga .GetMeteredState ()};

// LicenseKey represents a loaded license key.
type LicenseKey =_ga .LicenseKey ;

// LegacyLicenseType is the type of license
type LegacyLicenseType =_ga .LegacyLicenseType ;
