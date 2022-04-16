package types

type Bin struct {
	Binary  string
	Version struct {
		Requirement string
		Command     *string
		EnvVars     *[]string
	}
	InstallCommand []string
}

type Bins = []Bin
