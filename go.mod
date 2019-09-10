module go_admin

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/dchest/captcha v0.0.0-20120311190953-b80c582b3330
	github.com/go-sql-driver/mysql v1.4.1
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace (
	golang.org/x/build => github.com/golang/build v0.0.0-20190416225751-b5f252a0a7dd
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190411191339-88737f569e3a
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190413192849-7f338f571082
	golang.org/x/image => github.com/golang/image v0.0.0-20190417020941-4e30a6eb7d9a
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190415191353-3e0bab5405d6
	golang.org/x/net => github.com/golang/net v0.0.0-20190415214537-1da14a5a36f2
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190312170614-0655857e383f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190416152802-12500544f89f
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190417005754-4ca4b55e2050
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190410155217-1f06c39b4373
)
