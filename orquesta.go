package orquesta

// Version is the version of the SDK.
const SDKVersion = "1.0.0"

// userAgent is the User-Agent of outgoing HTTP requests.
const userAgent = "@orquestadev/go@" + SDKVersion

// Init initializes the SDK with options
func Init(options ClientOptions) (*Client, error) {
	return NewClient(options)
}
