# Pantopoda

Pantapoda is a simple web crawler/spider written in Go, as requested by the
DigitalOcean exercise spec.

It's fairly minimalistic but and no dependencies at this point, the crawl index
must fit in memory.

### Usage

Example:
    go run main/main.go -out ./out.json example.com

See
    go run main/main.go -h for more options.

### Caveats / TODO's
- 1) robots.txt is currently gnored whihc is bad. We could leverage
https://github.com/temoto/robotstxt-go in fetcher.AllowUrl
- 2) Response status: Currently ignored, but if it's 40x, 50x and such we should
probably skip parsing the body.
- 3) If a link specifies rel = "nofollow" we should respect that.
- 4) We should also deal with links "canonical name" if specified ?
- 5) We should also respect robots meta tags such as: meta name="robots" content="index, follow"

There are plenty of other pitfalls and caveats we could try to deal with ...

### Spec: (as specificed by DigitalOcean.)

> > We'd like you to write a web Fetcher in a modern language (something like
> > Ruby, Python, Go, or Coffeescript). It should be limited to one domain - so
> > when crawling digitalocean.com it would crawl all pages within the
> > digitalocean.com domain, but not follow the links to our Facebook or
> > Instagram accounts or subdomains like cloud.digitalocean.com. Given a
> > URL, it should output a site map, showing which static assets each page
> > depends on, and the links between pages. Choose the most appropriate data
> > structure to store & display this site map.
> >
> > Build this as you would build something for production - it's fine if you
> > don't finish everything, but focus on code quality and write tests. We're
> > interested in how you code and how you test your code.

