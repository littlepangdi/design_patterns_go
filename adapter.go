package main

import "fmt"

/*
	Convert the interface of a class to another interface that the client wants.
	The adapter pattern enables classes to work together that would otherwise not work together due to incompatible interfaces.

	Client--->requests for sth.--->Adapter converts request to Service language--->Service processes request.

*/

type player interface {
	PlayMusic(file string)
}

type MP3Player struct{}

func (p *MP3Player) PlayMusic(file string) {
	fmt.Printf("MP3 is playing music %s\n", file)
}

type GameSoundPlayer struct{}

func (g *GameSoundPlayer) PlayGameSound(file string) {
	fmt.Printf("GameSoundPlayer is playing %s\n", file)
}

type SoundAdapter struct {
	GameSoundPlayer
}

func (a *SoundAdapter) PlayMusic(file string) {
	fmt.Printf("convert music file into GameSoundPlayer file\n")
	a.GameSoundPlayer.PlayGameSound(file)
}

func Play(p player, file string) {
	p.PlayMusic(file)
}

func RunAdapter() {
	musicFile := "grow.mp3"
	Play(&MP3Player{}, musicFile)
	Play(&SoundAdapter{}, musicFile)
}
