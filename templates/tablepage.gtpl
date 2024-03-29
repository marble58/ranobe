// Table page template. Implements BasePage methods.

{% code
type TablePage struct {
	Rows []string
}
%}


{% func (p *TablePage) Title() %}
	This is table page
{% endfunc %}


{% func (p *TablePage) Body() %}
	<h1>Table page</h1>

	{%= p.form() %}

	{% if len(p.Rows) == 0 %}
		No rows. Click <a href="/table?rowsCount=5">here</a>.
	{% else %}
		<table>
			{%= emitRows(p.Rows) %}
		</table>
	{% endif %}
{% endfunc %}

{% func emitRows(rows []string) %}
	<tr>
		<th>#</th>
		<th>value</th>
	</tr>

	{% for n, r := range rows %}
		{% if r == "bingo" %}
			<tr><td colspan="2"><h1>BINGO!</h1></td></tr>
			{% return %}
		{% elseif n == 42 %}
			<tr><td colspan="2">42 rows already generated</td></tr>
			{% break %}
		{% endif %}

		<tr style="background: {% if n&1 == 1 %}white{% else %}#ddd{% endif %}">
			<td>{%d n+1 %}</td>
			<td>{%s r %}</td>
		</tr>
	{% endfor %}

	<tr><td colspan="2">No bingo found</td></tr>
{% endfunc %}

{% func (p *TablePage) form() %}
	<form>
		Rows: <input type="text" name="rowsCount" value="{%d len(p.Rows) %}"/><br/>
		<input type="submit" value="Generate!"/>
	</form>
{% endfunc %}