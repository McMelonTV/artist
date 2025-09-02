package captcha

import (
	"artist/internal/captcha/solver"
	"artist/internal/captcha/token"
	"time"
)

var WPlaceSiteKey = "0x4AAAAAABpqJe8FO0N84q0F"

func RunSolver() {
	initCounter()

	for range token.InitAmount {
		go solver.Solve()
	}

	go func() {
		time.Sleep(Window)
		needed := getRequestCount()
		difference := needed - token.GetTotal()
		if difference > 0 {
			for range difference {
				go solver.Solve()
			}
		}
	}()
}

func RequestTokenSync() token.Token {
	trackRequest()

	return grabToken()
}

func RequestTokenAsync() chan token.Token {
	trackRequest()

	tokenChannel := make(chan token.Token, 1)

	go func() {
		tokenChannel <- grabToken()
	}()

	return tokenChannel
}

func grabToken() token.Token {
	queueTokenIfNone()

	for {
		tk := <-token.AvailableTokens
		token.DecrementAvailable()

		if tk.Valid() {
			return tk
		}
		go solver.Solve()
	}
}

func queueTokenIfNone() {
	total := token.GetTotal()
	if total < 0 {
		panic("how the fuck did we get to negative total tokens")
	}
	if total == 0 {
		go solver.Solve()
	}
}
