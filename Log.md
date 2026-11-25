# Daily Tracker (Discipline wins over talent)

## Week 1 (23-29 Nov 2026): Ingest

**Goal (revised 24 Nov):** a running API accepts a valid event and rejects an invalid one, with tests. Storing events in Postgres moves to week 2, because each reading takes a full session with notes.

**Done when:**

- [ ] The repo is public on GitHub, with a README paragraph saying what relay is.
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

### 25 Nov 2026 (normal day) - 5 Min

1. Just updated the Log.md file and committed

**Next (25 Nov, build day):** go through the Mat Ryer notes, then write an HTTP server with a health endpoint that answers curl. Minimum: a server that starts and answers one request.

**Rest of the week:** Thu 26 reading day, second half of Mat Ryer's post with notes. Fri 27 build day, the event endpoint with validation. Sat 28 build day, tests for the validation, and buffer. Sun 29 check-in.
