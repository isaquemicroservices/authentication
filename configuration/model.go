package configuration

// Configuration main configuration struct
type Configuration struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Address     string   `json:"address"`
	SecretKey   string   `json:"jwt_secret_key"`
	Database    database `json:"database"`
}

type database struct {
	Driver string `json:"driver"`
	Url    string `json:"url"`
}
