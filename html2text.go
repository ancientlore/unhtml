/*
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
*/
package unhtml

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

// HtmlToTextString converts a string of HTML into a string of plain text.
func HtmlToTextString(in string) (string, error) {
	var buf bytes.Buffer
	rdr := bytes.NewBufferString(in)
	err := HtmlToText(rdr, &buf)
	return string(buf.Bytes()), err
}

// HtmlToText converts the HTML in the reader to text in the writer.
func HtmlToText(in io.Reader, out io.Writer) error {

	z := html.NewTokenizer(in)

	var lastTT html.TokenType = html.CommentToken
	var newLines int = 3
	var url string = ""
	var ignore bool = false
	var reallyIgnore bool = false

	for done := false; !done; {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err := z.Err()
			if err != io.EOF {
				return err
			}
			done = true
		case html.TextToken:
			if !ignore && !reallyIgnore {
				text := strings.TrimSpace(string(z.Text()))
				if text != "" && newLines == 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, text)
				if text != "" || url != "" {
					newLines = 0
				}
				if url != "" {
					fmt.Fprintf(out, " (%s)", url)
					url = ""
				}
			}
		case html.StartTagToken:
			tag, hasAttr := z.TagName()
			//if lastTT == html.TextToken && newLines == 0 {
			//	fmt.Fprint(out, " ")
			//}
			switch string(tag) {
			case "head":
				reallyIgnore = true
			case "script", "title", "link", "meta", "applet", "embed", "frame", "frameset", "iframe", "object", "style":
				ignore = true
			case "div", "ul", "tr", "ol", "p", "br", "table":
				if newLines < 2 {
					fmt.Fprintln(out)
					newLines++
				}
			case "h1", "h2", "h3", "h4", "h5", "h6", "pre", "blockquote":
				if newLines < 2 {
					fmt.Fprintln(out, "\n")
					newLines++
				}
			case "li":
				if newLines < 2 {
					fmt.Fprintln(out)
				}
				fmt.Fprint(out, "* ")
				newLines = 2
			case "hr":
				if newLines < 2 {
					fmt.Fprintln(out)
				}
				fmt.Fprintln(out, "---")
				newLines = 1
			case "td", "th":
				//fmt.Fprint(out, " ")
				newLines = 0
			case "a":
				if hasAttr {
					var key, val []byte
					more := true
					for more == true {
						key, val, more = z.TagAttr()
						if string(key) == "href" {
							url = string(val)
						}
					}
				}
				newLines = 0
			case "img":
				if hasAttr {
					var key, val []byte
					more := true
					src := ""
					alt := ""
					for more == true {
						key, val, more = z.TagAttr()
						if string(key) == "src" {
							src = string(val)
						} else if string(key) == "alt" {
							alt = string(val)
						}
					}
					s := " "
					if newLines > 0 {
						s = ""
					}
					if alt != "" {
						fmt.Fprintf(out, "%s%s", s, alt)
						if src != "" {
							fmt.Fprintf(out, " (%s)", src)
						}
					}
				}
				newLines = 0
			}
		case html.EndTagToken:
			tag, _ := z.TagName()
			switch string(tag) {
			case "head":
				reallyIgnore = false
			case "script", "title", "link", "meta", "applet", "embed", "frame", "frameset", "iframe", "object", "style":
				ignore = false
			case "ul", "ol", "pre", "table", "blockquote", "h1", "h2", "h3", "h4", "h5", "h6":
				if newLines < 2 {
					fmt.Fprintln(out, "\n")
					newLines += 2
				}
			case "hr":
				if newLines < 2 {
					fmt.Fprintln(out)
					newLines++
				}
			}
		case html.SelfClosingTagToken:
			tag, hasAttr := z.TagName()
			switch string(tag) {
			case "div", "li", "br", "p":
				if newLines < 2 {
					fmt.Fprintln(out)
					newLines++
				}
			case "hr":
				if newLines < 2 {
					fmt.Fprintln(out, "\n")
				}
				fmt.Fprintln(out, "---\n")
				newLines = 2
			case "a":
				if hasAttr {
					var key, val []byte
					more := true
					for more == true {
						key, val, more = z.TagAttr()
						if string(key) == "href" {
							url = string(val)
						}
					}
				}
				newLines = 0
			case "img":
				if hasAttr {
					var key, val []byte
					more := true
					src := ""
					alt := ""
					for more == true {
						key, val, more = z.TagAttr()
						if string(key) == "src" {
							src = string(val)
						} else if string(key) == "alt" {
							alt = string(val)
						}
					}
					s := " "
					if newLines > 0 {
						s = ""
					}
					if alt != "" {
						fmt.Fprintf(out, "%s%s", s, alt)
						if src != "" {
							fmt.Fprintf(out, " (%s)", src)
						}
					}
				}
				newLines = 0
			}
		case html.DoctypeToken:
		case html.CommentToken:
		default:
			return fmt.Errorf("Unhandled token type: %s", tt.String())
		}
		lastTT = tt
	}

	if lastTT == html.TextToken && url != "" {
		fmt.Fprintf(out, " (%s)", url)
	}

	return nil
}
