package striphtml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	maxInterations = 1000
	text           = `<div class="sect1" title="Validating Your Documents"><div class="titlepage"><div><div><h2 class="title" style="clear: both" id="validating_your_documents">Validating Your Documents</h2></div></div></div><p>One way that professional web developers catch errors in their markup is to validate their documents. What does that mean? To <span class="strong"><strong>validate</strong></span> a document is to check your markup to make sure that you have abided by all the rules of whatever version of HTML you are using (there are more than one, as we’ll discuss in <a class="xref" href="ch10.html" title="Chapter 10. What’s Up, HTML5?">Chapter 10</a>). Documents that are error-free are said to be valid. It is strongly recommended that you validate your documents, especially for professional sites. Valid documents are more consistent on a variety of browsers, they display more quickly, and they are more accessible.</p><p><a id="iddle1331" class="indexterm"></a><a id="iddle1542" class="indexterm"></a><a id="iddle1687" class="indexterm"></a><a id="iddle1939" class="indexterm"></a><a id="iddle2682" class="indexterm"></a><a id="iddle3040" class="indexterm"></a><a id="iddle3092" class="indexterm"></a>Right now, browsers don’t require documents to be valid (in other words, they’ll do their best to display them, errors and all), but any time you stray from the standard you introduce unpredictability in the way the page is displayed or handled by alternative devices.</p><p>So how do you make sure your document is valid? You could check it yourself or ask a friend, but humans make mistakes, and you aren’t really expected to memorize every minute rule in the specifications. Instead, you use a <span class="strong"><strong>validator</strong></span>, software that checks your source against the HTML version you specify. These are some of the things validators check for:</p><div class="itemizedlist"><ul class="itemizedlist"><li class="listitem"><p>The inclusion of a DOCTYPE declaration. Without it the validator doesn’t know which version of HTML or XHTML to validate against.</p></li><li class="listitem"><p>An indication of the character encoding for the document.</p></li><li class="listitem"><p>The inclusion of required rules and attributes.</p></li><li class="listitem"><p>Non-standard elements.</p></li><li class="listitem"><p>Mismatched tags.</p></li><li class="listitem"><p>Nesting errors.</p></li><li class="listitem"><p>Typos and other minor errors.</p></li></ul></div><p>Developers use a number of helpful tools for checking and correcting errors in HTML documents. The W3C offers a free online validator at <span class="emphasis"><em><a class="ulink" href="http://validator.w3.org" target="_top" onmousedown="dataLayer.push({'event': 'eventTracker', 'eventCat': 'Outbound Links', 'eventAct':'Non-Oreilly', 'eventLbl': 'validator.w3.org/', 'eventVal': 0, 'nonInteraction': 0});">validator.w3.org</a></em></span>. For HTML5 documents, use the online validator located at <span class="emphasis"><em>html5.validator.nu</em></span>. Browser developer tools like the Firebug plug-in for Firefox or the built-in developer tools in Safari and Chrome also have validators so you can check your work on the fly. If you use Dreamweaver to create your sites, there is a validator built into that as well.</p><div class="sidebar"><a id="having_problemsquestion_mark"></a><div class="sidebar-title">Having Problems?</div><p>The following are some typical problems that crop up when you are creating web pages and viewing them in a browser:</p><div class="variablelist"><dl><dt><span class="term"><span class="emphasis"><em>I’ve changed my document, but when I reload the page in my browser, it looks exactly the same.</em></span></span></dt><dd><p>It could be you didn’t save your document before reloading, or you may have saved it in a different directory.</p></dd><dt><span class="term"><span class="emphasis"><em>Half my page disappeared.</em></span></span></dt><dd><p>This could happen if you are missing a closing bracket (<span class="strong"><strong><code class="literal">&gt;</code></strong></span>) or a quotation mark within a tag. This is a common error when you’re writing HTML by hand.</p></dd><dt><span class="term"><span class="emphasis"><em>I put in a graphic using the</em></span> <span class="strong"><strong><code class="literal">img</code></strong></span> <span class="emphasis"><em>element, but all that shows up is a broken image icon.</em></span></span></dt><dd><p>The broken graphic could mean a couple of things. First, it might mean that the browser is not finding the graphic. Make sure that the URL to the image file is correct. (We’ll discuss URLs further in <a class="xref" href="ch06.html" title="Chapter 6. Adding Links">Chapter 6</a>.) Make sure that the image file is actually in the directory you’ve specified. If the file is there, make sure it is in one of the formats that web browsers can display (GIF, JPEG, or PNG) and that it is named with the proper suffix (<span class="emphasis"><em>.gif</em></span>, <span class="emphasis"><em>.jpeg</em></span> or <span class="emphasis"><em>.jpg</em></span>, or <span class="emphasis"><em>.png</em></span>, respectively).</p></dd></dl></div></div></div>`
	expected       = `Validating Your DocumentsOne way that professional web developers catch errors in their markup is to validate their documents. What does that mean? To validate a document is to check your markup to make sure that you have abided by all the rules of whatever version of HTML you are using (there are more than one, as we’ll discuss in Chapter 10). Documents that are error-free are said to be valid. It is strongly recommended that you validate your documents, especially for professional sites. Valid documents are more consistent on a variety of browsers, they display more quickly, and they are more accessible.Right now, browsers don’t require documents to be valid (in other words, they’ll do their best to display them, errors and all), but any time you stray from the standard you introduce unpredictability in the way the page is displayed or handled by alternative devices.So how do you make sure your document is valid? You could check it yourself or ask a friend, but humans make mistakes, and you aren’t really expected to memorize every minute rule in the specifications. Instead, you use a validator, software that checks your source against the HTML version you specify. These are some of the things validators check for:The inclusion of a DOCTYPE declaration. Without it the validator doesn’t know which version of HTML or XHTML to validate against.An indication of the character encoding for the document.The inclusion of required rules and attributes.Non-standard elements.Mismatched tags.Nesting errors.Typos and other minor errors.Developers use a number of helpful tools for checking and correcting errors in HTML documents. The W3C offers a free online validator at validator.w3.org. For HTML5 documents, use the online validator located at html5.validator.nu. Browser developer tools like the Firebug plug-in for Firefox or the built-in developer tools in Safari and Chrome also have validators so you can check your work on the fly. If you use Dreamweaver to create your sites, there is a validator built into that as well.Having Problems?The following are some typical problems that crop up when you are creating web pages and viewing them in a browser:I’ve changed my document, but when I reload the page in my browser, it looks exactly the same.It could be you didn’t save your document before reloading, or you may have saved it in a different directory.Half my page disappeared.This could happen if you are missing a closing bracket (&gt;) or a quotation mark within a tag. This is a common error when you’re writing HTML by hand.I put in a graphic using the img element, but all that shows up is a broken image icon.The broken graphic could mean a couple of things. First, it might mean that the browser is not finding the graphic. Make sure that the URL to the image file is correct. (We’ll discuss URLs further in Chapter 6.) Make sure that the image file is actually in the directory you’ve specified. If the file is there, make sure it is in one of the formats that web browsers can display (GIF, JPEG, or PNG) and that it is named with the proper suffix (.gif, .jpeg or .jpg, or .png, respectively).`
)

