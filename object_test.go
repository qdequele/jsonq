package jsonq

import (
	"log"
	"testing"
)

var (
	smallFixtureValue  *Value
	mediumFixtureValue *Value
	largeFixtureValue  *Value

	requestSmallSimple  *Query
	requestSmallMedium  *Query
	requestSmallHard    *Query
	requestMediumSimple *Query
	requestMediumMedium *Query
	requestMediumHard   *Query
	requestLargeSimple  *Query
	requestLargeMedium  *Query
	requestLargeHard    *Query
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

	requestSmallSimple = MustParseQuery(`(st>=1){st}`)
	requestSmallMedium = MustParseQuery(`(gr == 0){st,sid,tt}`)
	requestSmallHard = MustParseQuery(`{st,sid,tt,users(name == "Leonid"){name}}`)
	requestMediumSimple = MustParseQuery(`{person{name{fullName}}}`)
	requestMediumMedium = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio}}`)
	requestMediumHard = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio},users(id > 20){username}}`)
	requestLargeSimple = MustParseQuery(`{users{username}}`)
	requestLargeMedium = MustParseQuery(`{topics{topics{title,fancy_title}}}`)
	requestLargeHard = MustParseQuery(`{users{username},topics{topics(visible == true){posters{description}}}}`)
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
		TestKeepRequestOnly(t)
		TestKeepRequestParseCMD(t)
		TestKeepRequestParseAll(t)
	})

	t.Run("error", func(t *testing.T) {
		// testParseRawNumberError(t, "xyz", "xyz")
	})
}

func TestKeepRequestOnly(t *testing.T) {
	t.Helper()
	KeepSmallSimple(false, false)
	KeepSmallMedium(false, false)
	KeepSmallHard(false, false)
	KeepMediumSimple(false, false)
	KeepMediumMedium(false, false)
	KeepMediumHard(false, false)
	KeepLargeSimple(false, false)
	KeepLargeMedium(false, false)
	KeepLargeHard(false, false)
}

func TestKeepRequestParseCMD(t *testing.T) {
	t.Helper()
	KeepSmallSimple(false, true)
	KeepSmallMedium(false, true)
	KeepSmallHard(false, true)
	KeepMediumSimple(false, true)
	KeepMediumMedium(false, true)
	KeepMediumHard(false, true)
	KeepLargeSimple(false, true)
	KeepLargeMedium(false, true)
	KeepLargeHard(false, true)
}

func TestKeepRequestParseAll(t *testing.T) {
	t.Helper()
	KeepSmallSimple(true, true)
	KeepSmallMedium(true, true)
	KeepSmallHard(true, true)
	KeepMediumSimple(true, true)
	KeepMediumMedium(true, true)
	KeepMediumHard(true, true)
	KeepLargeSimple(true, true)
	KeepLargeMedium(true, true)
	KeepLargeHard(true, true)
}

//
// SMALL DATASET
//

func KeepSmallSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallSimple = MustParseQuery(`(st>=1){st}`)
	}
	_, err := smallFixtureValue.Keep(*requestSmallSimple)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepSmallMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallMedium = MustParseQuery(`(gr == 0){st,sid,tt}`)
	}
	_, err := smallFixtureValue.Keep(*requestSmallMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepSmallHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallHard = MustParseQuery(`{st,sid,tt,users(name == "Leonid"){name}}`)
	}
	_, err := smallFixtureValue.Keep(*requestSmallHard)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

//
// MEDIUM DATASET
//

func KeepMediumSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumSimple = MustParseQuery(`{person{name{fullName}}}`)
	}
	_, err := mediumFixtureValue.Keep(*requestMediumSimple)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepMediumMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumMedium = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio}}`)
	}
	_, err := mediumFixtureValue.Keep(*requestMediumMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepMediumHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumHard = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio},users(id > 20){username}}`)
	}
	_, err := mediumFixtureValue.Keep(*requestMediumHard)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

//
// LARGE DATASET
//

func KeepLargeSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeSimple = MustParseQuery(`{users{username}}`)
	}
	_, err := largeFixtureValue.Keep(*requestLargeSimple)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepLargeMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeMedium = MustParseQuery(`{topics{topics{title,fancy_title}}}`)
	}
	_, err := largeFixtureValue.Keep(*requestLargeMedium)
	if err != nil {
		log.Fatalf("cannot keep json: %s", err)
	}
}

func KeepLargeHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeHard = MustParseQuery(`{users{username},topics{topics(visible == true){posters{description}}}}`)
	}
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
		TestCheckRequestOnly(t)
		TestCheckRequestParseCMD(t)
		TestCheckRequestParseAll(t)
	})

	t.Run("error", func(t *testing.T) {
		// testParseRawNumberError(t, "xyz", "xyz")
	})
}

func TestCheckRequestOnly(t *testing.T) {
	t.Helper()
	CheckSmallSimple(false, false)
	CheckSmallMedium(false, false)
	CheckSmallHard(false, false)
	CheckMediumSimple(false, false)
	CheckMediumMedium(false, false)
	CheckMediumHard(false, false)
	CheckLargeSimple(false, false)
	CheckLargeMedium(false, false)
	CheckLargeHard(false, false)
}

func TestCheckRequestParseCMD(t *testing.T) {
	t.Helper()
	CheckSmallSimple(false, true)
	CheckSmallMedium(false, true)
	CheckSmallHard(false, true)
	CheckMediumSimple(false, true)
	CheckMediumMedium(false, true)
	CheckMediumHard(false, true)
	CheckLargeSimple(false, true)
	CheckLargeMedium(false, true)
	CheckLargeHard(false, true)
}

func TestCheckRequestParseAll(t *testing.T) {
	t.Helper()
	CheckSmallSimple(true, true)
	CheckSmallMedium(true, true)
	CheckSmallHard(true, true)
	CheckMediumSimple(true, true)
	CheckMediumMedium(true, true)
	CheckMediumHard(true, true)
	CheckLargeSimple(true, true)
	CheckLargeMedium(true, true)
	CheckLargeHard(true, true)
}

//
// SMALL DATASET
//

func CheckSmallSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallSimple = MustParseQuery(`(st>=1){st}`)
	}
	err := smallFixtureValue.Check(*requestSmallSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckSmallMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallMedium = MustParseQuery(`(gr == 0){st,sid,tt}`)
	}
	err := smallFixtureValue.Check(*requestSmallMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckSmallHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseSmall()
	}
	if parseCMD {
		requestSmallHard = MustParseQuery(`{st,sid,tt,users(name == "Leonid"){name}}`)
	}
	err := smallFixtureValue.Check(*requestSmallHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

//
// MEDIUM DATASET
//

func CheckMediumSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumSimple = MustParseQuery(`{person{name{fullName}}}`)
	}
	err := mediumFixtureValue.Check(*requestMediumSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckMediumMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumMedium = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio}}`)
	}
	err := mediumFixtureValue.Check(*requestMediumMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckMediumHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseMedium()
	}
	if parseCMD {
		requestMediumHard = MustParseQuery(`{person{name{fullName},email,geo{city,state},bio},users(id > 20){username}}`)
	}
	err := mediumFixtureValue.Check(*requestMediumHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

//
// LARGE DATASET
//

func CheckLargeSimple(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeSimple = MustParseQuery(`{users{username}}`)
	}
	err := largeFixtureValue.Check(*requestLargeSimple)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckLargeMedium(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeMedium = MustParseQuery(`{topics{topics{title,fancy_title}}}`)
	}
	err := largeFixtureValue.Check(*requestLargeMedium)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}

func CheckLargeHard(parseJSON, parseCMD bool) {

	if parseJSON {
		ParseLarge()
	}
	if parseCMD {
		requestLargeHard = MustParseQuery(`{users{username},topics{topics(visible == true){posters{description}}}}`)
	}
	err := largeFixtureValue.Check(*requestLargeHard)
	if err != nil {
		log.Fatalf("cannot Check json: %s", err)
	}
}
