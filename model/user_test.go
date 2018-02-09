package model

import (
	"testing"

	"github.com/qetuantuan/jengo_recap/vo"
	"github.com/qetuantuan/jengo_recap/util"
)

func TestSetToken(t *testing.T) {
	token := "iamatoken"
	user := &User{
		Auths: []Auth{
			Auth{
				AuthSourceId: "myid",
			},
		},
		Scms: []Scm{
			Scm{
				AuthSourceId: "myid",
			},
		},
	}
	user.SetTokenEncrypted("myid", util.KeyCoder, token)
	plainText, err := util.AESDecode([]byte(util.KeyCoder), user.Auths[0].Token)
	if err != nil || string(plainText) != token {
		t.Logf("plainText: %v, Error: %v", plainText, err)
		t.Error("Wrong in Auth.Token")
	}
	plainText, err = util.AESDecode([]byte(util.KeyCoder), user.Scms[0].Token)
	if err != nil || string(plainText) != token {
		t.Logf("plainText: %v, Error: %v", plainText, err)
		t.Error("Wrong in Scm.Token")
	}
}

func TestToApiUser(t *testing.T) {
	token := "iamatoken"
	user := &User{
		Id: "123",
		Auths: []Auth{
			Auth{
				AuthSourceId: vo.AUTH_SOURCE_GITHUB,
				AuthBase: vo.AuthBase{
					Primary: true,
				},
			},
			Auth{
				AuthSourceId: "something",
			},
		},
		Scms: []Scm{
			Scm{
				AuthSourceId: vo.AUTH_SOURCE_GITHUB,
			},
		},
	}
	user.SetTokenEncrypted(vo.AUTH_SOURCE_GITHUB, util.KeyCoder, token)

	tmp, _ := util.AESEncode([]byte(util.KeyCoder), []byte(token))
	if string(user.Auths[0].Token) != string(tmp) {
		t.Errorf("token not decrypted. %v", user)
	}

	if u, err := user.ToApiUser(); err != nil {
		t.Fatalf("ToApiUser failed: %v", err)
	} else if u.Auth.AuthSource.Name != vo.AUTH_SOURCE_GITHUB {
		t.Error("auth source not initialized. %v", u)
	}
}
