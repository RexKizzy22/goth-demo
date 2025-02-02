# Hypermedia Driven Application using HTMX and GO

This is a demo project showcasing the simplicity and efficiency of an application built using the GOTH stack.

GOTH is an acronym for Go, [Templ](https://templ.guide/), and [HTMX](https://htmx.org/). It is often combined with Alpine.js and Tailwind CSS
for client-side state management and rapid UI styling. Alpine.js can be swapped with other lightweight JavaScript
libraries or Web Components while Tailwind can be swapped with Vanilla CSS or other styling libraries.

[Alpine.js](https://alpinejs.dev/) is a lightweight JavaScript library that is used to handle client-side interactivity such as 
modal toggling, form validation, e.t.c.

Templ is a Go templating library that allows you to build and compose UI components similar to JavaScript frontend
libraries/frameworks like React, Vue, and Svelte.

However, HTMX is at the heart of this stack because it is language agnostic and it introduces a new way to build web
applications that have rich interactivity while drastically reducing the overhead of client-side state management.
This approach called HATEOAS (Hypermedia as the Engine of Application State) as popularized by HTMX enables agile
web development thereby collapsing the barrier between frontend and backend web development. It does this by using
fragments of HTML to communicate between client and server rather than using JSON. For more information, consult the
[HTMX DOCS](https://htmx.org/docs)

## How to Run this App

Easy-peasy. Ensure you have the **Go**, [Air](https://github.com/air-verse/air), and **Make** binaries installed in your machine and run these commands:

```shell
go mod download
make run
```

## Project Set Up

The project set up is an adaptation of [this very helpful guide](https://medium.com/ostinato-rigore/go-htmx-templ-tailwind-complete-project-setup-hot-reloading-2ca1ba6c28be) as I didn't make use of Tailwind CSS here.
