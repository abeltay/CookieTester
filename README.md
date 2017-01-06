# Cookie Tester

The purpose of this test is to check whether cookies are set through an AJAX request.

## How to test
```
var xhr = new XMLHttpRequest();
xhr.open('GET', 'http://lvh.me:8000/api/set', true);
xhr.send(null);
```
1. Start the Go server
1. Run the above javascript in the browser. I ran these commands using the console of Google Chrome.
1. Go to `http://lvh.me:8000/get`
1. You should see that the cookie is set for the domain

This shows that the cookie can be set though a XHR. You can edit the cookie params to try out different scenarios.
