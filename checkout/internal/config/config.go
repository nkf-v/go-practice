package config

type Config struct {
	App            AppConfig      `yaml:"app"`
	ProductService ProductService `yaml:"product_service"`
	Loms           Loms           `yaml:"loms"`
	DataBase       DataBase       `yaml:"database"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type ProductService struct {
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
}

type Loms struct {
	Url string `yaml:"url"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
