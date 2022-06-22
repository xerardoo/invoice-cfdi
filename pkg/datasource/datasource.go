package datasource

type DataSource interface {
	Connect() error
}
