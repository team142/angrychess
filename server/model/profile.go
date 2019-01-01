package model

type Profile struct {
	Nick string
	Out  chan []byte
	In   chan []byte
}
