package server

import "net/http"

// Contains the data extracted from the request and returned to the user.
type RequestResponse struct {
	// IP address of the remote client
	RemoteAddr string `json:"remoteAddress"`
	// Port of the remote client
	RemotePort string `json:"remotePort"`
	// Host header
	Host string `json:"host"`
	// Request method
	Method string `json:"method"`
	// Request protocol
	Proto string `json:"protocol"`
	// Request headers
	Header http.Header `json:"header"`
}

// Returns the html template string for the index page.
//
//nolint:funlen
func IndexTemplateString() string {
	cssStyle := `<style type="text/css">
    body {
      background-color: #121212;
      color: #e0e0e0;
      font-family: Verdana, Helvetica, Arial, sans-serif;
      font-size: 15px;
      text-align: center;
    }

    a {
      color: #81d4fa;
      /* Light blue for links */
      text-decoration: none;
    }

    a:hover {
      text-decoration: underline;
    }

    input,
    button,
    select {
      background-color: #333;
      color: #fff;
      border: 1px solid #555;
    }

    table {
      background-color: #121212;
      border: 0;
      overflow-wrap: anywhere;
      width: 15%;
      min-width: fit-content;
      margin-left: auto;
      margin-right: auto;
    }

    tr {
      background-color: #333;
      text-align: center;
    }

    td {
      font-size: 15px;
      text-align: center;
    }

    .title {
      font-size: 18px;
      font-weight: bold;
      color: white
    }

    .small {
      font-size: 11px;
    }

    .big {
      font-size: 22px;
    }
    .first {
      padding-top: 0.5em;
      padding-left: 0.5em;
      padding-right: 0.5em;
    }
    .last {
      padding-bottom: 0.5em;
      padding-left: 0.5em;
      padding-right: 0.5em;
    }

    .center {
      text-align: center;
    }
    
    .left {
      text-align: left;
    }

    :link,
    :visited {
      text-decoration: none;
      color: #81d4fa
    }
  </style>`

	javaScript := `<script>
    function genLink() {
      const port = document.getElementById("portInput").value;
      const linkTextV4 = "http://127.0.0.1:"+port;
      const linkTextV6 = "http://[::1]:"+port;
      document.getElementById("linkResult").innerHTML =
        '<a href="'+linkTextV4+'" target="_blank" style="color: #00db54;">'+linkTextV4+'</a><br>'+
        '<a href="${linkTextV6}" target="_blank" style="color: #00db54;">'+linkTextV6+'</a>';
    }
    function copyText(elementID) {
      // Get the text field
      var copyText = document.getElementById(elementID);
      if (window.isSecureContext) {
        // Copy the text inside the text field
        navigator.clipboard.writeText(copyText.innerText);
      } else {
        console.log("can't copy; insecure context")
      }
      // Log the copied text
      console.log(elementID + ":", copyText.innerText);
    }

    function checkContext() {
      if (!window.isSecureContext) {
        document.querySelectorAll('.copyButt').forEach(function(el) {
          el.style.display = 'none';
        });
      }

      genLink();
    }
    window.onload = checkContext;
  </script>`

	head := `<title>locallhost</title>
  <meta name="keywords" content="localhost, software, ip address, ip addresses, ip, http header, golang, host, tools">
  <meta charset="UTF-8">
  <meta name="description" content="locallhost linker">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  ` +
		cssStyle +
		javaScript

	body := `<body class="center">
  <main>
    <span class="small">Perhaps you intended to go here:</span><br>
    <span class="big">
      <a href="http://127.0.0.1:80/">http://127.0.0.1:80</a>
      <br>
      <a href="http://[::1]:80/">http://[::1]:80</a>
    </span>
    <br><br>

    <span class="small">Or perhaps here:</span><br>
    <div class="big">
      <input type="number" id="portInput" value="8080" onload="genLink()" onchange="genLink()">
      <br>
      <div class="result" id="linkResult"></div>
    </div>
    <br>
    <br>

    <table>
      <tr>
        <td style="background-color: #008030;" class="title">
          Your IP Address
        </td>
      </tr>
      <tr>
        <td class='first last'>
          <span class="big" id="ipAddr">{{ .RemoteAddr }}</span>
          &nbsp;
          <button class="copyButt" onclick="copyText('ipAddr')">Copy</button>
          <br>
          <span id="port">Port: {{ .RemotePort }}</span>
          &nbsp;
          <button class="copyButt" onclick="copyText('port')">Copy</button>
        </td>
      </tr>
    </table>
    <br><br>

    <table>
      <tr>
        <td style="background-color: #0040b0;" class="title">
          Your HTTP Request Header &nbsp;
          <button class="copyButt" onclick="copyText('headers')">Copy</button>
        </td>
      </tr>
      <tr>
        <td class="last left" id="headers">
          <span>
            {{ .Method }} / {{ .Proto }} <br><br>
            Host: {{ .Host }}<br>
            {{- if  index .Header "User-Agent" }}
            User-Agent: {{ index (index .Header "User-Agent") 0 }}<br>
            {{- end }}
            <br>
            {{- range $key, $value := .Header }}
            {{- if eq $key "User-Agent" }}
            {{- continue }}
            {{- end }}
            {{- if gt (len $value) 1 }}
            {{ $key }}: {{ $value }}<br>
            {{- else }}
            {{ $key }}: {{ index $value 0}}<br>
            {{- end }}
            {{- end }}
          </span>
        </td>
      </tr>
    </table>
    <br><br>
    <div>
    Addt'l Endpoints
    <li><a href="/ip">/ip</a></li>
    <li><a href="/json">/json</a></li>
    </div>
  </main>
  <!-- footer -->
  <footer>
    <p>
      <a href="https://b13rg.icecdn.tech/">B13rg</a> - <a href="https://github.com/B13rg/locallhost">Source</a>
      <br>
      Inspired by <a href="http://locallhost.com">locallhost.com</a>
      <br>
    </p>
  </footer>
</body>`

	return `<!DOCTYPE html>
<!-- Based on http://locallhost.com/ - &copy; 2004-2025 Tom Anderson-->
<html> 
<head>` +
		head +
		`</head>` +
		body +
		`</html>`
}
