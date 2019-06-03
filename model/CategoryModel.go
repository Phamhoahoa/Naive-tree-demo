package model
type Category struct{
	Id int
	Name string
	Code string
	Show bool
	Parent int
}
type Tree struct{
	Parent int
	Child int
}