package robotname

import (
	"errors"
	"github.com/zach-klippenstein/goregen"
	//"fmt"
	"math/rand"
)

var uppercaseLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(strSize int) string {
	result := make([]rune, strSize)
	for i := range result {
		result[i] = uppercaseLetters[rand.Intn(len(uppercaseLetters))]
	}

	return string(result)
}

type Robot struct {
	name string
}

var robots = map[string]int{}
func generateRobotName() (string, error) {
	var tmpName string
	const maxNames = 676000

	if len(robots) == maxNames {
		return "", errors.New("new name is not available")
	}

	for {
		tmpName, _ = regen.Generate("[A-Z]{2}[0-9]{3}")
		if _, ok := robots[tmpName]; !ok {
			robots[tmpName] = 0
			break
		}
	}

	return tmpName, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateRobotName()
		if err != nil {
			return "", err
		}

		r.name = name
	}

	return r.name, nil
}

