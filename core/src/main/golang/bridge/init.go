package bridge

import "github.com/Dreamacro/clash/constant"

import "github.com/kr328/cfa/profile"

import "github.com/Dreamacro/clash/tunnel"

func Init(home string) {
	constant.SetHomeDir(home)
}

func Reset() {
	profile.LoadDefault()
	tunnel.DefaultManager.Reset()
}
