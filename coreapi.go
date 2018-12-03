// Package ipfscorehttp implements the core_api interface with an IPFS HTTP api backing
// currently we're implementing on an as-needed basis
package ipfscorehttp

import (
	"net/http"

	coreiface "gx/ipfs/QmUJYo4etAQqFfSS2rarFAE97eNGB8ej64YkRT2SmsYD4r/go-ipfs/core/coreapi/interface"
)

// CoreHTTP implements github.com/ipfs/go-ipfs/core/coreapi/interface.CoreAPI
// by talking to an IPFS daemon via the HTTP API
type CoreHTTP struct {
	url string
	cli *http.Client
}

// NewCoreHTTP creates a core HTTP module
func NewCoreHTTP(urlstr string) CoreHTTP {
	return CoreHTTP{
		url: urlstr,
		cli: http.DefaultClient,
	}
}

// Block returns the BlockAPI interface implementation backed by an ipfs HTTP api string
func (api CoreHTTP) Block() coreiface.BlockAPI {
	return (BlockAPI)(api)
}

// Dag returns the DagAPI interface implementation backed by an ipfs HTTP api string
func (api CoreHTTP) Dag() coreiface.DagAPI {
	return (DagAPI)(api)
}
