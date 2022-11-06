package media

import (
	"database/sql/driver"
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
	ValueString string
}

func NewIndexIdentifierFromString(valueString string) IndexIdentifier {
	return IndexIdentifier{
		ValueString: valueString,
	}
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

func (indexIdentifier *IndexIdentifier) IdentifierType() IdentifierType {
	return DetermineIdentifierType(indexIdentifier.ValueString)
}

func (identifierType *IdentifierType) String() string {
	return string(*identifierType)
}

func (indexIdentifier *IndexIdentifier) ResolveToCID() (CID, error) {
	indexIdentifierType := indexIdentifier.IdentifierType()

	switch indexIdentifierType {
	case IdentifierTypeCID:
		return CID(indexIdentifier.ValueString), nil
	default:
		return CIDBlank, fmt.Errorf("resolving identifier type %q to a CID is not yet supported", indexIdentifierType)
	}
}

func (indexIdentifier *IndexIdentifier) Scan(value interface{}) error {
	valueString, ok := value.(string)
	if !ok {
		return fmt.Errorf("could not convert column value to string")
	}

	indexIdentifier.ValueString = valueString

	return nil
}

func (indexIdentifier IndexIdentifier) Value() (driver.Value, error) {
	return indexIdentifier.ValueString, nil
}
