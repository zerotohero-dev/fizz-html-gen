/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package conf

const InlineStyleToInject = `
<style type="text/css">
/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))';;,,               <zerotohero.dev>
 *   ,(@') ;)'))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;',,/7),)) )) )\,,
 *  (& )'   (,((,((;( ))\,
 */

html, body {padding: 0;margin:0;}
p {margin-top: 1.4em;margin-bottom: 1.4em;}
html body {cursor: default;font-family: courier-prime, monospace;font-size: 16px; 
  background: #073642;-webkit-text-size-adjust: none;text-size-adjust: none;}
pre, span, div {font-family: courier-prime, monospace;font-size: 16px;}
a {color: #2D37B1;padding: 4px 2px 2px 4px;transition: all 250ms ease-out;}
a:hover {color: #000000; background: #FFDB77;}
hr {max-width: 64px;margin: 0;border: 0;height: 2px;background: #CEC6B0;}
pre {line-height: 1.44;background: #fdf6e3;padding: 1em;margin: 6em auto;
  margin-bottom: 1em;border-radius: 4px;max-width: 795px; 
  border-top: 16px #FFDB77 solid;overflow: auto;tab-size: 4;}
div.preformatted {color:#000000;padding:1em;line-height: 1.64;
  background: #fdf6e3;padding: 1em;margin: 6em auto;
  margin-bottom: 1em;border-radius: 4px;
  max-width: 795px;border-top: 16px #FFDB77 solid;
  overflow: auto;tab-size: 4;}
div.preformatted pre p {margin: 0;padding: 0;}
div.preformatted code {color: #008800;background: #FFF0C6;
  border: 1px #FFDB77 solid;padding: 2px;}
div.preformatted pre code {border: 0;padding: 0;color: #382900;
  background: #FFF0C6;}
div.preformatted hr {margin-top: 2em;margin-bottom: 2em;}
div.preformatted pre {border-top: 0;color:#000000;background:#FFF0C6;
  padding:1em;margin: 1em;border: 1px #FFDB77 solid;}
div.preformatted a {display: inline;padding: 4px 2px 2px 4px;}
div.preformatted ul {margin:0;padding-left:1em;}
div.preformatted ol {margin:0;padding-left:2em;}
div.preformatted h2:before {content: '';margin-bottom: 1rem;
  width: 64px;height: 2px;display: block;background: #CEC6B0;}
div.preformatted h2 {font-weight: normal;font-size: 1.4em;
  display: inline-block;color: #d33682;margin-top: 2.2em;margin-bottom: 0.4em;}
nav span.ref {color:#000000;background:#CEC6B0;padding:4px;}
nav strong.heading {font-weight: normal;font-size: 1.4em;display: inline-block;
  color: #d33682;margin-top: 1.2em;}
nav a.inline {display: inline;padding: 4px;margin: 0;}
.btn-cmt {margin: 0.5em auto;display: block;background: #FFF0C6;
  color: #293084;text-align: center;
  margin-top: -1em;margin-left: -1em;margin-right: -1em;}
.btn-cmt:hover {background: #7A59AE;color: #ffffff;}
.copyright {font-size: 14px;padding: 2em;line-height: 2em;padding-top: 4em;
  padding-bottom: 2em;text-align: center;font-family: monospace;
  color:  #839496;}
.copyright a {padding: 0;color: #839496;display: inline;}
.copyright a:hover {background: inherit;}

td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
body .hll { background-color: #ffffcc }
body { background: #ffffff; }
body .c { color: #008800; font-style: italic } /* Comment */
body .err { color: #a61717; background-color: #e3d2d2 } /* Error */
body .k { color: #0000ff; font-weight: normal; } /* Keyword */
body .ch { color: #008800; font-style: italic } /* Comment.Hashbang */
body .cm { color: #008800; font-style: italic } /* Comment.Multiline */
body .cp { color: #008080 } /* Comment.Preproc */
body .cpf { color: #008800; font-style: italic } /* Comment.PreprocFile */

body .c1 { 
	color: #008800; font-style: normal; 
  padding: 4px;
  display: inline-block;
	cursor: default;
	border-left: 6px #FFD256 solid;
	transition: all 250ms ease-out;
} /* Comment.Single */

/* Syntax highlight below */

body .cs { color: #008800; font-weight: normal } /* Comment.Special */
body .gd { color: #000000; background-color: #ffdddd } /* Generic.Deleted */
body .ge { font-style: italic } /* Generic.Emph */
body .gr { color: #aa0000 } /* Generic.Error */
body .gh { color: #999999 } /* Generic.Heading */
body .gi { color: #000000; background-color: #ddffdd } /* Generic.Inserted */
body .go { color: #888888 } /* Generic.Output */
body .gp { color: #555555 } /* Generic.Prompt */
body .gs { font-weight: bold } /* Generic.Strong */
body .gu { color: #aaaaaa } /* Generic.Subheading */
body .gt { color: #aa0000 } /* Generic.Traceback */
body .kc { color: #0000ff; font-weight: normal } /* Keyword.Constant */
body .kd { color: #CC00A3; font-weight: normal } /* Keyword.Declaration */
body .kn { color: #0000ff; font-weight: normal } /* Keyword.Namespace */
body .kp { color: #0000ff; ont-weight: normal } /* Keyword.Pseudo */
body .kr { color: #0000ff; font-weight: normal } /* Keyword.Reserved */
body .kt { color: #0000ff; font-weight: normal; } /* Keyword.Type */
body .m { color: #0000FF } /* Literal.Number */
body .s { color: #0000FF } /* Literal.String */
body .na { color: #FF0000 } /* Name.Attribute */
body .nt { color: #000080; font-weight: bold } /* Name.Tag */
body .ow { font-weight: bold } /* Operator.Word */
body .w { color: #bbbbbb } /* Text.Whitespace */
body .mb { color: #0000FF } /* Literal.Number.Bin */
body .mf { color: #0000FF } /* Literal.Number.Float */
body .mh { color: #0000FF } /* Literal.Number.Hex */
body .mi { color: #0000FF } /* Literal.Number.Integer */
body .mo { color: #0000FF } /* Literal.Number.Oct */
body .sa { color: #0000FF } /* Literal.String.Affix */
body .sb { color: #0000FF } /* Literal.String.Backtick */
body .sc { color: #800080 } /* Literal.String.Char */
body .dl { color: #0000FF } /* Literal.String.Delimiter */
body .sd { color: #0000FF } /* Literal.String.Doc */
body .s2 { color: #0000FF } /* Literal.String.Double */
body .se { color: #0000FF } /* Literal.String.Escape */
body .sh { color: #0000FF } /* Literal.String.Heredoc */
body .si { color: #0000FF } /* Literal.String.Interpol */
body .sx { color: #0000FF } /* Literal.String.Other */
body .sr { color: #0000FF } /* Literal.String.Regex */
body .s1 { color: #0000FF } /* Literal.String.Single */
body .ss { color: #0000FF } /* Literal.String.Symbol */
body .il { color: #0000FF } /* Literal.Number.Integer.Long */
</style>
`