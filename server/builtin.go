package server

import (
	"github.com/Woop-Chain/woopchain/consensus"
	consensusDev "github.com/Woop-Chain/woopchain/consensus/dev"
	consensusDummy "github.com/Woop-Chain/woopchain/consensus/dummy"
	consensusIBFT "github.com/Woop-Chain/woopchain/consensus/ibft"
	"github.com/Woop-Chain/woopchain/secrets"
	"github.com/Woop-Chain/woopchain/secrets/awsssm"
	"github.com/Woop-Chain/woopchain/secrets/gcpssm"
	"github.com/Woop-Chain/woopchain/secrets/hashicorpvault"
	"github.com/Woop-Chain/woopchain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
