package main

import "github.com/yoctoMNS/soft_engine/core"

func main() {
	core.GetEngineInstance().Init()

	for core.GetEngineInstance().IsRunning() {
		core.GetEngineInstance().Events()
		core.GetEngineInstance().Update()
		core.GetEngineInstance().Render()
	}

	core.GetEngineInstance().Clean()
}
