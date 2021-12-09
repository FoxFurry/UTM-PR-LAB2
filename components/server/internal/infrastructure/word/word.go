package word

import (
	"github.com/tjarratt/babble"
	"sync"
)

var babbler babble.Babbler
var babblerOnce sync.Once


func Say() string {
	babblerOnce.Do(func(){
		babbler = babble.NewBabbler()

	})
	return babbler.Babble()
}
