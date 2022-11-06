package media

type CID string

const CIDBlank CID = ""

func (cID CID) Blank() bool {
	return len(string(cID)) == 0
}

func (cID CID) String() string {
	return string(cID)
}

func (cID CID) Valid() bool {
	// TODO (@JaksonKallio): this validation sucks
	return len(cID.String()) > 0
}
