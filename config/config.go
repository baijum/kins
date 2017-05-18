package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
)

// configuration represent all configurations
type configuration struct {
	DataDir           string   `default:"" split_words:"true"`
	HTTPAddress       string   `default:":8080" envconfig:"HTTP_ADDRESS"`
	TokenPrivateKey   string   `split_words:"true"`
	TokenPublicKey    string   `split_words:"true"`
	ClientID          string   `envconfig:"CLIENT_ID"`
	ClientSecret      string   `split_words:"true"`
	IdentityProviders []string `default:"google" split_words:"true"`
	RedirectURL       string   `envconfig:"REDIRECT_URL"`
}

// Config represent all configurations
var Config configuration

// DatabaseFile provides the location of database file
func (c *configuration) DatabaseFile() string {
	if c.DataDir == "" {
		p, _ := filepath.Abs(os.Args[0])
		p = filepath.Dir(p)
		p = filepath.Join(p, "data")
		os.Mkdir(p, 0755)
		c.DataDir = p
	}
	return c.DataDir + "kins.db"
}

func init() {
	err := envconfig.Process("kins", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
