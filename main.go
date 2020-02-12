package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
    app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

    // static source : css img js
    app.HandleDir("/css", "./static/css")
    app.HandleDir("/images", "./static/img")
	app.HandleDir("/js", "./static/js")
	app.HandleDir("/video", "./static/video")
	app.HandleDir("/brain3", "./brian3")
	app.HandleDir("/brain", "./brain")

    // static html
    tmpls := iris.HTML("./templates", ".html")
    app.RegisterView(tmpls)

    // file source
    files := "./static/files/"

    // server
    app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/1.html", func(ctx iris.Context) {
		ctx.View("1.html")
	})

	app.Get("/index.html", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/index2.html", func(ctx iris.Context) {
		ctx.View("index2.html")
	})

	app.Get("/index3.html", func(ctx iris.Context) {
		ctx.View("index3.html")
	})

	app.Get("/class1.html", func(ctx iris.Context) {
		ctx.View("class1.html")
	})

	app.Get("/class2.html", func(ctx iris.Context) {
		ctx.View("class2.html")
	})

	app.Get("/class3.html", func(ctx iris.Context) {
		ctx.View("class3.html")
	})

	app.Get("/brain/brain_index.html", func(ctx iris.Context) {
		ctx.View("brain_index.html")
	})

	app.Get("/brain/brain.html", func(ctx iris.Context) {
		ctx.View("brain.html")
	})

	app.Get("/code.html", func(ctx iris.Context) {
		//a, _ := ctx.URLParamInt("id")
		////fmt.Print("pa=",a)
		//ctx.ViewData("?",a)
		ctx.View("code.html")
	})

	//app.Get("/index3.html", func(ctx iris.Context) {
	//	a, _ := ctx.URLParamInt("id")
	//	fmt.Print("pa=",a)
	//	ctx.View("index3.html")
	//})files/source.zip

    // file download
    app.Get("/files/experiment.zip", func(ctx iris.Context) {
		file := files + "experiment.zip"
		ctx.SendFile(file,"experiment.zip")
	})

	app.Get("/files/class2.xlsx", func(ctx iris.Context) {
		file := files + "class2.xlsx"
		ctx.SendFile(file,"class2.xlsx")
	})

	app.Get("/files/class3.xlsx", func(ctx iris.Context) {
		file := files + "class3.xlsx"
		ctx.SendFile(file,"class3.xlsx")
	})


    // class3 file download
    app.Get("/files/source.zip", func(ctx iris.Context) {
		file := files + "source.zip"
		ctx.SendFile(file,"source.zip")
	})

	app.Get("/files/EEGprocess.zip", func(ctx iris.Context) {
		file := files + "EEGprocess.zip"
		ctx.SendFile(file,"EEGprocess.zip")
	})

	app.Get("/files/fNIRSprocess.zip", func(ctx iris.Context) {
		file := files + "fNIRSprocess.zip"
		ctx.SendFile(file,"fNIRSprocess.zip")
	})

    // class3 show the code
	app.Get("files/source/plot/plot_direction_avg.m", func(ctx iris.Context) {
		file := files + "/source/plot/" + "plot_direction_avg.m"
		ctx.SendFile(file, "plot_direction_avg.m")
	})

	app.Get("files/source/plot/plot_direction_5avg.m", func(ctx iris.Context) {
		file := files + "/source/plot/" + "plot_direction_5avg.m"
		ctx.SendFile(file, "plot_direction_5avg.m")
	})

	app.Get("files/source/plot/plot_direction.m", func(ctx iris.Context) {
		file := files + "/source/plot/" + "plot_direction.m"
		ctx.SendFile(file, "plot_direction.m")
	})

	app.Get("files/source/processing/average_psd_max_channels.m", func(ctx iris.Context) {
		file := files + "/source/processing/" + "average_psd_max_channels.m"
		ctx.SendFile(file, "average_psd_max_channels.m")
	})

	app.Get("files/source/processing/psd_5avg_max_channels.m", func(ctx iris.Context) {
		file := files + "/source/processing/" + "psd_5avg_max_channels.m"
		ctx.SendFile(file, "psd_5avg_max_channels.m")
	})

	app.Get("files/source/processing/psd_max_channels.m", func(ctx iris.Context) {
		file := files + "/source/processing/" + "psd_max_channels.m"
		ctx.SendFile(file, "psd_max_channels.m")
	})



    //app.Run(iris.Addr(":8080"),iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
    app.Run(iris.Addr(":80"), iris.WithCharset("UTF-8"))
}
