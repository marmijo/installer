package joiner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
)

const (
	addNodesParamsFile = ".addnodesparams"
)

// AddNodesConfig is used to store the current configuration
// for the command.
type AddNodesConfig struct {
	File   *asset.File
	Params Params
}

// Params is used to store the command line parameters.
type Params struct {
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Save stores the current parameters on disk.
func (p *Params) Save(assetsDir string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	fileName := filepath.Join(assetsDir, addNodesParamsFile)
	return os.WriteFile(fileName, data, 0o600)
}

// Name returns the human-friendly name of the asset.
func (*AddNodesConfig) Name() string {
	return "AddNodes Config"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*AddNodesConfig) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate it's empty for this asset, always loaded from disk.
func (*AddNodesConfig) Generate(dependencies asset.Parents) error {
	return nil
}

// Files returns the files generated by the asset.
func (a *AddNodesConfig) Files() []*asset.File {
	if a.File != nil {
		return []*asset.File{a.File}
	}
	return []*asset.File{}
}

// Load returns agent config asset from the disk.
func (a *AddNodesConfig) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(addNodesParamsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to load %s file: %w", addNodesParamsFile, err)
	}

	params := &Params{}
	if err := json.Unmarshal(file.Data, params); err != nil {
		return false, fmt.Errorf("failed to unmarshal %s: %w", addNodesParamsFile, err)
	}

	a.Params = *params
	return true, nil
}
