package config

type Mysql struct {
	// ,inline 内联结构体内容到外层
	// ,squash 压平 zap 配置文件内容到内层中
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
