package app

import (
	"larana.tech/go/electrostatic/meta"
	"larana.tech/go/electrostatic/types"
)

func ParsePageInfo(root string, f []byte) (*types.Page, error) {
	text := string(f)

	m, err := meta.Parse(text)

	if err != nil {
		return &types.Page{}, err
	}

	return &types.Page{
		Content: []byte(m.Text),
		Meta:    m.Meta,
	}, nil
}
