package query_builder

type IQueryBuilder interface {
	Parser(query string) ([]error, string)
}
