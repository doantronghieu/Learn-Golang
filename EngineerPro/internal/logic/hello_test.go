package logic_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"EngineerPro/internal/logic"
)

var currentTime time.Time

func TestSayHello(t *testing.T) {
	t.Parallel()

	output := logic.SayHello("World")
	assert.Equal(t, "Hello World", output, "incorrect output")

	output = logic.SayHello("")
	assert.Equal(t, "", output, "incorrect output")
	
}

func TestCurrentTime(t *testing.T) {
	t.Log(currentTime)
}

func TestMain(m *testing.M) {

	var err error

	timeEnvVar := os.Getenv("TIME")
	if timeEnvVar == "" {
		currentTime = time.Now()
	} else {
		currentTime, err = time.Parse(time.RFC3339, timeEnvVar)
		if err != nil {
			return
		}
	}

	m.Run()
}
