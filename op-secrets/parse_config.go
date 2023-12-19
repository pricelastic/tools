package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Providers []struct {
		Provider string `yaml:"provider" validate:"required,oneof=1password process_env dotenv"`
		Order    uint8  `yaml:"order,omitempty" validate:"required"`
		// Envs and EnvsAll is mutually exclusive
		Envs map[string]struct {
			Vault       string `yaml:"vault"`
			IgnoreError bool   `yaml:"ignore_error"`
		} `yaml:"envs" validate:"required_unless=Provider dotenv"`
		// EnvsAll is only allowed for dotenv provider
		EnvsAll struct {
			Vault       string `yaml:"vault"`
			IgnoreError bool   `yaml:"ignore_error"`
		} `yaml:"envs_all" validate:"required_if=Provider dotenv,excluded_with=Envs"`
	} `yaml:"providers" validate:"required,unique=Order,dive"`
}

type Secret struct {
	Provider    string
	Order       uint8
	Vault       string
	Key         string
	Value       string
	IgnoreError bool
}

// Parse the YAML config file
func ParseYamlConfig(configFile string) ([]Secret, error) {
	file, readErr := os.ReadFile(configFile)
	if readErr != nil {
		return nil, readErr
	}

	// Replace ${VAR} or $VAR in the yaml content
	content := os.ExpandEnv(string(file))

	// Initialize Config struct
	var conf YamlConfig

	// KnownFields ensures that the keys in decoded mappings exist
	// https://github.com/go-yaml/yaml/issues/639
	decoder := yaml.NewDecoder(bytes.NewReader([]byte(content)))
	decoder.KnownFields(true)
	if err := decoder.Decode(&conf); err != nil && err != io.EOF {
		return nil, err
	}

	// Validate the struct with the rules specified as tags
	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		// Format and print only first error
		validationError := err.(validator.ValidationErrors)[0]
		params := ""
		if validationError.Param() != "" {
			params = fmt.Sprintf("[%s]", validationError.Param())
		}
		return nil, fmt.Errorf("invalid config: '%s' %s %s",
			strings.TrimLeft(validationError.StructNamespace(), "YamlConfig."),
			validationError.Tag(), params)
	}

	// Prepare & normalize config to build secrets slice
	var secrets []Secret
	for _, provider := range conf.Providers {
		// For dotenv provider, get all keys and fake the individual Envs configuration
		if provider.Provider == "dotenv" {
			keys, err := GetAllKeysFromDotenv(provider.EnvsAll.Vault)
			if err != nil {
				if !provider.EnvsAll.IgnoreError {
					return nil, err
				}
			}
			for _, key := range keys {
				secrets = append(secrets, Secret{
					Provider:    provider.Provider,
					Order:       provider.Order,
					Vault:       provider.EnvsAll.Vault,
					Key:         key,
					IgnoreError: provider.EnvsAll.IgnoreError,
				})
			}
		} else {
			for key, value := range provider.Envs {
				// For process_env provider, vault is the key for the process
				vault, _ := lo.Coalesce(value.Vault, key)

				secrets = append(secrets, Secret{
					Provider:    provider.Provider,
					Order:       provider.Order,
					Vault:       vault,
					Key:         key,
					IgnoreError: value.IgnoreError,
				})
			}
		}
	}

	// Get secrets from the parsed YAML config
	// Loop through each of the secrets and get Value from providers
	var secretsWithValues []Secret
	for _, secret := range secrets {
		value, err := GetSecretFromProvider(secret.Provider, secret.Vault, secret.Key)

		// An error or if the value is blank
		if err != nil || value == "" {
			if !secret.IgnoreError {
				if err == nil {
					err = fmt.Errorf("provider '%s' has no key '%s'", secret.Provider, secret.Key)
				}
				return nil, err
			}
		} else {
			secret.Value = value
			secretsWithValues = append(secretsWithValues, secret)
		}
	}

	// Once we have the secretsWithValues slice, let's order & remove duplicates
	sort.Slice(secretsWithValues, func(i, j int) bool {
		return secretsWithValues[i].Order > secretsWithValues[j].Order
	})

	uniqSecrets := lo.UniqBy[Secret](secretsWithValues, func(s Secret) string {
		return s.Key
	})

	return uniqSecrets, nil
}
