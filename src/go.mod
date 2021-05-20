module github.com/sinhadotabhinav/cryptogeek

go 1.16

require (
	github.com/sinhadotabhinav/cryptogeek/api v0.0.0-00010101000000-000000000000
	github.com/sinhadotabhinav/cryptogeek/configs v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1 // indirect
)

replace github.com/sinhadotabhinav/cryptogeek/api => ./api

replace github.com/sinhadotabhinav/cryptogeek/configs => ./configs
