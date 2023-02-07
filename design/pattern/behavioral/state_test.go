package behavioral

import "testing"

func TestState(t *testing.T) {
	player := NewPlayer()
	player.clickNext()
	player.clickPrevious()


	player.clickPlay()
	player.clickNext()
	player.clickPrevious()

	player.clickLock()
	player.clickPrevious()

	player.clickPlay()
	player.clickNext()
	player.clickPrevious()

	player.clickPlay()
	player.clickNext()
	player.clickPrevious()
}
