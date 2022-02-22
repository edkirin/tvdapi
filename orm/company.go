package orm

type Company struct {
	ID       int
	Caption  string
	Address  string
	City     string
	Email    string `json:"-"`
	AccessID int
	Access   CompanyAccess
}

type CompanyAccess struct {
	ID                  int
	FiscalConfiguration bool
	Fiscalization       bool
	FiscalModuleB       bool
	FiscalModuleC       bool
	FiscalModuleD       bool
	FiscalModuleE       bool
	FiscalModuleF       bool
}

func (Company) TableName() string {
	return "vending_companies"
}

func (CompanyAccess) TableName() string {
	return "administration_companyaccess"
}
