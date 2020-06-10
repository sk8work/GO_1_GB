package contact

// Contact struct
type Contact struct {
	Name              string
	Mobile            int
	successfullPerson bool
}

func (p *Contact) SetsuccessfulPerson(successfullPerson bool) {
	p.successfullPerson = successfullPerson
}
