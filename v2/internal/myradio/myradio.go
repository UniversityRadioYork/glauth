package myradio

import (
	"fmt"

	"github.com/UniversityRadioYork/myradio-go"
	c "github.com/glauth/glauth/v2/pkg/config"
)

func MyRadioAuthenticator(user *c.User, pw string, session *myradio.Session) error {
	user, err := session.UserCredentialsTest(user.Name, pw)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("login invalid")
	}

	return nil
}

func HandleMyRadio(apiKey string) (users []c.User, groups []c.Group) {
	groups = []c.Group{c.Group{
		Name:      "myradio-user",
		GIDNumber: 1350,
	}}

	myrSession, err := myradio.NewSession(apiKey)
	if err != nil {
		// TODO
		panic(err)
	}

	myrUsers, err := myrSession.GetThisYearsMembers()
	if err != nil {
		// TODO
		panic(err)
	}

	for _, u := range myrUsers {
		users = append(users, c.User{
			Name:         u.Email,
			PrimaryGroup: 1350,
			UIDNumber:    u.MemberID,
			GivenName:    u.Fname,
			SN:           u.Sname,
			PassAppCustom: func(user *c.User, pw string) error {
				return MyRadioAuthenticator(user, pw, myrSession)
			},
		})
	}

	return
}
