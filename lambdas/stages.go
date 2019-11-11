package lambdas

/*
	TODO
	Break out params into helper methods.
	- Since helpers are typed we can still validate we have all info we need
	Turn Stage.flags into an array of some new helper struct/interface?
	Add more helper methods (signing etc)
	Can even the result be a helper, (like Json(entity)) to enable fancy decoding?
*/

// Post - create a http post stage, binding the response.body to @param result.
func Post(url string, result interface{}) *Stage {
	return &Stage{"post", []string{url}, result}
}

// Put - craete a http put stage, binding the response.body to @param result.
func Put(url string, result interface{}) *Stage {
	return &Stage{"put", []string{url}, result}
}

// RPC - create a rpc stage, binding the response.body to @param result.
func RPC(topic string, result interface{}) *Stage {
	return &Stage{"rpc", []string{topic}, result}
}
