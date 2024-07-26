# Pretty

Getting Started
===============

## Installing

## Value
pretty print all of [array, slice, struct, map ...]

Using this example:
```go
type DemoStruct struct {
	RequestId string   `json:"requestId"`
	Codes     []int    `json:"code"`
	Reasons   []string `json:"reasons"`
	Data      struct {
		RequestId string `json:"requestId"`
		Choices   []struct {
			Messages []struct {
				Name            string `json:"name"`
				RoleId          string `json:"roleId"`
				Role            string `json:"role"`
				Content         string `json:"content"`
				FinishReason    string `json:"finishReason"`
				ValidMessage    bool   `json:"validMessage"`
				UseMessage      bool   `json:"useMessage"`
				FunctionMessage bool   `json:"functionMessage"`
			} `json:"messages"`
			StopReason string `json:"stopReason"`
		} `json:"choices"`
		Usage struct {
			OutputTokens int `json:"outputTokens"`
			InputTokens  int `json:"inputTokens"`
			UserTokens   int `json:"userTokens"`
			PluginTokens int `json:"pluginTokens"`
		} `json:"usage"`
	    ...
		Stop bool `json:"stop"`
	} `json:"data"`
	Success bool `json:"success"`
}
```

The following code:
```go
result = pretty.Value(example)
```

```json
{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}
```

```
{
  RequestId: "request_id"
  Codes: [1,2,3,4,5,6]
  Reasons: ["beijing","nanjing"]
  Data: {
    RequestId: ""
    Choices: []
    Usage: {
      OutputTokens: 0
      InputTokens: 0
      UserTokens: 0
      PluginTokens: 0
    }
    ...
    Stop: false
  }
  Success: false
}
```

## Pretty

Using this example:

```json
{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
"children": ["Sara","Alex","Jack"],
"fav.movie": "Deer Hunter", "friends": [
    {"first": "Janet", "last": "Murphy", "age": 44}
  ]}
```

The following code:
```go
result = pretty.Pretty(example)
```

Will format the json to:

```json
{
  "name": {
    "first": "Tom",
    "last": "Anderson"
  },
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {
      "first": "Janet",
      "last": "Murphy",
      "age": 44
    }
  ]
}
```

## Color

Color will colorize the json for outputing to the screen.

```go
result = pretty.Color(json, nil)
```

Will add color to the result for printing to the terminal.
The second param is used for a customizing the style, and passing nil will use the default `pretty.TerminalStyle`.

