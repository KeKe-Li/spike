module spike

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.41.0
	golang.org/x/blog => github.com/golang/blog v0.0.0-20190708141629-e28c63452d36
	golang.org/x/build => github.com/golang/build v0.0.0-20190709001953-30c0e6b89ea0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190627132806-fd42eb6b336f
	golang.org/x/image => github.com/golang/image v0.0.0-20190703141733-d6a02ce849c9
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190711165009-e47acb2ca7f9
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190620143337-7c3f2128ad9b
	golang.org/x/review => github.com/golang/review v0.0.0-20190508204355-8102926ea734
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190710143415-6ec70d6a5542
	golang.org/x/talks => github.com/golang/talks v0.0.0-20190313194420-5ca518b26a55
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190711191110-9a621aea19f8
	golang.org/x/tour => github.com/golang/tour v0.0.0-20190611164551-1f1f3d2b450b
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.7.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190708153700-3bdd9d9f5532
	google.golang.org/grpc => github.com/grpc/grpc-go v1.22.0
	gopkg.in/jcmturner/gokrb5.v7 => github.com/jcmturner/gokrb5 v7.3.0+incompatible
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/ugorji/go/codec v0.0.0-20181206144755-e72634d4d386 // indirect
	github.com/unrolled/render v0.0.0-20180914162206-b9786414de4d
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
