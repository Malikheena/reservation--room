package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/malikheena/heena2-main/internal/models"
)

// AppConfig holds the application config
// type AppConfig struct {
// 	UseCache      bool
// 	TemplateCache map[string]*template.Template
// 	InfoLog       *log.Logger
// }

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Errorlog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan       chan models.MailData
}
