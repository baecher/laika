package models

import (
	"time"
)

type User struct {
	Username     string    `json:"name"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type Environment struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Feature struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type EnvFeature struct {
	Env     string
	Feature string
}

type State struct {
	Users        []User
	Environments []Environment
	Features     []Feature
	Enabled      map[EnvFeature]bool
}

func NewState() *State {
	return &State{
		Environments: []Environment{},
		Features:     []Feature{},
		Users:        []User{},
		Enabled:      map[EnvFeature]bool{},
	}

}

func (s *State) GetFeatures(env string) map[string]int {
	if s.getEnvByName(env) == nil {
		return nil
	}

	features := map[string]int{}

	for _, feature := range s.Features {
		ok, status := s.Enabled[EnvFeature{env, feature.Name}]
		if ok && status {
			features[feature.Name] = 1
		} else {
			features[feature.Name] = 0
		}
	}

	return features
}

func (s *State) getFeatureByName(name string) *Feature {
	for _, feature := range s.Features {
		if feature.Name == name {
			return &feature
		}
	}

	return nil
}

func (s *State) getEnvByName(name string) *Environment {
	for _, env := range s.Environments {
		if env.Name == name {
			return &env
		}
	}

	return nil
}

func (s *State) getUserByName(username string) *User {
	for _, user := range s.Users {
		if user.Username == username {
			return &user
		}
	}

	return nil
}
