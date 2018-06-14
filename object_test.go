package jsonQuerry

import (
	"encoding/json"
	"log"
	"testing"
)

const responseSmallSimple = `{"st":1}`
const responseSmallMedium = `{"sid":486,"st":1,"tt":"\"active\""}`
const responseSmallHard = `{"sid":486,"st":1,"tt":"\"active\"","users":[{"name":"\"Leonid\""},{"name":"\"Bugaev\""}]}`

const responseMediumSimple = `{"person":{"name":{"fullName":"\"Leonid Bugaev\""}}}`
const responseMediumMedium = `{"person":{"bio":"\"Senior engineer at Granify.com\"","email":"\"leonsbox@gmail.com\"","geo":{"city":"\"Saint Petersburg\"","state":"\"Saint Petersburg\""},"name":{"fullName":"\"Leonid Bugaev\""}}}`
const responseMediumHard = `{"person":{"bio":"\"Senior engineer at Granify.com\"","email":"\"leonsbox@gmail.com\"","geo":{"city":"\"Saint Petersburg\"","state":"\"Saint Petersburg\""},"name":{"fullName":"\"Leonid Bugaev\""}},"users":[{"username":"\"system\""},{"username":"\"zergot\""},{"username":"\"sameer\""},{"username":"\"HenryMirror\""},{"username":"\"fimp\""},{"username":"\"agilliland\""},{"username":"\"amir\""},{"username":"\"waseem\""},{"username":"\"tovenaar\""}]}`

const responseLargeSimple = `{"users":[{"username":"\"system\""},{"username":"\"zergot\""},{"username":"\"sameer\""},{"username":"\"HenryMirror\""},{"username":"\"fimp\""},{"username":"\"agilliland\""},{"username":"\"amir\""},{"username":"\"waseem\""},{"username":"\"tovenaar\""},{"username":"\"Ben\""},{"username":"\"MarkLaFay\""},{"username":"\"camsaul\""},{"username":"\"mhjb\""},{"username":"\"jbwiv\""},{"username":"\"Maggs\""},{"username":"\"andrefaria\""},{"username":"\"bencarter78\""},{"username":"\"vikram\""},{"username":"\"edchan77\""},{"username":"\"karthikd\""},{"username":"\"arthurz\""},{"username":"\"tom\""},{"username":"\"LeoNogueira\""},{"username":"\"ss06vi\""},{"username":"\"mattcollins\""},{"username":"\"krmmalik\""},{"username":"\"odysseas\""},{"username":"\"jonthewayne\""},{"username":"\"anandiyer\""},{"username":"\"alnorth\""},{"username":"\"j_at_svg\""},{"username":"\"styts\""}]}`
const responseLargeMedium = `{"topics":{"topics":[{"fancy_title":"\"Welcome to Metabase\u0026rsquo;s Discussion Forum\"","title":"\"Welcome to Metabase's Discussion Forum\""},{"fancy_title":"\"Formatting Dates\"","title":"\"Formatting Dates\""},{"fancy_title":"\"Setting for google api key\"","title":"\"Setting for google api key\""},{"fancy_title":"\"Cannot see non-US timezones on the admin\"","title":"\"Cannot see non-US timezones on the admin\""},{"fancy_title":"\"External (Metabase level) linkages in data schema\"","title":"\"External (Metabase level) linkages in data schema\""},{"fancy_title":"\"Query working on \u0026ldquo;Questions\u0026rdquo; but not in \u0026ldquo;Pulses\u0026rdquo;\"","title":"\"Query working on \\\"Questions\\\" but not in \\\"Pulses\\\"\""},{"fancy_title":"\"Pulses posted to Slack don\u0026rsquo;t show question output\"","title":"\"Pulses posted to Slack don't show question output\""},{"fancy_title":"\"Should we build Kafka connecter or Kafka plugin\"","title":"\"Should we build Kafka connecter or Kafka plugin\""},{"fancy_title":"\"Change X and Y on graph\"","title":"\"Change X and Y on graph\""},{"fancy_title":"\"Issues sending mail via office365 relay\"","title":"\"Issues sending mail via office365 relay\""},{"fancy_title":"\"I see triplicates of my mongoDB collections\"","title":"\"I see triplicates of my mongoDB collections\""},{"fancy_title":"\"Google Analytics plugin\"","title":"\"Google Analytics plugin\""},{"fancy_title":"\"With-mongo-connection failed: bad connection details:\"","title":"\"With-mongo-connection failed: bad connection details:\""},{"fancy_title":"\"\u0026ldquo;We couldn\u0026rsquo;t understand your question.\u0026rdquo; when I query mongoDB\"","title":"\"\\\"We couldn't understand your question.\\\" when I query mongoDB\""},{"fancy_title":"\"My bar charts are all thin\"","title":"\"My bar charts are all thin\""},{"fancy_title":"\"What is the expected return order of columns for graphing results when using raw SQL?\"","title":"\"What is the expected return order of columns for graphing results when using raw SQL?\""},{"fancy_title":"\"Set site url from admin panel\"","title":"\"Set site url from admin panel\""},{"fancy_title":"\"Internationalization (i18n)\"","title":"\"Internationalization (i18n)\""},{"fancy_title":"\"Returning raw data with no filters always returns We couldn\u0026rsquo;t understand your question\"","title":"\"Returning raw data with no filters always returns We couldn't understand your question\""},{"fancy_title":"\"Support for Cassandra?\"","title":"\"Support for Cassandra?\""},{"fancy_title":"\"Mongo query with Date breaks [solved: Mongo 3.0 required]\"","title":"\"Mongo query with Date breaks [solved: Mongo 3.0 required]\""},{"fancy_title":"\"Can this connect to MS SQL Server?\"","title":"\"Can this connect to MS SQL Server?\""},{"fancy_title":"\"Cannot restart metabase in docker\"","title":"\"Cannot restart metabase in docker\""},{"fancy_title":"\"Edit Max Rows Count\"","title":"\"Edit Max Rows Count\""},{"fancy_title":"\"Creating charts by querying more than one table at a time\"","title":"\"Creating charts by querying more than one table at a time\""},{"fancy_title":"\"Trying to add RDS postgresql as the database fails silently\"","title":"\"Trying to add RDS postgresql as the database fails silently\""},{"fancy_title":"\"Deploy to Heroku isn\u0026rsquo;t working\"","title":"\"Deploy to Heroku isn't working\""},{"fancy_title":"\"Can I use DATEPART() in SQL queries?\"","title":"\"Can I use DATEPART() in SQL queries?\""},{"fancy_title":"\"Feature Request: LDAP Authentication\"","title":"\"Feature Request: LDAP Authentication\""},{"fancy_title":"\"Migrating from internal H2 to Postgres\"","title":"\"Migrating from internal H2 to Postgres\""}]}}`
const responseLargeHard = `{"topics":{"topics":[{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster, Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""},{"description":"\"Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster, Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster, Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster, Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster, Frequent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Frequent Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster, Most Recent Poster\""}]},{"posters":[{"description":"\"Original Poster\""},{"description":"\"Most Recent Poster\""}]}]},"users":[{"username":"\"system\""},{"username":"\"zergot\""},{"username":"\"sameer\""},{"username":"\"HenryMirror\""},{"username":"\"fimp\""},{"username":"\"agilliland\""},{"username":"\"amir\""},{"username":"\"waseem\""},{"username":"\"tovenaar\""},{"username":"\"Ben\""},{"username":"\"MarkLaFay\""},{"username":"\"camsaul\""},{"username":"\"mhjb\""},{"username":"\"jbwiv\""},{"username":"\"Maggs\""},{"username":"\"andrefaria\""},{"username":"\"bencarter78\""},{"username":"\"vikram\""},{"username":"\"edchan77\""},{"username":"\"karthikd\""},{"username":"\"arthurz\""},{"username":"\"tom\""},{"username":"\"LeoNogueira\""},{"username":"\"ss06vi\""},{"username":"\"mattcollins\""},{"username":"\"krmmalik\""},{"username":"\"odysseas\""},{"username":"\"jonthewayne\""},{"username":"\"anandiyer\""},{"username":"\"alnorth\""},{"username":"\"j_at_svg\""},{"username":"\"styts\""}]}`

