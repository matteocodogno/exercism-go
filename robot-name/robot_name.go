package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	letterLength 	= 26
	numberLimit 	= 1000
	poolSize 		= 676000
)
var poolNames = generatePoolName()

func generatePoolName() []string {
	a, b := 'A', 'A'
	pool := make([]string, poolSize)
	var counter int

	for i := 0; i < letterLength; i ++ {
		for j := 0; j < letterLength; j ++ {
			for k := 0; k < numberLimit; k ++ {
				pool[counter] = string(a+int32(i)) + string(b+int32(j)) + fmt.Sprintf("%03d", k)
				counter ++
			}
		}
	}
	rand.Shuffle(poolSize, func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})

	return pool
}

type Robot struct {
	name string
}

func (r *Robot) Reset() {
	r.name = ""
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(poolNames) == 0 {
			return "", errors.New("new name is not available")
		}

		r.name = poolNames[0]
		poolNames = poolNames[1:]
	}

	return r.name, nil
}

