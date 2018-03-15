import "html/template"

var issureList = template.Must(template.New("issurelist").Parse(`
<h1>{{.TotalCount}} issues </h1>
<table>
<tr style='text-align:left'>
    <th>State</th>

<
`))
