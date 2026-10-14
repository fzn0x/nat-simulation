---
name: nat-simulation
description: NAT masquerading taught as a brick build-instruction booklet, drawn in outlined SVG plates for a GitHub README.
colors:
  sky-page: "#AEE1FF"
  sky-grid: "#86CBF3"
  ink: "#111111"
  paper: "#FFFFFF"
  slate-muted: "#3B5566"
  slate-subline: "#24435A"
  brick-red: "#E53935"
  brick-red-left: "#C42B28"
  brick-red-right: "#9F201D"
  brick-red-edge: "#5E120F"
  brick-red-stud: "#EB6561"
  router-blue: "#1E84DA"
  router-blue-left: "#1468B2"
  router-blue-right: "#0E4F8A"
  router-blue-edge: "#082E52"
  router-blue-stud: "#509FE2"
  build-yellow: "#FFCD00"
  build-yellow-left: "#E3B000"
  build-yellow-right: "#BF9100"
  build-yellow-edge: "#6E5300"
  build-yellow-stud: "#FFD838"
  plate-grey: "#D9D9D9"
  plate-grey-left: "#BEBEBE"
  plate-grey-right: "#A2A2A2"
  plate-grey-edge: "#5C5C5C"
  plate-grey-stud: "#E1E1E1"
  tile-white-left: "#E4E4E4"
  tile-white-right: "#C9C9C9"
typography:
  display:
    fontFamily: "Archivo, sans-serif"
    fontWeight: 900
    letterSpacing: "-0.025em"
    fontVariation: "'wght' 900, 'wdth' 100"
  numeral:
    fontFamily: "Archivo, sans-serif"
    fontSize: "150px"
    fontWeight: 900
    lineHeight: 1
    fontVariation: "'wght' 900, 'wdth' 62"
  headline:
    fontFamily: "Archivo, sans-serif"
    fontSize: "40px"
    fontWeight: 800
    lineHeight: 1.1
  label:
    fontFamily: "Archivo, sans-serif"
    fontWeight: 700
    letterSpacing: "0.08em"
  body:
    fontFamily: "Archivo, sans-serif"
    fontWeight: 600
  data:
    fontFamily: "Martian Mono, monospace"
    fontWeight: 500
    fontVariation: "'wdth' 87.5, 'wght' 500"
  data-strong:
    fontFamily: "Martian Mono, monospace"
    fontWeight: 700
    fontVariation: "'wdth' 87.5, 'wght' 700"
rounded:
  chip: "8px"
  box: "14px"
  page: "26px"
  logo: "20%"
spacing:
  stud: "32px"
  margin: "40px"
  gutter: "36px"
components:
  page-plate:
    backgroundColor: "{colors.sky-page}"
    rounded: "{rounded.page}"
    width: "1200px"
    height: "600px"
  page-plate-narrow:
    backgroundColor: "{colors.sky-page}"
    rounded: "{rounded.page}"
    width: "720px"
  callout-box:
    backgroundColor: "{colors.paper}"
    textColor: "{colors.ink}"
    rounded: "{rounded.box}"
    width: "476px"
  label-chip:
    backgroundColor: "{colors.paper}"
    textColor: "{colors.ink}"
    typography: "{typography.data}"
    rounded: "{rounded.chip}"
  verdict-chip:
    backgroundColor: "{colors.paper}"
    textColor: "{colors.ink}"
    typography: "{typography.data-strong}"
    rounded: "{rounded.chip}"
  logo-tile:
    backgroundColor: "{colors.paper}"
    rounded: "{rounded.logo}"
  badge:
    backgroundColor: "{colors.sky-page}"
    textColor: "{colors.ink}"
---

# Design System: nat-simulation

## Overview

**Creative North Star: "The Build-Instruction Booklet"**

Every visual in nat-simulation reads like a page from a toy brick instruction booklet. Each NAT translation is one numbered assembly step. The part that changes arrives bright, with an arrow pointing to its seat. Parts that leave are ghosted. A paper-white 1:1 call-out box shows exactly the one part the step changes. The page is instruction-sky blue with a faint stud grid. Every line is an ink keyline, and the lettering is heavy and grotesk.

The world lives entirely in static SVG plates, because GitHub markdown allows no page CSS. All text is outlined to paths so it renders the same everywhere. Each plate repeats its meaning in nearby markdown and alt text. The picture is never the only place the information lives, and colour is never the only signal: every brick also carries a written address label.

The system deliberately avoids the networking-tool default of a dark terminal banner, a monospace wordmark and an ASCII topology. Monospace appears only where data appears.

