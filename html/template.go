// Code generated by go generate; DO NOT EDIT.
package html

const pageTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    {{range .Metas}}<meta{{if .Name}} name="{{.Name}}"{{end}}{{if .HTTPEquiv}} http-equiv="{{.HTTPEquiv}}"{{end}}{{if .Content}} content="{{.Content}}"{{end}}>
    {{end}} 
    <title>{{.Title}}</title>
    {{if .AppStyle}}<style media="all" type="text/css">
        html {
            height: 100%;
            width: 100%;
            margin: 0;
        }
        
        body {
            height: 100%;
            width: 100%;
            margin: 0;
        }
    </style>{{end}}
    {{range .CSS}}
    <link type="text/css" rel="stylesheet" href="{{.}}">{{end}}
</head>
<body>
    <div data-goapp-root>{{if .DefaultComponent}}
        {{.DefaultComponent}}{{end}}
    </div>

    <script> {{if .AppJS}}
       {{.AppJS}}{{end}}
    </script>
    {{range .Javascripts}}
    <script src="{{.}}"></script>{{end}}
</body>
</html>
`
