package schedule

import (
	"io/ioutil"
	
	"os"
	"strings"
	"BearApp/internal/bootstrap"
	"github.com/naoina/toml"
)


func loadSchedule() ([]*CronJob, error) {
	configDir := bootstrap.GetAppRoot() + "/config/schedule/" + bootstrap.GetAppEnv()
	dir, err := os.Open(configDir)
	if err != nil {
		return nil, err
	}

	var fileList []os.FileInfo
	fileList, err = dir.Readdir(-1)
	if err != nil {
		dir.Close()
		return nil, err
	}
	defer dir.Close()

	var jobs []*CronJob

	for i := range fileList {
		file := fileList[i]

		if strings.HasSuffix(file.Name(), ".toml") {
			tomlData, readFileErr := ioutil.ReadFile(configDir + "/" + file.Name())
			if readFileErr != nil {
				return nil, readFileErr
			}

			var jobsConf struct {
				Jobs []*CronJob `toml:"job"`
			}
			err := toml.Unmarshal(tomlData, &jobsConf)
			if err != nil {
				return nil, err
			}
			jobs = append(jobs, jobsConf.Jobs...)
		}
	}

	return jobs, nil
}
