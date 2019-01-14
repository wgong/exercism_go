package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const totalNumberOfNames = 676000

// Robot is type for robot
type Robot struct {
	name string
	err  error
}

// ensure uniqueness
var mapNames = map[string]bool{"": false}

func init() {
	rand.Seed(time.Now().UnixNano())
	mapNames = map[string]bool{}
}

func genName() string {
	return string(rand.Intn(26)+'A') + string(rand.Intn(26)+'A') + fmt.Sprintf("%03d", rand.Intn(1000))
}

func genUniqName() (string, error) {
	if len(mapNames) >= totalNumberOfNames {
		return "", errors.New("ran out unique names")
	}
	name := genName()
	for mapNames[name] {
		name = genName()
	}
	mapNames[name] = true
	return name, nil
}

// Name names a robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name, r.err = genUniqName()
	}
	return r.name, r.err
}

// Reset re-names a robot
func (r *Robot) Reset() (string, error) {
	r.name, r.err = genUniqName()
	return r.name, r.err
}