var smallFixtureValue *Value
var mediumFixtureValue *Value
var largeFixtureValue *Value

func init() {
	var p Parser
	v, err := p.Parse(smallFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	smallFixtureValue = v

	v, err = p.Parse(mediumFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	mediumFixtureValue = v

	v, err = p.Parse(largeFixture)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	largeFixtureValue = v
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

	var querry = `{st}`
	response := responseSmallSimple

	request, _ := NewKeepRequest(querry)

	newvalue, err := smallFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepSmallMedium() {

	var querry = `{st,sid,tt}`
	response := responseSmallMedium

	request, _ := NewKeepRequest(querry)

	newvalue, err := smallFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepSmallHard() {

	var querry = `{st,sid,tt,users:{name}}`
	response := responseSmallHard

	request, _ := NewKeepRequest(querry)

	newvalue, err := smallFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

//
// MEDIUM DATASET
//

func KeepMediumSimple() {

	var querry = `{person:{name:{fullName}}}`
	response := responseMediumSimple

	request, _ := NewKeepRequest(querry)

	newvalue, err := mediumFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepMediumMedium() {

	var querry = `{person:{name:{fullName},email,geo:{city,state},bio}}`
	response := responseMediumMedium

	request, _ := NewKeepRequest(querry)

	newvalue, err := mediumFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepMediumHard() {

	var querry = `{person:{name:{fullName},email,geo:{city,state},bio},users:{username}}`
	response := responseMediumHard

	request, _ := NewKeepRequest(querry)

	newvalue, err := mediumFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

//
// LARGE DATASET
//

func KeepLargeSimple() {

	var querry = `{users:{username}}`
	response := responseLargeSimple

	request, _ := NewKeepRequest(querry)

	newvalue, err := largeFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepLargeMedium() {

	var querry = `{topics:{topics:{title,fancy_title}}}`
	response := responseLargeMedium

	request, _ := NewKeepRequest(querry)

	newvalue, err := largeFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}

func KeepLargeHard() {

	var querry = `{users:{username},topics:{topics:{posters:{description}}}}`
	response := responseLargeHard

	request, _ := NewKeepRequest(querry)

	newvalue, err := largeFixtureValue.Keep(request)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	json, err := json.Marshal(newvalue)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	if string(json) != response {
		// log.Fatalf("Wrong response\n >have : %s\n > should have : %s", string(json), response)
	}
}
