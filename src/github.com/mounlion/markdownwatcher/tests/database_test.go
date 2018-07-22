package tests

import (
	"testing"
	"github.com/mounlion/markdownwatcher/database"
)

func TestReverse(t *testing.T) {
	database.UpdateSubscribe(123456, true)
}