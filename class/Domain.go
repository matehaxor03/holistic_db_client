package class

type Domain struct {
	domain_name *string
}

func NewDomain(domain_name *string) (*Domain) {
	x := Domain{domain_name: domain_name}
	return &x
}

func (this *Domain) GetDomainName() *string {
	return (*this).domain_name
 }
