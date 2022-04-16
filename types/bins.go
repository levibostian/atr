package types

type Bin struct {
	Binary  string
	Version struct {
		Requirement string
		Command     *string
		EnvVars     *[]string
	}
	InstallMethods []struct {
		Command       string
		UpdateCommand *string
	}
}

type Bins = []Bin
