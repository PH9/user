// +build integration

package user_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/PH9/user/user"
)

func TestIntegrationGetUser(t *testing.T) {
	u, err := user.Get(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	fmt.Print(u)
}
