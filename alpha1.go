  package main

  import (
     "github.com/rootsdev/gofamilysearch"
     "log"
     "net/http"
  )

  func main() {
      // Context can be shared among go-routines
     ctx := &gofamilysearch.Context{
        Environment: "sandbox",
     }
     // Client is specific to a user
     c := &gofamilysearch.Client{
        Context: ctx,
        AccessToken: "access token for the requesting user goes here",
        // if running on app engine, pass in &urlfetch.Transport{Context: appengine.NewContext(request)}
        Transport: DefaultTransport,
     }

     user, err := c.GetCurrentUser()
     if err != nil {
        log.Panic(err)
     }
     log.Printf("ID=%s personID=%s treeUserID=%s\n", user.ID, user.PersonID, user.TreeUserID)
  }  
  
