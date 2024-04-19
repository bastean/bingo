package config

type Flag struct {
	Name string
	Run  func()
}
