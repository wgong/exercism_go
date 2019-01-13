package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const totalNumberOfNames = 676000

// Robot is type for robot
type Robot struct {
	name string
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

func genUniqName() string {
	if len(mapNames) > int(0.9*totalNumberOfNames) {
		fmt.Println("90 %% of available names are used. Consider to expand naming format")
	}
	name := genName()
	for mapNames[name] {
		name = genName()
	}
	mapNames[name] = true
	return name
}

// Name names a robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = genUniqName() //genName()
	}
	return r.name, nil
}

// Reset re-names a robot
func (r *Robot) Reset() (string, error) {
	r.name = genUniqName() //genName()
	return r.name, nil
}
