// Error page template. Implements BasePage methods.

{% code
type ErrorPage struct {
	// inherit from base page, so its' title is used in error page.
	BasePage

	// error path
	Path []byte
}
%}


{% func (p *ErrorPage) Body() %}
	<h1>Error page</h1>
	</div>
		Unsupported path <b>{%z p.Path %}</b>.
	</div>
	Base page body: {%= p.BasePage.Body() %}
{% endfunc %}