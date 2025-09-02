package solver

import (
	"artist/internal/captcha/token"
	"time"
)

func Solve() {
	if token.GetTotal() >= token.Limit {
		return
	}

	token.IncrementQueued()

	// TODO: chromedp stuff
	time.Sleep(time.Second * 30)

	// TODO: solving error
	if false {
		token.DecrementQueued()
		go Solve()
		return
	}

	tk := token.Token{
		Tk: "",                                 // TODO
		Ex: token.CreateExpiration(time.Now()), // TODO
	}

	addToken(tk)
}

func addToken(tk token.Token) {
	token.AvailableTokens <- tk
	token.DecrementQueued()
	token.IncrementAvailable()
}
