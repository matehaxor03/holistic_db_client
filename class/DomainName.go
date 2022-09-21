package class

type DomainName struct {
	domain_name *string
}

func NewDomainName(domain_name *string) (*DomainName) {
	x := DomainName{domain_name: domain_name}
	return &x
}

func (this *DomainName) GetDomainName() *string {
	return (*this).domain_name
 }
