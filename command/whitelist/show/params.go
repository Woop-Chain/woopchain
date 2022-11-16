package show

import (
	"fmt"

	"github.com/Woop-Chain/woopchain/chain"
	"github.com/Woop-Chain/woopchain/command"
	"github.com/Woop-Chain/woopchain/helper/config"
	"github.com/Woop-Chain/woopchain/types"
)

const (
	chainFlag = "chain"
)

var (
	params = &showParams{}
)

type showParams struct {
	// genesis file path
	genesisPath string

	// deployment whitelist
	whitelists Whitelists
}

type Whitelists struct {
	deployment []types.Address
}

func (p *showParams) initRawParams() error {
	// init genesis configuration
	if err := p.initWhitelists(); err != nil {
		return err
	}

	return nil
}

func (p *showParams) initWhitelists() error {
	// import genesis configuration
	genesisConfig, err := chain.Import(p.genesisPath)
	if err != nil {
		return fmt.Errorf(
			"failed to load chain config from %s: %w",
			p.genesisPath,
			err,
		)
	}

	// fetch whitelists
	deploymentWhitelist, err := config.GetDeploymentWhitelist(genesisConfig)
	if err != nil {
		return err
	}

	// set whitelists
	p.whitelists = Whitelists{
		deployment: deploymentWhitelist,
	}

	return nil
}

func (p *showParams) getResult() command.CommandResult {
	result := &ShowResult{
		Whitelists: p.whitelists,
	}

	return result
}
