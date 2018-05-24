package packages

import (
	"net/http"

	"github.com/antchfx/antch"
)

func init() {
	Packages["antch"] = map[string]interface{}{
		"ParseHTML":                antch.ParseHTML,
		"ParseXML":                 antch.ParseXML,
		"NewCrawler":               antch.NewCrawler,
		"Handler":                  Handler,
		"VoidHandler":              antch.VoidHandler,
		"ToHandlerFunc":            ToHandlerFunc,
		"HttpMessageHandler":       HttpMessageHandler,
		"HttpMessageHandlerFunc":   HttpMessageHandlerFunc,
		"ToHttpMessageHandlerFunc": ToHttpMessageHandlerFunc,
		"Item":                  Item,
		"Logger":                Logger,
		"ParseMediaType":        antch.ParseMediaType,
		"Middleware":            Middleware,
		"ToMiddleware":          ToMiddleware,
		"CompressionMiddleware": antch.CompressionMiddleware,
		"CookiesMiddleware":     antch.CookiesMiddleware,
		"ProxyMiddleware":       antch.ProxyMiddleware,
		"RobotstxtMiddleware":   antch.RobotstxtMiddleware,
		"Pipeline":              Pipeline,
		"ToPipeline":            ToPipeline,
		"PipelineHandler":       PipelineHandler,
		"PipelineHandlerFunc":   PipelineHandlerFunc,
		"ToPipelineHandlerFunc": ToPipelineHandlerFunc,
	}
	PackageTypes["antch"] = map[string]interface{}{
		"Crawler":   antch.Crawler{},
		"MediaType": antch.MediaType{},
		"ProxyKey":  antch.ProxyKey{},
	}
}

func Handler() antch.Handler {
	var temp antch.Handler
	return temp
}

func ToHandlerFunc(f func(chan<- antch.Item, *http.Response)) antch.HandlerFunc {
	return antch.HandlerFunc(f)
}

func HttpMessageHandler() antch.HttpMessageHandler {
	var temp antch.HttpMessageHandler
	return temp
}

func HttpMessageHandlerFunc() antch.HttpMessageHandlerFunc {
	var temp antch.HttpMessageHandlerFunc
	return temp
}

func ToHttpMessageHandlerFunc(f func(*http.Request) (*http.Response, error)) antch.HttpMessageHandlerFunc {
	return antch.HttpMessageHandlerFunc(f)
}

func Item() antch.Item {
	var temp antch.Item
	return temp
}

func Logger() antch.Logger {
	var temp antch.Logger
	return temp
}

func Middleware() antch.Middleware {
	var temp antch.Middleware
	return temp
}

func ToMiddleware(f func(antch.HttpMessageHandler) antch.HttpMessageHandler) antch.Middleware {
	return antch.Middleware(f)
}

func Pipeline() antch.Pipeline {
	var temp antch.Pipeline
	return temp
}

func ToPipeline(f func(antch.PipelineHandler) antch.PipelineHandler) antch.Pipeline {
	return antch.Pipeline(f)
}

func PipelineHandler() antch.PipelineHandler {
	var temp antch.PipelineHandler
	return temp
}

func PipelineHandlerFunc() antch.PipelineHandlerFunc {
	var temp antch.PipelineHandlerFunc
	return temp
}

func ToPipelineHandlerFunc(f func(antch.Item)) antch.PipelineHandlerFunc {
	return antch.PipelineHandlerFunc(f)
}
