package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot is type for robot
type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func genName() string {
	return string(rand.Intn(26)+'A') + string(rand.Intn(26)+'A') + fmt.Sprintf("%03d", rand.Intn(1000))
}

// Name names a robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = genName()
	}
	return r.name, nil
}

// Reset re-names a robot
func (r *Robot) Reset() (string, error) {
	r.name = genName()
	return r.name, nil
}
