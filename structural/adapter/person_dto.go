package adapter

type PersonDTO struct {
	FullName string
}

func ToDto(p Person) PersonDTO {
	return PersonDTO{
		FullName: p.Name + " " + p.LastName,
	}
}
