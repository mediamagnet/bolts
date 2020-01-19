<div align='center'>
  <img src="/docs/disgord-draft-8.jpeg" alt='Build Status' />
  <p>
    <a href="https://codecov.io/gh/andersfylling/disgord">
      <img src="https://codecov.io/gh/andersfylling/disgord/branch/develop/graph/badge.svg" />
    </a>
    <a href='https://goreportcard.com/report/github.com/andersfylling/disgord'>
      <img src='https://goreportcard.com/badge/github.com/andersfylling/disgord' alt='Code coverage' />
    </a>
    <a href='http://godoc.org/github.com/andersfylling/disgord'>
      <img src='https://godoc.org/github.com/andersfylling/disgord?status.svg' alt='Godoc' />
    </a>
  </p>
  <p>
    <a href='https://discord.gg/fQgmBg'>
      <img src='https://img.shields.io/badge/Discord%20Gophers-%23disgord-blue.svg' alt='Discord Gophers' />
    </a>
    <a href='https://discord.gg/HBTHbme'>
      <img src='https://img.shields.io/badge/Discord%20API-%23disgord-blue.svg' alt='Discord API' />
    </a>
  </p>
</div>

## About
Go module with context support that handles some of the difficulties from interacting with Discord's bot interface for you; websocket sharding, auto-scaling of websocket connections, advanced caching, helper functions, middlewares and lifetime controllers for event handlers, etc.

## Warning
The develop branch is under continuous breaking changes, as the interface and exported funcs/consts are still undergoing planning. Because DisGord is under development and pushing for a satisfying interface, the SemVer logic is not according to spec. Until v1.0.0, every minor release is considered possibly breaking and patch releases might contain additional features. Please see the issue and current PR's to get an idea about coming changes before v1.

There might be bugs in the cache, or the cache processing might not exist yet for some REST methods. Bypass the cache for REST methods by supplying the flag argument `disgord.IgnoreCache`. eg. `client.GetCurrentUser(disgord.IgnoreCache)`.

Remember to read the docs/code for whatever version of disgord you are using. This README file reflects the latest state in the develop branch, or at least, I try to reflect the latest state.

## Data types & tips
 - Use disgord.Snowflake, not snowflake.Snowflake.
 - Use disgord.Time, not time.Time when dealing with Discord timestamps. This is because Discord returns a weird time format.

## Starter guide
> This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dealing with dependencies, remember to activate module support in your IDE

