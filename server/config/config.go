package config

type Server struct {
	Zap    Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System          `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
