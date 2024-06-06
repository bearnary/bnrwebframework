package bnrwebframework

import (
	"encoding/json"
	"os"
	"time"
)

type BuildDescription struct {
	Commit string `json:"commit"`
	Time   string `json:"time"`
}

func readBuildDescription() (buildDescription *BuildDescription) {
	data, err := os.ReadFile("./build.json")
	if err != nil {
		return &BuildDescription{
			Commit: "LOCAL",
			Time:   time.Now().String(),
		}
	}
	_ = json.Unmarshal(data, &buildDescription)
	return buildDescription
}
