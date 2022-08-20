# Posts Api

In this API you can create, get all posts or by id, update or delete posts.

### For start using this API need

```
make build && make run
```
### After building docker-compose you need to make a migration

```
make migration
```
if you've migrated, you don't need to do this.

#### Server starten on localhost:8000
This API has the following routes:
#### <p>Create our post in body we send json:</p>

>* [POST] /posts
```
{
  "title":"Your title",
  "content":"Your content ps.Booster awesome!"
}
```
#### <p>Get all posts</p>

>* [GET] /posts

 Response

```
"data":{
          { 
            "id":1,
            "title":"Your title",
            "content":"Your content ps.Booster awesome!",
            "create_at": 2022-20-08T11:11:11
          },
          { 
            "id":2,
            "title":"Your title",
            "content":"Your content ps.Booster awesome!",
            "create_at": 2022-20-08T11:11:12
          }
        }
```

#### <p>Get all posts</p>

>* [GET] /posts/{id}

 Response
```
{ 
  "id":1,
  "title":"Your title",
  "content":"Your content ps.Booster awesome!",
  "create_at": 2022-20-08T11:11:11
}
```

#### <p>Update our post in body we send json:</p>

>* [PUT] /posts/{id}

```
{
  "title":"Your title",
  "content":"Your content ps.Booster awesome!"
}
```
Response
```
{
    "status": "deleted"
}
```

#### <p>Delete our post by in <b>url</b>:</p>

>* [DELETE] /posts/{id}

Response
```
{
    "status": "deleted"
}
```













