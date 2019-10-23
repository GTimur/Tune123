{{define "main"}}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{template "title" .}}</title>
    <script src="/static/js/jquery-3.3.1.min.js"></script>    
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
    <script src="/static/js/bootstrap.min.js"></script>
    <!-- Include Fancytree skin and library -->
    <link href="/static/fancytree/dist/skin-win8/ui.fancytree.min.css" rel="stylesheet">
    <script src="/static/fancytree/dist/jquery.fancytree-all-deps.min.js"></script>
  {{template "head"}}
</head>
<body>
{{template "body" .}}        
{{template "scripts"}}
</body>
</html>
{{end}}
