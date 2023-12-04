package autoload

/*
struct 后面 单引号括起来的，叫 tag
相当于 java 里的 注解
*/
type Admin struct {
	/*
		相当于一个字段上，有三个注解

		@mapstructure("enable")
		@json("enable")
		@yaml("enable")
		boolean enable
	*/
	Enable bool `mapstructure:"enable" json:"enable" yaml:"enable"`
	Auth   bool `mapstructure:"auth" json:"auth" yaml:"auth"`
}
