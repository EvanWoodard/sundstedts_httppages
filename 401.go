package httppages

const page401 string = `
{{ define "401" }}
<!DOCTYPE html>
<html>
<head>
    <title>Sundstedt's - 401</title>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Roboto&display=swap" rel="stylesheet">
    <link href="shared/css/theme.mobile.css" rel="stylesheet" >
    <link href="shared/css/theme.css" rel="stylesheet" >
</head>
<body>
    <div class="topbar">
        <div class="logo"><img src="shared/img/android-chrome-512x512.png" /></div>
        <div class="title">The Sundstedt's</div>
        <div class="spacer"></div>
        <form class="redirect-form" action="https://iam.sundstedt.us/logout" method="POST">
            <input id="home_redirect" type="hidden" name="redirect" value="https://evenson.sundstedt.us/">
            <button class="btn login-btn" type="submit">Home</button>
        </form>
    </div>
    <div class="page-content">
        <div class="http-error">
            <strong class="code">401</strong>
            <div class="message">Uh-Oh! You're not authorized to view this page!</div>
        </div>
    </div>
</body>
</html>
{{ end }}
`