**Key Characteristics:**
- Bricks with fixed jobs: red is the client's private address, blue is the router's public address, yellow is the server, grey is the packet and its payload.
- Isometric bricks at 30°, flat-shaded in three tones plus a dark edge, with no gradients.
- State is shown with marks, not hues: seated, arriving, ghosted, rejected, accepted.
- Ink keylines everywhere; paper call-outs on a sky page.
- Huge condensed step numerals. Reading order follows the arrows.
- Static by design. There is no motion.

## Colors

The palette is a primary-toy brick set on a pale instruction-sky page, held together by one ink black.

### Primary
- **Brick Red** (brick-red): the client's private address and nothing else. Use it for every src brick before the masquerade, the client entry in the parts list, and the hidden brick in the logo.
- **Router Blue** (router-blue): the router's public address and nothing else. Use it for the src brick after the masquerade, the router entry in the parts list, and the front brick in the logo.

### Secondary
- **Build Yellow** (build-yellow): the server, that is, the dst brick and the server parts entry.

### Tertiary
- **Plate Grey** (plate-grey): the packet plate and its payload.
- **Tile White** (paper top, with tile-white-left and tile-white-right sides): the flat hardware-address tile. It has no studs.

Each brick colour ships as a four-tone set: top (the base token), left face (`-left`), right face (`-right`) and edge keyline (`-edge`). Stud tops use the `-stud` tone, which is the top colour lightened by 22%. These tones are fixed. Do not mix new ones.

### Neutral
- **Instruction Sky** (sky-page): the page ground of every plate and the value side of README badges.
- **Grid Sky** (sky-grid): dashed stud-grid cell lines (dash 3/5, opacity .75) and one stud ring per 32-unit cell (opacity .7).
- **Ink** (ink): all keylines, lettering, arrows, seats, stop bars and leader lines.
- **Paper** (paper): call-out boxes, chips, the logo tile, and the halos around arrows and leader dots.
- **Muted Slate** (slate-muted): table column headers inside call-outs, and the label side of README badges.
- **Subline Slate** (slate-subline): the second line of the banner tagline, which states the scope.

### Named Rules
**The Fixed Jobs Rule.** A brick colour is a meaning, not a decoration. Red is only ever a client's private address. Blue is only ever the router's public address. Yellow is only ever the server. Grey is only ever the packet and its payload. A new plate may not reuse any of them for emphasis.

**The Blue Is Spoken For Rule.** Arrows, links in plates, highlights and selection marks are never blue, because blue already means the router. Arrows are ink.

**The Label Always Rides Rule.** Every brick has a written address label in a mono chip or call-out. Colour alone never carries meaning.

## Typography

**Display Font:** Archivo variable (OFL), instanced per role
**Body Font:** Archivo
**Label/Mono Font:** Martian Mono variable (OFL), width 87.5

**Character:** A heavy, slightly industrial grotesk sets the booklet's voice: loud numerals and titles, plain captions. A narrow, workmanlike mono is reserved for the addresses the lesson is actually about.

All text in the SVGs is converted to outlines. The values below are the instances that were outlined, measured in SVG user units on a 1200-wide plate.

### Hierarchy
- **Display** (900, width 100, tracking -0.025em): the `nat-simulation` wordmark on the banner only.
- **Numeral** (900, width 62, 150 units): the step numerals 1–4 at the top left of each plate, and the "1x" counts in call-outs (set smaller).
- **Headline** (800, 40 units wide / 36 narrow): plate titles, written as sentences ("The router swaps in its own address"). On narrow plates they wrap to at most two lines.
- **Label** (700, all caps, tracking +0.08em): box titles such as PIECES IN THIS SET, TRANSLATION TABLE, TABLE LOOKUP and ALLOW-LIST, each sitting above a 3px ink rule.
- **Body** (600): the banner tagline.
- **Data** (Martian Mono 500): addresses, ports and hardware addresses on chips, in call-outs and in table rows.
- **Data Strong** (Martian Mono 700): verdict chips (`200 Hello!`, `403 Not Permitted`) and table column headers (HWADDR, INSIDE ADDRESS).

### Named Rules
**The Mono Means Data Rule.** Martian Mono appears only on literal values a program would print: addresses, ports, hardware addresses, status lines and table headers. Titles, counts and the wordmark are never mono.

**The Outlined Text Rule.** SVG text is always converted to paths, because an SVG loaded through `<img>` cannot load fonts.

## Layout

Plates are measured in stud units. One stud pitch is 1 unit, which is 32 units on the page grid.

