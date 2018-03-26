package test

import "testing"
import "../src/galio/entities"

func TestEntities(t *testing.T)  {
	champion := entities.Champion{Id: 1, Key: "a", Name: "b"}
	if champion.Id != 1 {
		t.Errorf("Champion does not match ID")
	}
}
