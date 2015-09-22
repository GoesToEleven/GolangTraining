package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>

<h1>LOG IN</h1>
<form method="post">
    <h3>User name</h3>
    <input type="text" name="userName" id="userName">
    <h3>Password</h3>
    <input type="text" name="password" id="password">
    <input type="submit">
</form>

</body>
</html>`)
	})

	http.ListenAndServe(":9000", nil)
}
