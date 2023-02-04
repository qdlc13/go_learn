package model
//大写其他包可访问
type Student struct{
	Name string
	Age int
}

//小写其它包需要访问 需要用到工厂模式
type teacher struct{
	Name string
	Age int
}

func NewTeacher(n string,a int) *teacher{
	return &teacher{n,a}
}