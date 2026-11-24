# How to desing http service

## Mat Ryer's Blog

### Findings

- Try to seprate out the main from errors
- Use seprate setup method named `run` to do that
- `run()` can take on the all required parameters foe exmaple below code.

  ```go
        func run(
        ctx    context.Context,
        args   []string,
        getenv func(string) string,
        stdin  io.Reader,
        stdout, stderr io.Writer,
        ) error
    ```

- Also using the function to create the `http.handler` is better type sefty wise and provides user idea about which parameters are exactly getting used and where. That doesn't discoount user of `structs` where it makes sense to use.
- He generally prefers one mthod doing the setting up `NewServer()` for example belwo code.

```go
    srv := NewServer(
        logger,
        config,
        tenantsStore,
        commentsStore,
        conversationService,
        chatGPTService,
    )
```

- This method evantually will look like belwo,

```go
    func NewServer(
        logger *Logger
        config *Config
        commentStore *commentStore
        anotherStore *anotherStore
    ) http.Handler {
        mux := http.NewServeMux()
        addRoutes(
            mux,
            Logger,
            Config,
            commentStore,
            anotherStore,
        )
        var handler http.Handler = mux
        handler = someMiddleware(handler)
        handler = someMiddleware2(handler)
        handler = someMiddleware3(handler)
        return handler
    }
```

- So if you look at the code you can see that he likes to wrap everything in single place where it just returns the `http.handler`. Error is also not returned so its same to assume its been propogated till th actual middleware or handler and then buck stops there.
- Then standard stuff in run method
    . Start with signal Nofity context to catch the intrrupt
    . Create the server
    . Start rutine to start the server
    . Waitgroup to handle termination/server closing rutine
    . Termination rutine setup
    . waitgroup blocking wait.

    The termination rutine looks like bleow,

    ```go
        go func() {
            defer wg.Done()
            <-ctx.Done()
            shutdownCtx := context.Background()
            shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10 * time.Second)
            defer cancel()
            if err := httpServer.Shutdown(shutdownCtx); err != nil {
                fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
            }
        }()
    ```

- Main just calls the run with requried parameters
- He likes to handler implement the `handlerFunc` and not the handler in itself so they look like below,

```go
    // handleSomething handles one of those web requests
    // that you hear so much about.
    func handleSomething(logger *Logger) http.Handler {
        thing := prepareThing()
        return http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                // use thing to handle request
                logger.Info(r.Context(), "msg", "handleSomething")
            }
        )
    }
```

- Likes to also keep the routes to one function so its easy to manage. like below,

```go
    func addRoutes(
        mux                 *http.ServeMux,
        logger              *logging.Logger,
        config              Config,
        tenantsStore        *TenantsStore,
        commentsStore       *CommentsStore,
        conversationService *ConversationService,
        chatGPTService      *ChatGPTService,
        authProxy           *authProxy
    ) {
        mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
        mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
        mux.HandleFunc("/healthz", handleHealthzPlease(logger))
        mux.Handle("/", http.NotFoundHandler())
    }
```