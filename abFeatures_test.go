package abFeatures

import (
	"strconv"
	"testing"
	"net/http"
	"net/url"
)

func TestSetNewFeature(t *testing.T){

	features[1] = &TestFeature{}

	SetNewFeature(&TestFeature{
		TestId:1,
		Chance:100,
		CookieMaxAge:100,
		Active:true,
	})

	if _, ok := features[1]; !ok {
		t.Fatalf("Could not set new Feature")
	}
}

func TestSetVersion(t *testing.T){
	SetVersion(1)
	if version != version_prefix+strconv.Itoa(1){
		t.Fatalf("Could not set new Version")
	}
}

func TestGetVersion(t *testing.T){
	SetVersion(1)
	if getVersion() != version{
		t.Fatalf("Could not get new Version")
	}
}

func TestSetParamName(t *testing.T){
	SetParamName("ABC")
	if param_name != "ABC"{
		t.Fatalf("Could not Set Param Name")
	}
}

func TestGetParamName(t *testing.T){
	SetParamName("ABC")
	if getParamName() != "ABC"{
		t.Fatalf("Could not Get Param Name")
	}
}

func TestSetCookieName(t *testing.T){
	SetCookieName("ABC")
	if cookie_name != "ABC"{
		t.Fatalf("Could not Set Param Name")
	}
}

func TestGetCookieName(t *testing.T){
	SetCookieName("ABC")
	if cookie_name != getCookieName(){
		t.Fatalf("Could not Get Param Name")
	}
}

func TestHasFeature(t *testing.T){
}


//
func TestSetTestCookie(t *testing.T){
}



//
func TestIsLuckyUser(t *testing.T){

	tf := TestFeature{
		TestId:1,
		Chance:1,
		CookieMaxAge:100,
		Active:true,
	}

	if tf.isLuckyUser() == false{
		t.Fatalf("somthing is wrong with isLuckyUser")
	}

}


//
func TestHasTestParam(t *testing.T){

	SetParamName("tid")

	tf := TestFeature{
		TestId:1,
		Chance:1,
		CookieMaxAge:100,
		Active:true,
	}

	r1 := &http.Request{
		URL:&url.URL{RawQuery:"tid=1"},
	}

	r2 := &http.Request{
		URL:&url.URL{RawQuery:"tid=2"},
	}

	if !tf.hasTestParam(r1) || tf.hasTestParam(r2){
		t.Fatalf("has no tets param")
	}

}


//
func TestHasTestCookie(t *testing.T) {

	SetCookieName("ABC")
	SetVersion(1)

	//ABC_1__V1

	tf := TestFeature{
		TestId:1,
		Chance:1,
		CookieMaxAge:100,
		Active:true,
	}

	//should pass
	r1 := &http.Request{
		Header:http.Header{"Cookie": {tf.getCookieName()+"=1"}},
	}

	//should fail
	r2 := &http.Request{
		Header:http.Header{"Cookie": {"ABC=1"}},
	}

	if !tf.hasTestCookie(r1) || tf.hasTestCookie(r2){
		t.Fatalf("has no tets cookie")
	}

}

//
func TestGetTestIdUrlParam(t *testing.T){

	SetParamName("tid")

	tf := TestFeature{
		TestId:1,
		Chance:1,
		CookieMaxAge:100,
		Active:true,
	}

	r1 := &http.Request{
		URL:&url.URL{RawQuery:"tid=1"},
	}

	if tf.getTestIdUrlParam(r1) != "1"{
		t.Fatalf("getTestIdUrlParam")
	}

}

//
func TestIsValid(t *testing.T){

	tf := TestFeature{
		Active:true,
	}

	if tf.Active != tf.isValid(){
		t.Fatalf("is valid is wrong")
	}

}

//
func TestGetFeatureByTestId(t *testing.T){

	features[1] = &TestFeature{
		TestId:1,
	}

	f,err := getFeatureByTestId(1)

	if err != nil || f.TestId != 1{
		t.Fatalf("TestGetFeatureByTestId")
	}


}
