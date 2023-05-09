package myradio

import (
	"github.com/UniversityRadioYork/myradio-go"
	c "github.com/glauth/glauth/v2/pkg/config"
)

func AllowAccess(_ *c.User, _ string) error {
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
			Name:          u.Email,
			PrimaryGroup:  1350,
			UIDNumber:     u.MemberID,
			GivenName:     u.Fname,
			SN:            u.Sname,
			PassAppCustom: AllowAccess,
		})
	}

	return
}
