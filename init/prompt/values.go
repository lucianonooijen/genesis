package prompt

var (
	ReplaceValueIdNameCapitalized = "name_capitalized"
	ReplaceValueIdNameLowercase   = "name_lowercase"
)

var replaceValuesBase = []ReplaceValue{
	{
		ID:           ReplaceValueIdNameCapitalized,
		OldValue:     "Genesis",
		Description:  "Project name without spaces, with capitals",
		ReplaceOrder: 99,
	},
	{
		ID:           ReplaceValueIdNameLowercase,
		OldValue:     "genesis",
		Description:  "Project name without spaces, lowercase",
		ReplaceOrder: 98,
	},
	{
		OldValue:     "nl.bytecode.genesis",
		Description:  "Android package name",
		ReplaceOrder: 98,
	},
	{
		OldValue:     "registry.digitalocean.com/dawny/genesis-server",
		Description:  "Docker registry for server",
		ReplaceOrder: 6,
	},
	{
		OldValue:     "git.bytecode.nl/bytecode/genesis/server",
		Description:  "Golang module name for server",
		ReplaceOrder: 5,
	},
	{
		OldValue:     "development@genesis.bytecode.dev",
		Description:  "Default from-email",
		ReplaceOrder: 4,
	},
	{
		OldValue:     "https://placekitten.com/400/400",
		Description:  "Placeholder image being used",
		ReplaceOrder: 1,
	},
	{
		OldValue:     "https://bytecode.nl",
		Description:  "Link in emails",
		ReplaceOrder: 2,
	},
}
