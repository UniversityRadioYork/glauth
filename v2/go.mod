module github.com/UniversityRadioYork/glauth/v2

go 1.16

// Do not mistake /vendored for /vendor!
replace github.com/hydronica/toml => ./vendored/toml

replace github.com/glauth/glauth/v2 => ./

require (
	github.com/GeertJohan/yubigo v0.0.0-20190917122436-175bc097e60e
	github.com/UniversityRadioYork/myradio-go v0.0.0-20231128173314-37a34932a4b1 // indirect
	github.com/arl/statsviz v0.4.0
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/docopt/docopt-go v0.0.0-20180111231733-ee0de3bc6815
	github.com/fsnotify/fsnotify v1.4.9
	github.com/glauth/glauth/v2 v2.1.0
	github.com/hydronica/toml v0.4.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/nmcclain/ldap v0.0.0-20210720162743-7f8d1e44eeba
	github.com/pquerna/otp v1.3.0
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.28.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/yaegashi/msgraph.go v0.1.1-0.20200221123608-2d438cf2a7cc
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	gopkg.in/amz.v3 v3.0.0-20201001071545-24fc1eceb27b
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
