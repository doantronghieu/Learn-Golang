package logic_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"EngineerPro/internal/database"
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
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	userDataAccessor := database.NewMockUserDataAccessor(mockController)
	userDataAccessor.EXPECT().GetUser(gomock.Any(), uint64(1)).Return(database.User{
		ID:   1,
		Name: "Harry",
	}, nil).AnyTimes()

	user, err := userDataAccessor.GetUser(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, database.User{
		ID:   1,
		Name: "Harry",
	}, user)
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
