package abFeatures

import (
	"net/http"
	"errors"
	"time"
	"math/rand"
	"strconv"
	"fmt"
)

var features = make(map[int]*TestFeature)
var version = "_V1"
var param_name = "tid"
var cookie_name = "T_"

type TestFeature struct {
	// feature id
	TestId 		int
	// chance is a number to rand between
	Chance		int
	// unix timestamp of test expiration
	//ExpireAt	int64
	// max-age of each test Cookie
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now
	// MaxAge>0 means Max-Age attribute present and given in seconds
	CookieMaxAge		int
	// set weather feature is active or not
	Active		bool
}

type abTest struct {}

// set new test feature
// if expireAt is zero , its means it will never expire
func SetNewFeature(feature *TestFeature){
	features[feature.TestId] = feature
}

// set a new version
func SetVersion(v int){
	version = "_^V"+strconv.Itoa(v)
}

// get current version
func getVersion()string{
	return version
}

// change the default param name to custom name
func SetParamName(p string){
	param_name = p
}

// get current version
func getParamName()string{
	return param_name
}

func SetCookieName(name string){
	cookie_name = name
}

func getCookieName()string{
	return cookie_name
}

// main flow for checking if user can access requested feature
func HasFeature(testId int,w *http.ResponseWriter,r *http.Request)bool{

	// get feature by test and check if its still active
	feature,err := getFeatureByTestId(testId)

	if err != nil || !feature.isValid(){
		return false
	}

	// check if user got an active test cookie
	if feature.hasTestCookie(r){
		return true
	}

	// check if user got an test id url param
	if !feature.hasTestParam(r) && !feature.isLuckyUser(){
		return false
	}

	// i guess your user is lucky
	// lets give him a Cookie
	return feature.setTestCookie(w)
}

//
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

//
func (f *TestFeature)setTestCookie(w *http.ResponseWriter)bool{
	http.SetCookie(*w,&http.Cookie{
		Name:f.getCookieName(),
		Value:"TRUE",
		MaxAge:f.CookieMaxAge,
		Secure:false,
		HttpOnly:false,
	})
	return true
}

//
func (f *TestFeature)isLuckyUser()bool{
	if f.Chance <= 0{
		return true
	}
	return random(0,f.Chance) == 0
}


//
func (f *TestFeature)hasTestParam(r *http.Request)bool{
	return f.getTestIdUrlParam(r) == strconv.Itoa(f.TestId)
}

//
func (f *TestFeature)getCookieName()string{
	return fmt.Sprintf("%s_%d_%s",getCookieName(),f.TestId,getVersion())
}

//
func (f *TestFeature)hasTestCookie(r *http.Request) bool{
	cookie,err := r.Cookie(f.getCookieName())
	if err != nil{
		return false
	}
	return cookie.Value == strconv.Itoa(f.TestId)
}

//
func (f *TestFeature)getTestIdUrlParam(r *http.Request)string{
	return r.URL.Query().Get(getParamName())
}

//
func (f *TestFeature)isValid()bool{
	return f.Active /*&& f.ExpireAt > time.Now().Unix()*/
}

//
func getFeatureByTestId(testId int)(*TestFeature,error){
	if feature, ok := features[testId]; ok {
		return feature,nil
	}
	return nil,errors.New("testId cannot be found")
}
