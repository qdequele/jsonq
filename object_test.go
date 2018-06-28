package jsonq

import (
	"log"
	"testing"
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

// Needed for bench on Keep and Check. We assume that the parsing was previously done.
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

	requestSmallSimple = MustparseQuery(`(st>=1){st}`)
	requestSmallMedium = MustparseQuery(`(gr == 0){st,sid,tt}`)
	requestSmallHard = MustparseQuery(`{st,sid,tt,users(name == "Leonid"){name}}`)
	requestMediumSimple = MustparseQuery(`{person{name{fullName}}}`)
	requestMediumMedium = MustparseQuery(`{person{name{fullName},email,geo{city,state},bio}}`)
	requestMediumHard = MustparseQuery(`{person{name{fullName},email,geo{city,state},bio},users(id > 20){username}}`)
	requestLargeSimple = MustparseQuery(`{users{username}}`)
	requestLargeMedium = MustparseQuery(`{topics{topics{title,fancy_title}}}`)
	requestLargeHard = MustparseQuery(`{users{username},topics{topics(visible == true){posters{description}}}}`)
}

//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------
//                                                    PARSE
//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

func TestParseRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Helper()
		ParseSmall()
		ParseMedium()
		ParseLarge()
	})
}

func ParseSmall() {
	var p Parser
	v, err := p.Parse(smallFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	_ = v
}

func ParseMedium() {
	var p Parser
	v, err := p.Parse(mediumFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	_ = v
}

func ParseLarge() {
	var p Parser
	v, err := p.Parse(largeFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	_ = v
}

//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------
//                                                    KEEP
//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

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
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepSmallMedium() {

	_, err := smallFixtureValue.Keep(*requestSmallMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepSmallHard() {
	_, err := smallFixtureValue.Keep(*requestSmallHard)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

//
// MEDIUM DATASET
//

func KeepMediumSimple() {
	_, err := mediumFixtureValue.Keep(*requestMediumSimple)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepMediumMedium() {
	_, err := mediumFixtureValue.Keep(*requestMediumMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepMediumHard() {
	_, err := mediumFixtureValue.Keep(*requestMediumHard)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

//
// LARGE DATASET
//

func KeepLargeSimple() {
	_, err := largeFixtureValue.Keep(*requestLargeSimple)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepLargeMedium() {
	_, err := largeFixtureValue.Keep(*requestLargeMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepLargeHard() {
	_, err := largeFixtureValue.Keep(*requestLargeHard)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------
//                                                    CHECK
//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

func TestCheckRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		TestCheckSmallRequestSuccess(t)
		TestCheckMediumRequestSuccess(t)
		TestCheckLargeRequestSuccess(t)
	})

	t.Run("error", func(t *testing.T) {
		// testParseRawNumberError(t, "xyz", "xyz")
	})
}

func TestCheckSmallRequestSuccess(t *testing.T) {
	t.Helper()
	CheckSmallSimple()
	CheckSmallMedium()
	CheckSmallHard()
}

func TestCheckMediumRequestSuccess(t *testing.T) {
	t.Helper()
	CheckMediumSimple()
	CheckMediumMedium()
	CheckMediumHard()
}
func TestCheckLargeRequestSuccess(t *testing.T) {
	t.Helper()
	CheckLargeSimple()
	CheckLargeMedium()
	CheckLargeHard()
}

//
// SMALL DATASET
//

func CheckSmallSimple() {

	err := smallFixtureValue.Check(*requestSmallSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckSmallMedium() {

	err := smallFixtureValue.Check(*requestSmallMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckSmallHard() {
	err := smallFixtureValue.Check(*requestSmallHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

//
// MEDIUM DATASET
//

func CheckMediumSimple() {
	err := mediumFixtureValue.Check(*requestMediumSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckMediumMedium() {
	err := mediumFixtureValue.Check(*requestMediumMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckMediumHard() {
	err := mediumFixtureValue.Check(*requestMediumHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

//
// LARGE DATASET
//

func CheckLargeSimple() {
	err := largeFixtureValue.Check(*requestLargeSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckLargeMedium() {
	err := largeFixtureValue.Check(*requestLargeMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckLargeHard() {
	err := largeFixtureValue.Check(*requestLargeHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}
