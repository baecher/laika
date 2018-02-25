package test

import (
	"net/http/httptest"
	"testing"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/client"
	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestGetEnabledFeatures(t *testing.T) {
	state := models.NewState()
	state.Environments = []models.Environment{
		models.Environment{
			Name: "e1",
		},
		models.Environment{
			Name: "e 2",
		},
	}
	state.Features = []models.Feature{
		models.Feature{
			Name: "f1",
		},
		models.Feature{
			Name: "f2",
		},
	}
	state.Enabled = map[models.EnvFeature]bool{
		models.EnvFeature{"e 2", "f2"}: true,
	}

	store, _ := store.NewMemoryStore(state)

	handler, err := api.NewServer(api.ServerConfig{
		RootUsername: "root",
		RootPassword: "root",
		Store:        store,
	})
	require.NoError(t, err)

	s := httptest.NewServer(handler)
	require.NotNil(t, s)

	c, err := client.NewClient(client.Config{
		Addr:        s.URL,
		Username:    "root",
		Password:    "root",
		Environment: "e 2",
	})
	require.NoError(t, err)

	require.False(t, c.IsEnabled("f1", false))
	require.True(t, c.IsEnabled("f2", false))
}
