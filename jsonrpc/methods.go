package jsonrpc

// Manifest describes information for the Prisma Engine Go generator for the Prisma CLI.
type Manifest struct {
	PrettyName         string   `json:"prettyName"`
	DefaultOutput      string   `json:"defaultOutput"`
	Denylist           []string `json:"denylist"`
	RequiresGenerators []string `json:"requiresGenerators"`
	RequiresEngines    []string `json:"requiresEngines"`
}

// ManifestResponse sets the response Prisma Engine Go returns when Prisma asks for the Manifest.
type ManifestResponse struct {
	Manifest Manifest `json:"manifest"`
}
