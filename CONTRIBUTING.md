# Contributing to nat-simulation

Thanks for helping make NAT easier to learn. This project is a teaching toy: three small Go programs that act out what `iptables -t nat -A POSTROUTING -j MASQUERADE` does to a packet. Every contribution is judged by one question:

> **Does this make the masquerade easier to see or understand?**

If the answer is yes, it belongs here, even if it's a one-word fix to a log line.

## What we welcome

- **Clearer output.** A log line that explains a rewrite better than the current one.
- **Better explanations.** README sections, code comments that explain the *networking* (not the Go), diagrams.
- **New experiments.** Flags or scenarios that let a learner poke at NAT: a second client, a port clash, a dropped reply.
- **Tests** that pin down what the router, server or client must do.
- **Bug reports**, including "I followed the README and got lost here". Confusion is a bug in a teaching project.

## What we'll usually decline

- **Real packet handling.** Raw sockets, netfilter bindings, TUN/TAP devices. The point is to run anywhere, without root.
- **Third-party dependencies.** The standard library keeps the code readable from top to bottom.
- **Features that blur the lesson,** such as authentication, persistence, config files or a web UI.

Not sure which side your idea falls on? [Open an issue](https://github.com/fzn0x/nat-simulation/issues/new/choose) and ask before you write code.

## Getting set up

You need [Go 1.21 or newer](https://go.dev/dl/). Nothing else.

```sh
git clone https://github.com/fzn0x/nat-simulation.git
cd nat-simulation
go test ./...
```

Run the simulation in three terminals as the [README](README.md#build-it) describes, and keep it running while you work: the logs are the fastest way to see what your change did.

## Before you open a pull request

Run the same checks CI runs:

```sh
gofmt -l .        # should print nothing
go vet ./...
go test ./...
```

Then:

1. **Keep each program to one idea.** The server knows allow-lists, the router knows translation, the client knows the experiment. Shared code goes in `internal/packet`.
2. **Show the change.** If your change alters what the programs print, paste the before and after output in the pull request.
3. **Update the README** when you add a flag, change a default or change the output it quotes.
4. **Keep pull requests small.** One change per pull request is easier to review and easier to learn from in the history.

## Commit messages

We use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). The history doubles as a changelog, so each commit should describe one change.

```text
<type>(<optional scope>): <what changed, in the imperative>
```

| Type       | Use it for                                    |
| ---------- | --------------------------------------------- |
| `feat`     | New behaviour, flags or experiments           |
| `fix`      | Something that behaved wrongly                |
| `docs`     | README, CONTRIBUTING, comments-only changes   |
| `test`     | Adding or fixing tests                        |
| `refactor` | Code changes that don't change behaviour      |
| `ci`       | GitHub Actions workflows                      |
| `chore`    | Repo settings, tooling, housekeeping          |

Scopes match the folders: `client`, `router`, `server`, `packet`, plus `brand` for the logo and artwork in `assets/`.

```text
feat(router): print the translation table after every change
fix(client): wrap the source port back to 49152 after 65535
docs: explain why replies carry the hardware address
```

## Artwork

The logo and README plates in [`assets/`](assets) follow a brick build-instruction style. Each colour has a fixed job, so keep them consistent:

| Colour                          | Means                              |
| ------------------------------- | ---------------------------------- |
| Brick red `#E53935`             | A client's private address         |
| Accent blue `#147BD1`           | The router's public address        |
| Build yellow `#FFCD00`          | The server                         |
| Plate grey `#D9D9D9`            | The packet and its payload         |
| Sky blue `#AEE1FF`              | The page                           |

Text in the SVGs is converted to outlines so it renders the same everywhere. If you change a plate, keep its alt text and the README text beside it in sync, because the picture is never the only place the information lives.

## Code of conduct

Everyone taking part is expected to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## License

By contributing, you agree that your contributions are licensed under the [MIT License](LICENSE).
