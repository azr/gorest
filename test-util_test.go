package gorest

import (
	//"log"
	"testing"
)

func TestingAuthorizer(id string, role string) (bool, bool, SessionData) {
	if idsInRealm == nil {
		idsInRealm = make(map[string][]string, 0)
		idsInRealm["12345"] = []string{"var-user", "string-user", "post-user"}
		idsInRealm["fox"] = []string{"postInt-user"}
	}

	if roles, found := idsInRealm[id]; found {
		for _, r := range roles {
			if role == r {
				return true, true, nil
			}
		}
		return true, false, nil
	}

	return false, false, nil
}

func AssertEqual(given interface{}, expecting interface{}, compared string, t *testing.T) {
	if expecting != given {
		t.Error("Fail Assert:", compared, " Expecting:", expecting, "; but is:", given)
	} else {
		//log.Println("Pass Assert:", compared)
	}
}
