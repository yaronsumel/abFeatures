# abFeatures [![Build Status](https://travis-ci.org/yaronsumel/abFeatures.svg?branch=master)](https://travis-ci.org/yaronsumel/abFeatures)

abFeatures allows you to create and check features in your web projects.

Written by[Yaron Sumel](http://sumel.me).

## Install

    $ go get github.com/yaronsumel/abFeatures

## Usage

```go

    	import "github.com/yaronsumel/abFeatures"

	//Set a new Test Feature
	abFeatures.SetNewFeature(&abFeatures.TestFeature{
	// feature id
		TestId:2,
	// chance is a number to rand between
		Chance:100,
	// unix timestamp of test expiration
		ExpireAt:1000,
	// max-age of each test Cookie
		CookieMaxAge:100,
    	// is the feature active
		Active:true,
	})
```


## Check a Feature

```go
	if abFeatures.HasFeature(2,&http.ResponseWriter,&http.Request}){
		// user got the test feature !!
	}
```

## Advanced Usage


```go
    // set a new version in order to ignore old cookies
	abFeatures.SetVersion(3)

    // change cookie name as you wish
	abFeatures.SetCookieName("new_cookie_name")

    // change the GET param to something you like more than tid
	abFeatures.SetParamName("new_test_param")
```


## TBD

    * Unit-Testing
