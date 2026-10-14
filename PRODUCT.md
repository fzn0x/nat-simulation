# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

(The visual surfaces are the GitHub-rendered README and its SVG assets. The programs themselves are command-line tools.)

## Stack

Go, standard library only. Three programs (server, router, client), each in its own folder with a launcher script (`serverexec`, `routerexec`, `clientexec`). Chosen by the user.

## Users

Self-learners: students and developers studying on their own what `iptables -t nat -A POSTROUTING -j MASQUERADE` actually does to a packet. They read the README and the source side by side, run the three programs in three terminals, and watch the logs. (Inferred from the brief's "a toy for learning"; the user did not pick an audience explicitly.)

## Product Purpose

nat-simulation simulates NAT masquerading with three cooperating programs and never touches the real kernel firewall. It's a toy for learning how a router hides clients behind its own address. It is not a firewall tool.

Success: after one run, the learner can explain why the routed request gets "Hello!" and the direct one gets "Not Permitted", and can point to the router code that rewrites the source address on the way out and restores the destination on the way back.

## Positioning

It shows the masquerade on the wire you can read. Every rewrite is visible as a field changing in a JSON-ish packet and in the router's translation table, without needing root, network namespaces, or Linux. The direct-vs-routed pair, sent every 4 seconds, is the built-in controlled experiment.

## Operating Context

- Three terminals, one per program, on any OS with Go installed.
- The client sends the same request twice every 4 seconds: once through the router, once directly to the server.
- The server holds an allow-list of source addresses. It trusts only the router's public address.
- The router records each original source against a hardware (MAC-style) address. On the way out it rewrites the source to its own address. On the way back it looks up the hardware address and restores the real destination.

## Capabilities and Constraints

- Educational only. Never claim security, performance, or production use.
- No real packets are rewritten. Addresses are simulated fields carried in the request, because all three programs run on localhost.
- No third-party dependencies.
- License: MIT, copyright holder **fzn0x**.

## Brand Commitments

- Name: `nat-simulation`.
- Attribution only under the handle **fzn0x**. Never use the maintainer's real name.

## Evidence on Hand

No stars, users, testimonials, or benchmarks exist. Don't invent any. The only proof is the program's own output, which the README may show verbatim from a real run.

## Product Principles

1. Show the rewrite, don't describe it: every translation the router makes is printed.
2. One concept per file: the server knows allow-lists, the router knows translation, the client knows the experiment.
3. Honest scope: say plainly what is simulated and how real MASQUERADE (conntrack, port rewriting) differs.
4. Zero setup friction: stdlib Go, no root, any OS.
