# 2016-11-Challenge
Challenge for November, 2016

## Objective

Use struct tags to build a form dynamically. Use Go templates to display the form.

## Example data

```
type UserData struct {
	FullName string `name:"fullname" desc:"Full name"`
	Email    string `name:"email" desc:"Email address"`
	Phone    string `name:"phone" desc:"Phone number"`
}
```

## Challenge 1

Build the form.  Suggested steps:

1. Read the struct tags using the `reflect` package
2. Use the struct tag data to build a form of user data
3. Add more user data fields and/or tags that affect behavior (ex: field is required)

Resources:

* [Go reflect package](https://golang.org/pkg/reflect/)
* [Go template package](https://golang.org/pkg/text/template/)
* [Sample code](https://gist.github.com/sosedoff/b373623a9572cf1a992486d2d87dcd85)

## Challenge 2

Validate the form data using struct tags.

Suggested steps:

1. Add `valid` struct tag to the `UserData` struct
2. Handle form POST of user data
3. Validate submitted data and display result

Resources:

* [govalidator](https://github.com/asaskevich/govalidator)

## Challenge 3

Maintain list of users submitted and provide a way to display user data.
