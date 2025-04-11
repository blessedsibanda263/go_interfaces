package examples

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strings"
)

type Shuffleable interface {
	contents() string
	shuffle(r *rand.Rand)
}

type shuffleString string

func (s *shuffleString) shuffle(r *rand.Rand) {
	tmp := strings.Split(string(*s), "")
	r.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
	*s = shuffleString(strings.Join(tmp, ""))
}

func (s *shuffleString) contents() string {
	return string(*s)
}

func NewShuffleString(init string) *shuffleString {
	var s shuffleString = shuffleString(init)
	return &s
}

type shuffleSlice []interface{}

func (sl shuffleSlice) contents() string {
	data, _ := json.Marshal(sl)
	return fmt.Sprintf("%v", string(data))
}

func (sl shuffleSlice) shuffle(r *rand.Rand) {
	r.Shuffle(len(sl), func(i, j int) {
		sl[i], sl[j] = sl[j], sl[i]
	})
}

func ShuffleableInterfaceExamples() {
	r := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))

	var myShuffle Shuffleable

	myShuffle = NewShuffleString("my name is blessed sibanda")
	myShuffle.shuffle(r)
	fmt.Println(myShuffle.contents())

	myShuffle = &shuffleSlice{1, 2, 3, 4, 5}
	myShuffle.shuffle(r)
	fmt.Println(myShuffle.contents())
}
