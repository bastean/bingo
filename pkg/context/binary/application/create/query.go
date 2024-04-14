package create

type Command struct {
	Name           string
	ShouldRecorded bool
}

type Argument struct {
	Name           string
	ShouldRecorded bool
}

type Option struct {
	Name           string
	ShouldRecorded bool
}

type Flag struct {
	Name           string
	ShouldRecorded bool
}

type Record struct {
	*Command
	Arguments   []*Argument
	Options     []*Option
	Flags       []*Flag
	Subcommands []*Record
}

type Query struct {
	Name, Description string
	*Record
}

func NewQuery(name, description string) *Query {
	person := &Record{
		Command: &Command{Name: "person"},
		Flags:   []*Flag{{Name: "firstName"}, {Name: "lastName"}, {Name: "ssn"}, {Name: "email"}, {Name: "phone"}},
	}

	address := &Record{
		Command: &Command{Name: "address"},
		Flags:   []*Flag{{Name: "country"}, {Name: "state"}, {Name: "city"}, {Name: "street"}, {Name: "zip"}},
	}

	payment := &Record{
		Command: &Command{Name: "payment"},
		Flags:   []*Flag{{Name: "creditCardNumber"}, {Name: "creditCardCvv"}, {Name: "creditCardExp"}, {Name: "bitcoinAddress"}, {Name: "bitcoinPrivateKey"}},
	}

	internet := &Record{
		Command: &Command{Name: "internet"},
		Flags:   []*Flag{{Name: "url"}, {Name: "domainName"}, {Name: "domainSuffix"}, {Name: "macAddress"}, {Name: "userAgent"}},
	}

	language := &Record{
		Command: &Command{Name: "language"},
		Flags:   []*Flag{{Name: "programmingLanguage"}, {Name: "programmingLanguageBest"}},
	}

	animal := &Record{
		Command: &Command{Name: "animal"},
		Flags:   []*Flag{{Name: "petName"}, {Name: "cat"}, {Name: "dog"}, {Name: "bird"}},
	}

	binaryModel := &Record{
		Command: &Command{ShouldRecorded: true},
		Subcommands: []*Record{
			person,
			address,
			payment,
			internet,
			language,
			animal,
		},
	}

	return &Query{
		Name:        name,
		Description: description,
		Record:      binaryModel,
	}
}
