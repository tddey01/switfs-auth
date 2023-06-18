package jwtclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tddey01/switfs-auth/config"
	"github.com/tddey01/switfs-auth/core"
)

func TestLocalAuthClient(t *testing.T) {
	secret, err := config.RandSecret()
	assert.NoError(t, err)

	clientFromSecret, token, err := NewLocalAuthClientWithSecret(secret)
	if err != nil {
		t.Fatal(err)
	}

	testClientWithAdminPerm(t, clientFromSecret, string(token))

	client, token2, err := NewLocalAuthClient()
	if err != nil {
		t.Fatal(err)
	}
	testClientWithAdminPerm(t, client, string(token2))
}

func testClientWithAdminPerm(t *testing.T, client *LocalAuthClient, token string) {
	ctx := context.Background()
	permission, err := client.Verify(ctx, token)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, core.PermAdmin, permission)
}
