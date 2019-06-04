package ui

import (
	"fmt"
	"html/template"
)

// language=JavaScript
const cssFmtString string = `
	(function(css){
	var style = document.createElement('style');
	var head = document.head || document.getElementsByTagName('head')[0];
				style.setAttribute('type', 'text/css');
				if (style.styleSheet) {
					style.styleSheet.cssText = css;
				} else {
					style.appendChild(document.createTextNode(css));
				}
				head.appendChild(style);
	})("%s")`

// language=JavaScript
const htmlFmtString string = `
	(function(parentSelector, html){
	    var parentElement = document.querySelector(parentSelector);
	    parentElement.innerHTML = html
	})("%s", "%s")`

func injectCss(css string) string {
	return fmt.Sprintf(cssFmtString, template.JSEscapeString(css))
}

func injectHtml(parentSelector, html string) string {
	return fmt.Sprintf(htmlFmtString, parentSelector, template.JSEscapeString(html))
}

func injectJs(js string) string {
	return js
}
