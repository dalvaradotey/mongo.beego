# Mongo session connection with Beego

All credits for **Kyaw Myint Thein** in [Building simple blog API with Golang + Mongodb](https://medium.com/@kyawmyintthein/building-simple-blog-api-with-golang-mongodb-part-1-d9de449c1fd6#.feazcum6r). I make a small change in connection string for my own purpose. **THANKS!!!!**

First, you need get mongo library for go. We use **mgo** (https://labix.org/mgo), type this in your terminal

```sh
    go get gopkg.in/mgo.v2
```

Right! now add theses files in your project, you need change the connection data in ```conf/mongo.conf```. **Don't forget** include this config file in your ```conf/app.conf``` adding this line:
```sh
    include "mongo.conf"
```
Later in your collection file (you can make collections folder or add file in model folder) add a make session function for collection like this:
```sh
    func newMyCollectionSession() *db.Collection {
        return db.NewCollectionSession("my_collection")
    }
```
Don't forget import mongo utilities in your package
```sh
    import "your/app/path/utilities/mongo"
```
Finally you can use this code something like that

```sh
    conn := newMyCollectionSession()
    defer conn.Close()
    
    err := conn.Session.Find(nil).All(&my_collection)
```
You can check the API documentation for **mgo** [here](https://godoc.org/gopkg.in/mgo.v2).
