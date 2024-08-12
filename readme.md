S# Google Analytics Scanner

I wrote this program to check a URL for Google Analytics tracking code. It uses the `requests` library to fetch the HTML content of the URL and then searches for the Google Analytics tracking code in the HTML source.

It checks for the following Google Analytics references:
- `google-analytics.com/ga.js`
- `google-analytics.com/analytics.js`
- `google-analytics.com/ga.js`
- `google-analytics.com/dc.js`
- `google-analytics.com/urchin.js`
- `google-analytics.com/gtag/js`
- `googletagmanager.com/gtag/js`


Example of Google Analytics tracking code in the source:
```html
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-XXXXXX-X"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-XXXXXX-X');
</script>
```

## Usage

To use the program, simply run the `main.go` program and provide the URL you want to check as an argument. For example:

```bash

$ go run main.go https://www.nist.gov

```