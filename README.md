# **MoGo**
#### **Mongo DB wrapper for golang**
###### Quick Start

### Install

```bash
go get github.com/risusanto/mogo
```

### Usage
To get started, import `mogo` package, and setup for default config:
```go

import "github.com/risusanto/mogo"

func init() {
   // Setup mgm default config
   var mongoConfig  corey.DBConfig
   mongoConfig.MongoURI = "mongodb://root:12345@localhost:27017"
   mongoConfig.DBName = "mogo_db"
   _,_,err := corey.NewConnection(mongoConfig)
   if err != nil{
     panic(err)
   }
}
```

Define model:
```go
type Podcast struct {
    corey.BaseModel		  	  `bson:",inline"`
    Title  string             `json:"title" bson:"title,omitempty"`
    Author string             `json:"author" bson:"author,omitempty"`
}
```

Insert data:
```go

podcast := &Podcast{}

podcast.Tittle = "Learn GoLang"
podcast.Author = "Ari"

podcast.Create(podcast)

```

