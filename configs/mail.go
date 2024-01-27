package configs

type Mailer struct {
	Host        string `envconfig:"MAIL_HOST" required:"true"`
	Port        int    `envconfig:"MAIL_PORT" required:"true"`
	UserName    string `envconfig:"MAIL_USERNAME" required:"true"`
	Password    string `envconfig:"MAIL_PASSWORD" required:"true"`
	FromAddress string `envconfig:"MAIL_FROM_ADDRESS" required:"true"`
}
