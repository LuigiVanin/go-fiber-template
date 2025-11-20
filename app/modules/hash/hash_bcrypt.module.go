package hash

import "boilerplate/infra/configuration"

type HashBcryptModule struct {
	HashService IHashService
}

func NewHashBcryptModule(config configuration.Config) *HashBcryptModule {
	return &HashBcryptModule{
		HashService: New(config),
	}
}
