package anaconda

import (
	"net/url"
)

// Verify the credentials by making a very small request
func (a TwitterApi) VerifyCredentials() (ok bool, err error) {
	v := cleanValues(nil)
	v.Set("include_entities", "false")
	v.Set("skip_status", "true")

	_, err = a.GetSelf(v)
	return err == nil, err
}

// Get the user object for the authenticated user. Requests /account/verify_credentials
func (a TwitterApi) GetSelf(v url.Values) (u User, err error) {
	v = cleanValues(v)
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account/verify_credentials.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}
