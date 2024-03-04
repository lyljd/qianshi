package router

type Router struct {
	Path      string // no prefix "/"
	To        string // "http(s)://host:port"
	NeedLogin bool
	Power     int
	Children  []Router
}

var Routers = []Router{
	{
		Path: "user",
		To:   "http://localhost:9007",
		Children: []Router{
			{
				Path: "login",
				Children: []Router{
					{Path: "email"},
					{Path: "pass"},
				},
			},
			{
				Path:      "me",
				NeedLogin: true,
				Children: []Router{
					{Path: "exp"},
					{Path: "info"},
					{Path: "security"},
					{Path: "coin"},
					{
						Path: "pass",
						Children: []Router{
							{Path: "verify"},
							{Path: "change"},
						},
					},
				},
			},
		},
	},
	{
		Path: "captcha",
		To:   "http://localhost:9001",
		Children: []Router{
			{
				Path: "image",
				Children: []Router{
					{Path: "reload"},
					{Path: "verify"},
				},
			},
		},
	},
	{
		Path: "vcode",
		To:   "http://localhost:9003",
		Children: []Router{
			{
				Path: "email",
				Children: []Router{
					{Path: "login"},
					{Path: "change-password"},
				},
			},
		},
	},
	{
		Path: "auth",
		To:   "http://localhost:9005",
		Children: []Router{
			{
				Path: "token",
				Children: []Router{
					{Path: "refresh"},
				},
			},
		},
	},
}
