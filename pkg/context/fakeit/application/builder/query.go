package builder

type Query struct{}

func NewQuery() *Query {
	return new(Query)
}
