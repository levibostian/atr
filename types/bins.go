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
	PostInstall *struct {
		Command string
	}
}

type Bins = []Bin
