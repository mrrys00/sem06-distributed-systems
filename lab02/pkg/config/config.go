package config

const (
	HtmlPage = "<!DOCTYPE html>\n<html>\n<body>\n\n<p>Show information about city</p>\n\n" +
		"<form action=\"http://localhost:8080/weather\" id=\"frm1\" method=\"get\">\n" +
		"    City: <input name=\"city\" type=\"text\"><br><br>\n    API Forecast: <input name=\"forecast\" type=\"text\"><br><br>\n" +
		"    API UX Index: <input name=\"index\" type=\"text\"><br><br>\n" +
		"    <input onclick=\"myFunction()\" type=\"button\" value=\"Submit\">\n</form>\n\n<script>\n" +
		"    function myFunction() {\n        document.getElementById(\"frm1\").submit();\n" +
		"    }\n</script>\n\n</body>\n</html>"

	WeatherPath = "/weather"
	DefaultPath = "/"
	Localhost   = "localhost:8080"
	M3OURL      = "https://api.m3o.com/v1/weather/Forecast"
	UVURL       = "https://api.openuv.io/api/v1/forecast"
)
