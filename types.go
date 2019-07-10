package gonexus

// Asset struct
type Asset struct {
	DownloadURL string
	Path        string
	ID          string
	Repository  string
	Format      string
	CheckSum    interface{}
}

// Component struct
type Component struct {
	ID         string
	Repository string
	Format     string
	Group      string
	Name       string
	Version    string
	Assets     []Asset
}

// Search struct
type Search struct {
	Items             []Component
	ContinuationToken string
}
