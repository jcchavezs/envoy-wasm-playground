package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &testContext{contextID: contextID}
}

const tickMilliseconds uint32 = 100

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
		return types.OnPluginStartStatusFailed
	}
	proxywasm.LogInfof("set tick period milliseconds: %d", tickMilliseconds)

	return types.OnPluginStartStatusOK
}

func (ctx *pluginContext) OnTick() {
	// proxywasm.LogDebugf("OnTick")
}

type testContext struct {
	types.DefaultHttpContext
	contextID uint32
}

// Override types.DefaultHttpContext.
func (ctx *testContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
	}

	for _, h := range hs {
		proxywasm.LogDebugf("request header --> %q: %q", h[0], h[1])
	}

	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *testContext) OnHttpRequestBody(bodySize int, endOfStream bool) types.Action {
	if endOfStream {
		proxywasm.LogDebugf("request body at end of stream, body size: %d", bodySize)
		return types.ActionPause
	}

	proxywasm.LogDebugf("about to get request body")
	body, err := proxywasm.GetHttpRequestBody(0, bodySize)
	if err != nil {
		proxywasm.LogCriticalf("failed to get request body: %v", err)
	}

	proxywasm.LogDebugf("request body --> %q", body)

	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *testContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpResponseHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get response headers: %v", err)
	}

	for _, h := range hs {
		proxywasm.LogDebugf("response header <-- %q: %q", h[0], h[1])
	}
	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *testContext) OnHttpResponseBody(bodySize int, endOfStream bool) types.Action {
	if endOfStream {
		proxywasm.LogDebugf("response body end of stream, body size: %d", bodySize)
		return types.ActionPause
	}

	proxywasm.LogDebugf("about to get response body")
	body, err := proxywasm.GetHttpResponseBody(0, bodySize)
	if err != nil {
		proxywasm.LogCriticalf("failed to get response body: %v", err)
	}

	proxywasm.LogDebugf("response body --> %q", body)

	return types.ActionContinue
}
