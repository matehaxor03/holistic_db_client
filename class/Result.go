package class

type Result struct {
	errors []error
	data *Map
}

func NewResult() (*Result) {
	errors := []error{}
	data := Map{}
	x := Result{errors: errors, data: &data}
	return &x
}

func (this *Result) LogError(err error) (*Result)  {
	errors := (*this).GetErrors()
	errors = append(errors, err)
	return this
}

func (this *Result) GetErrors() ([]error)  {
	return (*this).errors
}

func (this *Result) GetData() (*Map)  {
	return (*this).data
}

