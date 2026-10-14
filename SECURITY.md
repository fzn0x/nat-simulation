# Security policy

nat-simulation is a teaching toy. It never touches the real kernel firewall, and it must not be used to protect anything. The "router" only rewrites fields in a JSON message.

That said, the three programs do open local HTTP listeners. If you find a way they could harm the machine they run on, for example by crashing on crafted input in a way that matters or by listening beyond `127.0.0.1` without being asked to, please report it privately:

1. Go to the repository's **Security** tab.
2. Choose **Report a vulnerability**.

Please don't open a public issue for this kind of report. You'll get a reply as soon as the maintainer can look at it. Only the latest commit on `main` is supported.
