UNHTML
======

[![Build Status](https://travis-ci.org/ancientlore/unhtml.svg?branch=master)](https://travis-ci.org/ancientlore/unhtml)
[![GoDoc](https://godoc.org/github.com/ancientlore/unhtml?status.svg)](https://godoc.org/github.com/ancientlore/unhtml)

Package unhtml is designed to remove HTML tags from text and do minor formatting updates.
It is intended to avoid changing the text very much - in particular so that it doesn't
mess up formatted plain text when HTML tags are not present. This allows you to run the
converter on data that you received without needing to bother checking if HTML tags are
in there.

That said, the package is primarily intended to handle minor HTML snippets. It isn't a
full-fledged formatter. Below are tags that are ignored, handled, and intentionally
skipped.

Ignored tags:

	(comments)
	DOCTYPE
	abbr
	acronym
	address
	area
	article
	aside
	audio
	b
	base
	basefont
	bdi
	bdo
	big
	body
	button
	canvas
	caption
	center
	cite
	code
	col
	colgroup
	datalist
	dd
	del
	details
	dfn
	dialog
	dir
	dl
	dt
	em
	fieldset
	figcaption
	figure
	font
	footer
	form
	header
	html
	i
	input
	ins
	kbd
	keygen
	label
	legend
	main
	map
	mark
	menu
	menuitem
	meter
	nav
	noframes
	noscript
	optgroup
	option
	output
	param
	progress
	q
	rp
	rt
	ruby
	s
	samp
	section
	select
	small
	source
	span
	strike
	strong
	sub
	summary
	sub
	summary
	sup
	tbody
	textarea
	tfoot
	thead
	time
	track
	tt
	u
	var
	video
	wbr

Skipped tags:

	applet
	embed
	frame
	frameset
	head
	iframe
	link
	meta
	object
	script
	style
	title


Handled tags:

	a
	blockquote
	br
	div
	h1 to h6
	hr
	img
	li
	ol
	p
	pre
	table
	td
	th
	tr
