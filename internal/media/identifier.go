package media

import (
	"fmt"
	"strings"
)

const (
	IdentifierTypeCID               IdentifierType = "cid"
	IdentifierTypeIPNS              IdentifierType = "ipns_name"
	IdentifierTypeTraditionalDomain IdentifierType = "traditional_domain"
	IdentifierTypeENSDomain         IdentifierType = "ens_domain"
	IdentifierTypeUnknown           IdentifierType = "unknown"
)

type IdentifierType string

type IndexIdentifier struct {
	Value string
}

func DetermineIdentifierType(identifier string) IdentifierType {
	if strings.Contains(identifier, ".") {
		if strings.HasSuffix(identifier, ".eth") {
			return IdentifierTypeENSDomain
		} else {
			return IdentifierTypeTraditionalDomain
		}
	} else if strings.HasPrefix(identifier, "Qm") {
		return IdentifierTypeCID
	} else if strings.HasPrefix(identifier, "k51") {
		return IdentifierTypeIPNS
	}

	return IdentifierTypeUnknown
}

func (indexIdentifier IndexIdentifier) Type() IdentifierType {
	return DetermineIdentifierType(indexIdentifier.Value)
}

func (identifierType IdentifierType) String() string {
	return string(identifierType)
}

func (indexIdentifier IndexIdentifier) ResolveToCID() (CID, error) {
	indexIdentifierType := indexIdentifier.Type()

	switch indexIdentifierType {
	case IdentifierTypeCID:
		return CID(indexIdentifier.Value), nil
	default:
		return CIDBlank, fmt.Errorf("resolving identifier type %q to a CID is not yet supported", indexIdentifierType)
	}
}
