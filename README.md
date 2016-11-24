Go DeployGate
=========

[![wercker status](https://app.wercker.com/status/34cfff3fe4ffc0a83bc518aa60d7a7b1/m/master "wercker status")](https://app.wercker.com/project/byKey/34cfff3fe4ffc0a83bc518aa60d7a7b1)

Go DeployGate is a Golang API client for interacting with most facets of the
[DeployGate API](https://docs.deploygate.com/api).

Installation
------------
This is a client library, so there is nothing to install.

Usage
-----
Download the library into your `$GOPATH`:

    $ go get github.com/recruit-mp/go-deploygate

Import the library into your tool:

```go
import "github.com/recruit-mp/go-deploygate"
```

Examples
--------
DeployGate's API is designed to work in the following manner:

1. Create (or clone) a new configuration version for the service
2. Make any changes to the version
3. Validate the version
4. Activate the version

This flow using the Golang client looks like this:

```go
// Create a client object. The client has no state, so it can be persisted
// and re-used. It is also safe to use concurrently due to its lack of state.
// There is also a DefaultClient() method that reads an environment variable.
// Please see the documentation for more information and details.
client, err := deploygate.NewClient("YOUR_DEPLOYGATE_API_KEY")

// Get users who collaborate with your App
g := &deploygate.GetAppCollaboratorInput{
	Owner:    "owner-name",
	Platform: "ios", // ios or android
	AppId:    "your.app.id",
}

// Add collaborator to your App
collaborator, err := client.GetAppCollaborator(g)

a := &deploygate.AddAppCollaboratorInput{
	Owner:    "owner-name",
	Platform: "ios",
	AppId:    "your.app.id",
	Users:    "username",
	Role:     2, //`1`(for developer) or `2`(for tester)
}

addResponse, err2 := client.AddAppCollaborator(a)

// Delete collaborator from your App
d := &deploygate.DeleteAppCollaboratorInput{
	Owner:    "owner-name",
	Platform: "ios",
	AppId:    "your.app.id",
	Users:    "username",
	Role:     2, //`1`(for developer) or `2`(for tester)
}

deleteResponse, err3 := client.DeleteAppCollaborator(d)
```

More information can be found in the
[DeployGate Godoc](https://godoc.org/github.com/recruit-mp/go-deploygate).

Reference
----------

- [terraform-provider-deploygate](https://github.com/recruit-mp/terraform-provider-deploygate)

TODO
----

- Support following [DeployGate API](https://docs.deploygate.com/reference#deploygate-api):
  - [ ]: [App Upload API](https://docs.deploygate.com/reference#upload)

License
-------

- Copyright 2016 Naoki Ainoya, Recruit Marketing Partners Co., Ltd., Apache License, Version 2.0
- This base client implementation is copied from [sethvargo/go-fastly](https://github.com/sethvargo/go-fastly) which is licensed as Copyright 2015 Seth Vargo, Apache Lisense, Version 2.0
