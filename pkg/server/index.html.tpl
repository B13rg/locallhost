<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">

<!-- Based on https://locolhost.com/ -->
<!-- &copy; 2004-2025 Tom Anderson-->

<html>

<head>
  <title>localhost</title>
  <meta name="keywords" content="localhost, software, ip address, ip addresses, ip, http header, golang, host, tools">
  <style type="text/css">
    body {
      background-color: #121212;
      color: #e0e0e0;
      font-family: Arial, sans-serif;
    }

    a {
      color: #81d4fa; /* Light blue for links */
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
    <!--
    p {
      font-size: 15px;
      font-family: verdana, arial, helvetica, sans-serif
    }

    td {
      font-size: 15px;
      font-family: verdana, arial, helvetica, sans-serif
    }

    .title {
      font-size: 18px;
      font-family: verdana, arial, helvetica, sans-serif;
      font-weight: bold;
      color: white
    }

    .small {
      font-size: 10px;
      font-family: verdana, arial, helvetica, sans-serif
    }

    .big {
      font-size: 22px;
      font-family: verdana, arial, helvetica, sans-serif;
    }

    .notices {
      font-size: 10px;
      font-family: arial, helvetica, sans-serif;
      color: gray
    }

    :link,
    :visited {
      text-decoration: none;
      color: #0040b0
    }
    -->
  </style>
</head>

<body bgcolor="#ffffff">

  <center>
    <table border=0 cellspacing=0 cellpadding=0 width="98%" height="98%">
      <tr height=60 align=center valign=top>
        <td align="left" width="336px" rowspan="2">
        </td>
        <td align="center">
        </td>
        <td align="right" width="336px" rowspan="2">
        <td />
      </tr>

      <tr align="middle">
        <td align="center">
          <span class="small">Perhaps you intended to go here:</span><br>
          <span class="big">
            <a href="http://127.0.0.1:80/">http://127.0.0.1:80</a>
          </span>
          <span class="big">
            <a href="http://127.0.0.1:8080/">http://127.0.0.1:8080</a>
          </span>
        </td>
      </tr>
      <tr>
        <td align="center">
          <span class="small">Perhaps Here:</span><br>
          <span class="big">
            http://127.0.0.1:
            <input type="number" id="portInput" placeholder="8080">
            <button onclick="genLink()">Gen</button>
          </span>
          <span class="big"><div class="result" id="linkResult"></div></span><br>
        </td>

        

    <script>
        function genLink() {
            const port = document.getElementById('portInput').value;
            const linkText = `http://127.0.0.1:${port}`;
            document.getElementById('linkResult').innerHTML = 
                `<a href="${linkText}" target="_blank" style="color: #4CAF50; text-decoration: none;">${linkText}</a>`;
        }
    </script>  
      </tr>

      <tr valign=middle align=center>
        <td colspan="3" align=center nowrap>

          <table bgcolor="#008030" border=0 cellpadding=1 cellspacing=0>
            <tr>
              <td>

                <table border=0 cellpadding=3 cellspacing=0 bgcolor="#ffffff">

                  <tr height=25>
                    <td bgcolor="#008030">&nbsp;</td>
                    <td bgcolor="#008030" align=middle height=30 class="title">Your IP Address</td>
                    <td bgcolor="#008030">&nbsp;</td>
                  </tr>

                  <tr>
                    <td width=35 nowrap>&nbsp;</td>
                    <td align=middle>
                      <table border=0 cellpadding=25>
                        <tr>
                          <td align=middle>
                            <span class="big">35.131.63.214</span><br>
                          </td>
                        </tr>
                      </table>
                    </td>
                    <td width=35 nowrap>&nbsp;</td>
                  </tr>

                </table>
              </td>
            </tr>
          </table>

        </td>
      </tr>

      <tr>
        <td colspan="3" align=center>&nbsp;</td>
      </tr>

      <tr>
        <td colspan="3" align="center">

          <table bgcolor="#0040b0" border=0 cellpadding=1 cellspacing=0>
            <tr>
              <td>
                <table border=0 cellpadding=3 cellspacing=0 bgcolor="#ffffff">

                  <tr height=25>
                    <td bgcolor="#0040b0" align="center" height=30 class="title">Your HTTP Request Header</td>
                  </tr>

                  <tr>
                    <td align="left">
                      GET / HTTP/1.1 <br />
                      {.key}: {.value} <br />
                    </td>
                  </tr>

                </table>
              </td>
            </tr>
          </table>

        </td>
      </tr>
      
      <!-- footer -->
      <tr height=50 align=center valign=bottom>
        <td colspan="3" align=center>
        </td>
      </tr>

    </table>
  </center>

</body>

</html>
