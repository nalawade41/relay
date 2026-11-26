# Daily Tracker (Discipline wins over talent)

## Week 1 (23-29 Nov 2026): Ingest

**Goal (revised 24 Nov):** a running API accepts a valid event and rejects an invalid one, with tests. Storing events in Postgres moves to week 2, because each reading takes a full session with notes.

**Done when:**

- [x] The repo is public on GitHub, with a README paragraph saying what relay is.
- [ ] One command starts the API, and its health endpoint answers.
- [ ] A valid event sent with curl is accepted.
- [ ] An invalid event is rejected with a clear error.
- [ ] Tests cover the validation and pass.
- [ ] Notes for each reading are in `reading/`.

### 23 Nov 2026 (minimum day) - 15 Min

1. Add Log.md file and record today's activity
2. Intilize the git with `git init`
3. Create git repo on github
4. Link github repo with local folder
5. Intilize go with `go mod init`
6. Verfiy the `.gitignore`
7. Push the files
8. Add README.md

### 24 Nov 2026 (normal day) - 90 Min

1. Read through Mat Ryer's first half off blog
2. Understand and take notes
3. Create reading material folder to keep the notes
4. Add README.md short paragraph about project.
5. Create review skill for claude.
6. Record in Log.md
7. Commit and Push

### 25 Nov 2026 (Worst day) - 5 Min

1. Just updated the Log.md file and committed

### 26 Nov 2026 (Very Good day) - 70 Min

1. Read thorugh notes on mat's blog
2. Added main.go with runner
3. Created config, encoder, handlers, middlweare, router, runner, server, store packages
4. Added files in above created packages
5. Added minimum code requried to build the project
6. Added `/healthz` end point and its code
7. Added code for server setup and middleware to log details
8. Added code to serve the `healthz` end point
9. Updated LOG.md

**Next (27 Nov, big day):** tests, the health endpoint checked with curl, then a review. Do the steps in order. Each one stands on its own, so stop wherever time ends and the rest moves to Saturday.

1. Start the server and call `/healthz` with curl. Ticks the second box above.
2. Run `go vet ./...` and fix what it reports in `runner`. `go test` will not run that package until it is clean.
3. Read, with notes in `reading/`: four sections of Mat Ryer's post: "Designing for testability", "What is the unit when unit testing?", "Testing with the run function", "Waiting for readiness".
4. Write tests for the handler, encoder, middleware and router, then one test that starts the whole program through `run` and waits on `/healthz`.
5. Measure coverage per package. Target: above 95%.
6. Run the review skill.
7. Fix what the review puts under "Fix before commit". Commit and push.

**Look up if stuck (no notes needed):**

- Calling a handler from a test: <https://pkg.go.dev/net/http/httptest>
- Measuring coverage and seeing which lines are missed: <https://go.dev/blog/cover>
- One test function, many cases: <https://go.dev/wiki/TableDrivenTests>

**Minimum:** steps 1 and 2, plus one test for the health handler.

**Rest of the week:** Sat 28 anything left from Friday first, then reading with notes: the sections of Mat Ryer's post from "Validating data" to "`sync.Once` to defer setup". Sun 29 check-in. The event endpoint with validation moves to the start of week 2.
