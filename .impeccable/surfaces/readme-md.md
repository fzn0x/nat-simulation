---
version: 1
slug: "readme-md"
primary_target: "README.md"
related_targets: ["assets"]
---

# Surface brief: README.md

Scope: the GitHub-rendered README and every image it embeds (logo, banner, step plates). Visitor mode: **Read**.

Reader: a self-learner who wants to know what MASQUERADE does to a packet, then run it. Their question: "What does the router actually change, and how does the reply find its way back?" Success: they run three terminals, see "Hello!" vs "Not Permitted", and can point at the rewrite.

Constraints: GitHub markdown only (no page CSS), so the world lives in SVG plates. All SVG text is outlined to paths (an `<img>` SVG can't load fonts). Every plate keeps its meaning in nearby markdown text and alt text. Colour is never the only signal: every brick carries a written address label.

## Direction contract

THESIS: NAT taught as a brick build-instruction booklet. Each translation is one numbered assembly step: the part that changes arrives bright with an arrow to its seat, and earlier parts are ghosted. This refuses the category default of a dark terminal banner, a mono wordmark and an ASCII topology.

OWN-WORLD: Instruction-page sky blue (#AEE1FF) as the page ground, with a faint dashed stud grid. Ink black (#111111) keylines and heavy grotesk lettering. Bricks have fixed jobs: brick red = a client's private address, accent blue = the router's public address, build yellow = the server, plate grey = the payload. Paper-white 1:1 call-out boxes with ink keylines, "1x" counts, and huge condensed step numerals. Isometric bricks are flat-shaded in three tones, with no gradients.

STORY: The reader sees the parts inventory (1x client, 1x router, 1x server), follows four numbered steps (send, swap the source, check the allow-list, restore the reply), then builds it themselves in three terminals and compares the real log output.

FIRST VIEWPORT: A full-width sky-blue banner page. The logo tile (a blue brick hiding a red one) is at top left. Below it, a huge black "nat-simulation" wordmark and a one-line tagline. On the right is a "pieces in this set" inventory box listing 1x client, 1x router and 1x server, each with its brick. Under the banner come the plain README column, the intro paragraph, then "What you'll see" with real client output.

FORM: Brick build-instruction booklet (catalog challenger "games-toys-physics-play-brick-build-instructions", chosen by the user over the assigned Cloakroom Ticket). Seed key 4ea1e387. Signature move: the 1:1 call-out box showing exactly the one part a step changes. Signature motion: none needed on GitHub. Plates are static, with an arrow-led reading order.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance
