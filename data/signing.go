package data

import (
	"github.com/rubblelabs/ripple/crypto"
)

func Sign(s Signer, key crypto.Key, sequence *uint32) error {
	s.InitialiseForSigning()
	hash, msg, err := SigningHash(s)
	if err != nil {
		return err
	}
	sig, err := crypto.Sign(key.Private(sequence), hash.Bytes(), msg)
	if err != nil {
		return err
	}
	hash, _, err = Raw(s)
	if err != nil {
		return err
	}
	copy(s.GetHash().Bytes(), hash.Bytes())
	copy(s.GetPublicKey().Bytes(), key.Public(sequence))
	*s.GetSignature() = VariableLength(sig)
	return nil
}

func CheckSignature(s Signer) (bool, error) {
	hash, msg, err := SigningHash(s)
	if err != nil {
		return false, err
	}
	return crypto.Verify(s.GetPublicKey().Bytes(), hash.Bytes(), msg, s.GetSignature().Bytes())
}
