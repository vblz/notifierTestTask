module github.com/vblz/notifierTestTask

go 1.17

replace github.com/vblz/notifierTestTask/notifier => ./notifier

require (
	github.com/go-pkgz/lgr v0.10.4
	github.com/jessevdk/go-flags v1.5.0
	github.com/stretchr/testify v1.7.1
	github.com/vblz/notifierTestTask/notifier v0.0.0-00010101000000-000000000000
	golang.org/x/time v0.0.0-20220411224347-583f2d630306
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