- **Wide plate** (1200 × 600): the numeral and title run across the top. The isometric assembly sits on the left (x 40–640). The call-out column starts at x 676 and is 476 wide. Outer margin is 40.
- **Narrow plate** (720 wide, height follows the content): the numeral and title sit at the top, the title wraps to two lines at most, and the call-outs stack below the assembly at full width.
- **Responsive delivery:** each step ships as a `<picture>`. The narrow variant is served at `max-width: 600px` so the call-out lettering stays legible on phones. The banner and logo are single images.
- **Reading order** is led by the arrows: numeral, then title, then the moving part, then the call-out.
- **README column:** the plain GitHub column is not restyled. Badges use `flat-square` with slate-muted labels and sky values.

## Elevation & Depth

Depth is drawn, not cast. Everything is isometric at 30°, flat-shaded in three face tones with a darker edge keyline. Brick height is 1.2 units, plate thickness 0.4, stud radius 0.3 and stud height 0.24. There are no drop shadows, no blur and no gradients. Call-out boxes and chips are flat paper shapes with ink keylines.

### Named Rules
**The Three Faces Rule.** A solid has exactly a top, a left face and a right face, each a flat fill from its brick's tone set and outlined in the edge tone. A light source is never implied by gradient or shadow.

**The Halo, Not Shadow Rule.** When an ink mark crosses a brick, it is separated with a paper halo (4px on arrows, a white ring on leader dots), never with a shadow.

## Shapes

The language is rounded rectangles with heavy ink keylines around isometric solids.

- **Page plate:** radius 26 with a 4px ink keyline.
- **Call-out box:** radius 14 with a 4px ink keyline.
- **Chip:** radius 8 with a 3px ink keyline.
- **Logo tile:** radius about 20% of the tile, with a heavy ink keyline.
- **Seats:** a dashed ink outline of the footprint a part will occupy.
- **Ghosted parts:** dashed edges at opacity .4.

A packet is always a 5×2 grey plate carrying a 2×2 src brick, a 1×2 white hardware-address tile, and a 2×2 dst brick, in that order.

## Components

### Call-out Box (signature)
The 1:1 box shows exactly the one part a step changes.
- **Shape:** paper fill, 4px ink keyline, radius 14, inside the 476-wide call-out column.
- **Content:** the brick drawn at true size, a "1x" count in the condensed numeral, and the address in a mono label.
- **Titled variant:** a caps label title over a 3px ink rule (PIECES IN THIS SET, ALLOW-LIST).
- **Table variant:** two columns, HWADDR and INSIDE ADDRESS. Headers are Data Strong in slate-muted; rows are Data in ink.

### Chips
- **Label chip:** paper fill, 3px ink keyline, radius 8, mono text. It names an address on the assembly (`src 192.168.1.10`, `hw 02:42:c0:a8:01:0a`, `body "Hello!"`).
- **Verdict chip:** the same shape with Data Strong text, led by a drawn ✓ (accepted) or ✕ (rejected) mark.

### Leader Line
A 3px ink line from a chip to its part, ending in an ink dot with a white halo.

### Arrows and State Marks
- **Arrow:** ink with a 4px paper halo. Never blue.
- **Seated:** normal rendering.
- **Arriving:** the part hovers above a dashed ink seat, with a solid ink arrow pointing down into it.
- **Lifted off:** the part is ghosted (opacity .4, dashed edges).
- **Rejected:** a solid ink bar stops the arrow, followed by a ✕ verdict chip.
- **Accepted:** the arrow reaches its box, followed by a ✓ verdict chip.

### Logo Tile
A paper tile with an ink keyline and a radius of about 20%, holding a blue brick standing in front of a red one: the router hiding its client.

### Badges
Shields.io `flat-square` badges with a slate-muted label and a sky-page value.

## Do's and Don'ts

### Do:
- **Do** give every changing part its own numbered step, a call-out box with "1x", and a mono address label.
- **Do** build every solid from its four-tone brick set (top, left, right, edge) plus the stud tone.
- **Do** show state with marks (seat, arrow, ghost, stop bar, ✓/✕ chip), keeping each brick its job colour in every state.
- **Do** ship a 720-wide narrow variant of every step plate and serve it through `<picture>` at `max-width: 600px`.
- **Do** outline all SVG text, and keep each plate's alt text and the README text beside it in sync.

### Don't:
- **Don't** use red, blue, yellow or grey for anything other than their fixed jobs.
- **Don't** draw arrows, highlights or links in blue; blue is reserved for the router.
- **Don't** add gradients, drop shadows or blur to bricks or boxes; depth comes from the three flat faces.
- **Don't** set titles, counts or the wordmark in monospace; Martian Mono is for data only.
- **Don't** let colour be the only signal; every brick carries a written label.