// TESTS

func Test_StripHtmlTagsWithRegexpP1(t *testing.T) {
	output := StripHtmlTagsWithRegexpP1(text)
	assert.Equal(t, expected, output)
}

func Test_StripHtmlTagsWithRegexpP2(t *testing.T) {
	output := StripHtmlTagsWithRegexpP2(text)
	assert.Equal(t, expected, output)
}

func Test_StripHtmlTagsWithStringBuilderV1(t *testing.T) {
	output := StripHtmlTagsWithStringBuilderV1(text)
	assert.Equal(t, expected, output)
}

func Test_StripHtmlTagsWithStringBuilderV2(t *testing.T) {
	output := StripHtmlTagsWithStringBuilderV2(text)
	assert.Equal(t, expected, output)
}

func Test_StripHtmlTagsWithBlueMonday(t *testing.T) {
	output := StripHtmlTagsWithBlueMonday(text)
	assert.Equal(t, expected, output)
}

// BENCHMARKS

func Benchmark_StripHtmlTagsWithRegexpP1(b *testing.B) {
	for i := 0; i < maxInterations; i++ {
		StripHtmlTagsWithRegexpP1(text)
	}
}

func Benchmark_StripHtmlTagsWithRegexpP2(b *testing.B) {
	for i := 0; i < maxInterations; i++ {
		StripHtmlTagsWithRegexpP2(text)
	}
}

func Benchmark_StripHtmlTagsWithStringBuilderV1(b *testing.B) {
	for i := 0; i < maxInterations; i++ {
		StripHtmlTagsWithStringBuilderV1(text)
	}
}

func Benchmark_StripHtmlTagsWithStringBuilderV2(b *testing.B) {
	for i := 0; i < maxInterations; i++ {
		StripHtmlTagsWithStringBuilderV2(text)
	}
}

func Benchmark_StripHtmlTagsWithBlueMonday(b *testing.B) {
	for i := 0; i < maxInterations; i++ {
		StripHtmlTagsWithBlueMonday(text)
	}
}
