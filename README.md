# abFeatures [![Build Status](https://travis-ci.org/yaronsumel/abFeatures.svg?branch=master)](https://travis-ci.org/yaronsumel/abFeatures)

abFeatures allows you to create and check features in your web projects.

Written by [Yaron Sumel](http://sumel.me).

## Install

    $ go get github.com/yaronsumel/abFeatures

## Usage

```go

import "github.com/yaronsumel/abFeatures"
const NEW_HOME_PAGE_FEATURE = 1

	//Set a new Test Feature
	abFeatures.SetNewFeature(&abFeatures.TestFeature{
	// feature id
		TestId:NEW_HOME_PAGE_FEATURE,
	// chance is a number to rand between
		Chance:100,
	// max-age of each test Cookie
		CookieMaxAge:100,
    	// is the feature active
		Active:true,
	})
```


## Check a Feature

```go
	if abFeatures.HasFeature(NEW_HOME_PAGE_FEATURE,&http.ResponseWriter,&http.Request){
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
