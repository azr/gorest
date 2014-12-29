package gorest

var authorizers map[string]Authorizer

//Signiture of functions to be used as Authorizers
type Authorizer func(string, string) (bool, bool, SessionData)

//Registers an Authorizer for the specified realm.
func RegisterRealmAuthorizer(realm string, auth Authorizer) {
	if authorizers == nil {
		authorizers = make(map[string]Authorizer, 0)
	}

	if _, found := authorizers[realm]; !found {
		authorizers[realm] = auth
	}
}

//Returns the registred Authorizer for the specified realm.
func GetAuthorizer(realm string) (a Authorizer) {
	if authorizers == nil {
		authorizers = make(map[string]Authorizer, 0)
	}
	a, _ = authorizers[realm]
	return
}

//This is the default and exmaple authorizer that is used to authorize requests to endpints with no security realms.
//It always allows access and returns nil for SessionData.
func DefaultAuthorizer(id string, role string) (bool, bool, SessionData) {
	return true, true, nil
}
