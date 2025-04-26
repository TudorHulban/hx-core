package widgets

import pagecss "github.com/TudorHulban/hx-core/page-css"

func CSSBase() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		html {
	line-height: 1.15;
	-webkit-text-size-adjust: 100%;
}

body {
	margin: 0;
}

main {
	display: block;
}

h1 {
	font-size: 2em;
	margin: 0.67em 0;
}

hr {
	box-sizing: content-box;
	height: 0;
	overflow: visible;
}

pre {
	font-family: monospace, monospace;
	font-size: 1em;
}

a {
	background-color: transparent;
}

abbr[title] {
	border-bottom: none;
	text-decoration: underline;
	text-decoration: underline dotted;
}

b,
strong {
	font-weight: bolder;
}

code,
kbd,
samp {
	font-family: monospace, monospace;
	font-size: 1em;
}

small {
	font-size: 80%;
}

sub,
sup {
	font-size: 75%;
	line-height: 0;
	position: relative;
	vertical-align: baseline;
}

sub {
	bottom: -0.25em;
}

sup {
	top: -0.5em;
}

img {
	border-style: none;
}

button, .mpa-employees-list-shortcode .entry-link,
.wp-block-getwid-post-carousel .entry-link,
input,
optgroup,
select,
textarea {
	font-family: inherit;
	font-size: 100%;
	line-height: 1.15;
	margin: 0;
}

button, .mpa-employees-list-shortcode .entry-link,
.wp-block-getwid-post-carousel .entry-link,
input {
	overflow: visible;
}

button, .mpa-employees-list-shortcode .entry-link,
.wp-block-getwid-post-carousel .entry-link,
select {
	text-transform: none;
}

button, .mpa-employees-list-shortcode .entry-link,
.wp-block-getwid-post-carousel .entry-link,
[type="button"],
[type="reset"],
[type="submit"] {
	-webkit-appearance: button;
}

button::-moz-focus-inner, .mpa-employees-list-shortcode .entry-link::-moz-focus-inner,
.wp-block-getwid-post-carousel .entry-link::-moz-focus-inner,
[type="button"]::-moz-focus-inner,
[type="reset"]::-moz-focus-inner,
[type="submit"]::-moz-focus-inner {
	border-style: none;
	padding: 0;
}

button:-moz-focusring, .mpa-employees-list-shortcode .entry-link:-moz-focusring,
.wp-block-getwid-post-carousel .entry-link:-moz-focusring,
[type="button"]:-moz-focusring,
[type="reset"]:-moz-focusring,
[type="submit"]:-moz-focusring {
	outline: 1px dotted ButtonText;
}

fieldset {
	padding: 0.35em 0.75em 0.625em;
}

legend {
	box-sizing: border-box;
	color: inherit;
	display: table;
	max-width: 100%;
	padding: 0;
	white-space: normal;
}

progress {
	vertical-align: baseline;
}

textarea {
	overflow: auto;
}

[type="checkbox"],
[type="radio"] {
	box-sizing: border-box;
	padding: 0;
}

[type="number"]::-webkit-inner-spin-button,
[type="number"]::-webkit-outer-spin-button {
	height: auto;
}

[type="search"] {
	-webkit-appearance: textfield;
	outline-offset: -2px;
}

[type="search"]::-webkit-search-decoration {
	-webkit-appearance: none;
}

::-webkit-file-upload-button {
	-webkit-appearance: button;
	font: inherit;
}

details {
	display: block;
}

summary {
	display: list-item;
}

template {
	display: none;
}

[hidden] {
	display: none;
}

*,
*::before,
*::after {
	box-sizing: inherit;
}

html {
	box-sizing: border-box;
}
		`,
	}
}
