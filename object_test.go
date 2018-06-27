package jsonq

import (
	"log"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var (
	smallFixtureValue  *Value
	mediumFixtureValue *Value
	largeFixtureValue  *Value

	requestSmallSimple  *Level
	requestSmallMedium  *Level
	requestSmallHard    *Level
	requestMediumSimple *Level
	requestMediumMedium *Level
	requestMediumHard   *Level
	requestLargeSimple  *Level
	requestLargeMedium  *Level
	requestLargeHard    *Level
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	var p1 Parser
	v1, err := p1.Parse(smallFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	smallFixtureValue = v1

	var p2 Parser
	v2, err := p2.Parse(mediumFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	mediumFixtureValue = v2

	var p3 Parser
	v3, err := p3.Parse(largeFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	largeFixtureValue = v3

	requestSmallSimple = MustParseCMD(`(st>=1){st}`)
	requestSmallMedium = MustParseCMD(`(gr == 0){st,sid,tt}`)
	requestSmallHard = MustParseCMD(`{st,sid,tt,users(name == "Leonid"){name}}`)
	requestMediumSimple = MustParseCMD(`{person{name{fullName}}}`)
	requestMediumMedium = MustParseCMD(`{person{name{fullName},email,geo{city,state},bio}}`)
	requestMediumHard = MustParseCMD(`{person{name{fullName},email,geo{city,state},bio},users(id > 20){username}}`)
	requestLargeSimple = MustParseCMD(`{users{username}}`)
	requestLargeMedium = MustParseCMD(`{topics{topics{title,fancy_title}}}`)
	requestLargeHard = MustParseCMD(`{users{username},topics{topics{posters{description}}}}`)
}

func TestKeepRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		TestKeepSmallRequestSuccess(t)
		TestKeepMediumRequestSuccess(t)
		TestKeepLargeRequestSuccess(t)
	})

	t.Run("error", func(t *testing.T) {
		// testParseRawNumberError(t, "xyz", "xyz")
	})
}

func TestKeepSmallRequestSuccess(t *testing.T) {
	t.Helper()

	KeepSmallSimple()
	KeepSmallMedium()
	KeepSmallHard()

}

func TestKeepMediumRequestSuccess(t *testing.T) {
	t.Helper()

	KeepMediumSimple()
	KeepMediumMedium()
	KeepMediumHard()

}
func TestKeepLargeRequestSuccess(t *testing.T) {
	t.Helper()

	KeepLargeSimple()
	KeepLargeMedium()
	KeepLargeHard()

}

//
// SMALL DATASET
//

func KeepSmallSimple() {

	_, err := smallFixtureValue.Keep(*requestSmallSimple)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepSmallMedium() {

	_, err := smallFixtureValue.Keep(*requestSmallMedium)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepSmallHard() {
	_, err := smallFixtureValue.Keep(*requestSmallHard)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

//
// MEDIUM DATASET
//

func KeepMediumSimple() {
	_, err := mediumFixtureValue.Keep(*requestMediumSimple)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepMediumMedium() {
	_, err := mediumFixtureValue.Keep(*requestMediumMedium)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepMediumHard() {
	_, err := mediumFixtureValue.Keep(*requestMediumHard)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

//
// LARGE DATASET
//

func KeepLargeSimple() {
	_, err := largeFixtureValue.Keep(*requestLargeSimple)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepLargeMedium() {
	_, err := largeFixtureValue.Keep(*requestLargeMedium)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}

func KeepLargeHard() {
	_, err := largeFixtureValue.Keep(*requestLargeHard)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
}
