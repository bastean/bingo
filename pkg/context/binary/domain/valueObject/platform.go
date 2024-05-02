package valueObject

type Platform struct {
	*OS
	*Arch
}

func NewPlatform(os, arch string) *Platform {
	return &Platform{
		OS:   NewOS(os),
		Arch: NewArch(arch),
	}
}
