package main

import "github.com/yoctoMNS/soft_engine/core"

func main() {
	core.GetInstance().Init()

	for core.GetInstance().IsRunning() {
		core.GetInstance().Events()
		core.GetInstance().Update()
		core.GetInstance().Render()
	}

	core.GetInstance().Clean()
}
