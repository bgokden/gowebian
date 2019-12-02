package mdb

import (
	"github.com/bgokden/gowebian/pkg/page"
)

type Page struct {
	page.BasePage
}

func NewPage() *Page {
	pg := &Page{}
	page.InitDefaults(pg)
	return pg
}

func (p *Page) Render() string {
	return `
  <!doctype html>
  <html lang="en">
  	<head>
    {{ range $key, $value := .GetHeaders }}
      {{ Generate $value }}
    {{ end }}
		<!-- Font Awesome -->
		<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.2/css/all.css">
		<!-- Bootstrap core CSS -->
		<link href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
		<!-- Material Design Bootstrap -->
		<link href="https://cdnjs.cloudflare.com/ajax/libs/mdbootstrap/4.8.11/css/mdb.min.css" rel="stylesheet">
  	</head>
  	<body {{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>
		{{if .HasChildren }}
	    {{ range $key, $value := .GetChildren }}
	      {{ Generate $value }}
	    {{ end }}
		{{else}}
			<div style="height: 100vh">
				<div class="flex-center flex-column">
					<h1 class="animated fadeIn mb-2">Material Design for Bootstrap</h1>

					<h5 class="animated fadeIn mb-1">Thank you for using our product. We're glad you're with us.</h5>

					<p class="animated fadeIn text-muted">MDB Team</p>
				</div>
			</div>
		{{end}}
    <!-- SCRIPTS -->
    <!-- JQuery -->
		<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
		<!-- Bootstrap tooltips -->
		<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.4/umd/popper.min.js"></script>
		<!-- Bootstrap core JavaScript -->
		<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.3.1/js/bootstrap.min.js"></script>
		<!-- MDB core JavaScript -->
		<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/mdbootstrap/4.8.11/js/mdb.min.js"></script>
  	</body>
  </html>
  `
}
