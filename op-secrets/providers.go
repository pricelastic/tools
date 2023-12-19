package main

import (
	"fmt"
	"os"

	"github.com/dvcrn/go-1password-cli/op"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
)

// Common data types
type (
	EnvsMap  map[string]string
	CacheMap map[string]map[string]string
)

// Cache key for the providers
var (
	dotenvCache = make(CacheMap)
)

// Get all keys from dotenv file
func GetAllKeysFromDotenv(vault string) ([]string, error) {
	var envs EnvsMap
	envs, err := godotenv.Read(vault)
	if err != nil {
		return nil, err
	}
	return lo.Keys[string, string](envs), nil
}

// Get secret value from a provider vault for a key
func GetSecretFromProvider(provider, vault, key string) (string, error) {
	switch provider {
	case "1password":
		return getOnePasswordSecret(vault)
	case "process_env":
		return getProcessEnvSecret(vault), nil
	case "dotenv":
		return getDotenvSecret(vault, key)
	default:
		return "", fmt.Errorf("invalid provider: '%s'", provider)
	}
}

// Get secret value for process_env provider
func getProcessEnvSecret(vaultKey string) string {
	return os.Getenv(vaultKey)
}

// Get secret value for dotenv provider
func getDotenvSecret(vault, key string) (string, error) {
	// Check if dotenv file data exists in cache
	cache, ok := dotenvCache[vault]
	if ok {
		return cache[key], nil
	}

	// Load dotenv file
	var envs EnvsMap
	envs, err := godotenv.Read(vault)
	if err != nil {
		return "", err
	}

	// Add dotenv file data to the cache
	dotenvCache[vault] = envs
	return envs[key], nil
}

// Get secret value for 1password provider
func getOnePasswordSecret(vault string) (string, error) {
	client := op.NewOpClient()

	value, err := client.Read("op://" + vault)
	if err != nil {
		return "", err
	}
	return value, nil
}
