package characters

import (
	"math/rand"
	"testing"
	"time"
)
import "fmt"

func TestSkill(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	guy := RandPerson()
	fmt.Println(guy.GetWeaponSkill())
}