> Examples can be found in [docs/examples](docs/examples) and some open source projects DisGord projects in the [wiki](https://github.com/andersfylling/disgord/wiki/A-few-DisGord-Projects)

I highly suggest reading the [Discord API documentation](https://discordapp.com/developers/docs/intro) and the [DisGord go doc](http://godoc.org/github.com/andersfylling/disgord).

Here is a basic bot program that prints out every message. Save it as `main.go`, run `go mod init bot` and `go mod download`. You can then start the bot by writing `go run .`

```go
package main

import (
	  "context"
    "fmt"
    "github.com/andersfylling/disgord"
    "os"
)

func printMessage(session disgord.Session, evt *disgord.MessageCreate) {
    msg := evt.Message
    fmt.Println(msg.Author.String() + ": "+ msg.Content) // Anders#7248{435358734985}: Hello there
}

func main() {
    client := disgord.New(disgord.Config{
        BotToken: os.Getenv("DISGORD_TOKEN"),
        // You can inject any logger that implements disgord.Logger interface (eg. logrus)
        // DisGord provides a simple logger to get you started. Nothing is logged if nil.
        Logger: disgord.DefaultLogger(false), // debug=false
    })
    // connect, and stay connected until a system interrupt takes place
    defer client.StayConnectedUntilInterrupted(context.Background())
    
    // create a handler and bind it to new message events
    // handlers/listener are run in sequence if you register more than one
    // so you should not need to worry about locking your objects unless you do any
    // parallel computing with said objects
    client.On(disgord.EvtMessageCreate, printMessage)
}
```

#### Linux script
To create a new bot you can use the disgord.sh script to get a bot with some event middlewares and Dockerfile. Paste the following into your terminal:

```bash
bash <(curl -s -L https://git.io/disgord-script)
``` 

Starter guide as a gif: https://terminalizer.com/view/469961d0695


## Architecture & Behavior
Discord provide communication in different forms. DisGord tackles the main ones, events (ws), voice (udp + ws), and REST calls.

You can think of DisGord as layered, in which case it will look something like:
![Simple way to think about DisGord architecture from a layered perspective](docs/disgord-layered-version.png)

#### Events
For Events, DisGord uses the [reactor pattern](https://dzone.com/articles/understanding-reactor-pattern-thread-based-and-eve). Every incoming event from Discord is processed and checked if any handler is registered for it, otherwise it's discarded to save time and resource use. Once a desired event is received, DisGord starts up a Go routine and runs all the related handlers in sequence; avoiding locking the need to use mutexes the handlers. 

In addition to traditional handlers, DisGord allows you to use Go channels. Note that if you use more than one channel per event, one of the channels will randomly receive the event data; this is how go channels work. It will act as a randomized load balancer.

But before either channels or handlers are triggered, the cache is updated.

#### REST
The "REST manager", or the `httd.Client`, handles rate limiting for outgoing requests, and updated the internal logic on responses. All the REST methods are defined on the `disgord.Client` and checks for issues before the request is sent out.

If the request is a standard GET request, the cache is always checked first to reduce delay, network traffic and load on the Discord servers. And on responses, regardless of the http method, the data is copied into the cache.

Some of the REST methods (updating existing data structures) will use the builder+command pattern. While the remaining will take a simple config struct. 

> Note: Methods that update a single field, like SetCurrentUserNick, does not use the builder pattern.
```go
// bypasses local cache
client.GetCurrentUser(context.Background(), disgord.IgnoreCache)
client.GetGuildMembers(context.Background(), guildID, disgord.IgnoreCache)

// always checks the local cache first
client.GetCurrentUser(context.Background())
client.GetGuildMembers(context.Background(), guildID)
```

#### Voice
Whenever you want the bot to join a voice channel, a websocket and UDP connection is established. So if your bot is currently in 5 voice channels, then you have 5 websocket connections and 5 udp connections open to handle the voice traffic.

#### Cache
The cache tries to represent the Discord state as accurate as it can. Because of this, the cache is immutable by default. Meaning the does not allow you to reference any cached objects directly, and every incoming and outgoing data of the cache is deep copied.

## Contributing
> Please see the [CONTRIBUTING.md file](CONTRIBUTING.md) (Note that it can be useful to read this regardless if you have the time)

You can contribute with pull requests, issues, wiki updates and helping out in the discord servers mentioned above.

To notify about bugs or suggesting enhancements, simply create a issue. The more the better. But be detailed enough that it can be reproduced and please provide logs.

To contribute with code, always create an issue before you open a pull request. This allows automating change logs and releases.

## Q&A
> **NOTE:** To see more examples go to the [docs/examples folder](docs/examples). See the GoDoc for a in-depth introduction on the various topics.

```Markdown
1. How do I find my bot token and/or add my bot to a server?

Tutorial here: https://github.com/andersfylling/disgord/wiki/Get-bot-token-and-add-it-to-a-server
```

```Markdown
2. Is there an alternative Go package?

Yes, it's called DiscordGo (https://github.com/bwmarrin/discordgo). Its purpose is to provide a minimalistic API wrapper for Discord, it does not handle multiple websocket sharding, scaling, etc. behind the scenes such as DisGord does.
Currently I do not have a comparison chart of DisGord and DiscordGo. But I do want to create one in the 
future, for now the biggest difference is that DisGord does not support self bots.
```

```Markdown
3. Why make another Discord lib in Go?

I'm trying to take over the world and then become a intergalactic war lord. Have to start somewhere.
```

```Markdown
4. Does this project re-use any code from DiscordGo?

Yes. See guild.go. The permission consts are pretty much a copy from DiscordGo.
```

```Markdown
5. Will DisGord support self bots?

No. Self bots are againts ToS and could result in account termination (see
https://support.discordapp.com/hc/en-us/articles/115002192352-Automated-user-accounts-self-bots-). 
In addition, self bots aren't a part of the official Discord API, meaning support could change at any 
time and DisGord could break unexpectedly if this feature were to be added.
```

