package bodyshop

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/db"

	"github.com/sirupsen/logrus"
)

type App struct {
	Config      *conf.Config
	Logger      *logrus.Entry
	TemplateDAO db.TemplateDAO
}
