package route

var (
	CtxKeyProxyName    = contextKey("proxy name")
	CtxKeyProviderName = contextKey("provider name")
	CtxKeyProxy        = contextKey("proxy")
)

type contextKey string

func (c contextKey) String() string {
	return "clash context key " + string(c)
}
