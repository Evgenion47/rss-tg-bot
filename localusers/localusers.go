package localusers

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadUsers() []int64 {
	var tmp []int64
	source, err := ioutil.ReadFile("./localusers/localusers.yaml")
	if err != nil {
		log.Println(err)
		source, err = ioutil.ReadFile("/app/localusers/localusers.yaml")
		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(source, &tmp)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return tmp
}

func WriteUsers(in []int64) {
	data, err := yaml.Marshal(&in)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile("./localusers/localusers.yaml", data, 0)
	if err != nil {
		log.Println(err)
		err = ioutil.WriteFile("/app/localusers/localusers.yaml", data, 0)
		if err != nil {
			panic(err)
		}
	}
}
