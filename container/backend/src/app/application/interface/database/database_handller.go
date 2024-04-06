package database

type IDatabaseHandller interface {
	OpenDB(url string) error
	CloseDB() error
	Transaction(fn func(IRepositoryFactory) error) error
	GetRepositoryFactory() IRepositoryFactory
}
