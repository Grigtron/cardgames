package main

import (
	"errors"
	"fmt"
)

func commandPlayTurn(args ...string) error {
	if currentGame == nil {
		return fmt.Errorf("no game in progress. start one with 'play war'")
	}

	if args[0] == "playturn" {
		err := currentGame.PlayTurn()
		if err != nil {
			return fmt.Errorf("error playing turn: %w", err)
		}

		return nil
		
	}
	return errors.New("unknown command for war")
}