package cmd

import (
	"CacheCLI/kvcache"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}
