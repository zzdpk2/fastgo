package options

type ServerOptions struct {
	MySQLOptions *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: NewMySQLOptions(),
	}
}
