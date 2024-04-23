package access_log

import (
	"bufio"
	"time"

	"go-learn/iris/bootstrap"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/sessions"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func New(b *bootstrap.Bootstrapper) *accesslog.AccessLog {
	return makeAccessLog()
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	ac := New(b)
	//b.UseGlobal(h)
	b.UseRouter(ac.Handler)
	broker := ac.Broker()
	b.Get("/logs", accesslog.SkipHandler, accessLogBrokerHandler(broker))
}

// Default line format:
// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
//
// Read the example and its comments carefully.
func makeAccessLog() *accesslog.AccessLog {
	//pathToAccessLog := "./runtimes/logs/access_log.%Y-%m-%d-%H"
	pathToAccessLog := "./runtimes/logs/%Y-%m/access_log.%Y-%m-%d.log"
	w, err := rotatelogs.New(
		pathToAccessLog,
		rotatelogs.WithMaxAge(24*time.Hour*7),     //最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour)) //文件的轮转时间间隔
	if err != nil {
		panic(err)
	}
	//方便看效果自定一下长度10
	//ac := accesslog.New(bufio.NewWriter(w))
	ac := accesslog.New(bufio.NewWriterSize(w, 10))
	ac.Delim = ' '
	ac.ResponseBody = false
	//json Formatter
	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})
	//ac.AddOutput(io.MultiWriter(w, os.Stdout))
	ac.AddFields(func(ctx iris.Context, fields *accesslog.Fields) {
		reqId := ctx.GetID()
		fields.Set("request-id", reqId)
	})
	ac.AddFields(func(ctx iris.Context, fields *accesslog.Fields) {
		if sess := sessions.Get(ctx); sess != nil {
			fields.Set("session_id", sess.ID())

			sess.Visit(func(k string, v interface{}) {
				fields.Set(k, v)
			})
		}
	})
	// Add a custom field of "auth" when basic auth is available.
	ac.AddFields(func(ctx iris.Context, fields *accesslog.Fields) {
		if username, password, ok := ctx.Request().BasicAuth(); ok {
			fields.Set("auth", username+":"+password)
		}
	})

	return ac
}

func accessLogBrokerHandler(b *accesslog.Broker) iris.Handler {
	return func(ctx iris.Context) {
		// accesslog.Skip(ctx) // or inline skip.
		logs := b.NewListener() // <- IMPORTANT

		ctx.Header("Transfer-Encoding", "chunked")
		notifyClose := ctx.Request().Context().Done()
		for {
			select {
			case <-notifyClose:
				b.CloseListener(logs) // <- IMPORTANT

				err := ctx.Request().Context().Err()
				ctx.Application().Logger().Infof("Listener closed [%v], loop end.", err)
				return
			case log := <-logs: // <- IMPORTANT
				_ = ctx.JSON(log, iris.JSON{Indent: "  ", UnescapeHTML: true})
				ctx.ResponseWriter().Flush()
			}
		}
	}
}
