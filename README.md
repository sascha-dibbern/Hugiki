# What is Hugiki ?

Hugiki is a WYSIWYW (what you see, is what you __want__) editing frontend for creating Hugo-based websites.
Using Hugiki helps to reduce page-editing turnaround time when building Hugo-based webpages and provide a better "look and feel" experience as the current Hugo-approach.

## What is Hugo ?

Hugo is a fast building system for fast static websites. One can say, that Hugo transforms markdown-text content, other content and templates into a complete static website.

See more about Hugo
* [Hugo website](https://gohugo.io/)
* [Hugo on Github](https://github.com/gohugoio/hugo) 

## What is the problem in build websites with Hugo only ?
Hugo, as a static markdown-text based website builder, might not be complete userfriendly solution for many users, as the website construction process for them

1. requires knowledge of markdown language and has a steeper learning curve than just "clicking" in a graphical user interface
1. assumes the user has a good imagination and experience of how the markdown-text will look like in the final website ie. no "look and feel"
1. leads to many interations of editing & saving markdown-text documents and reloading the generated webpage in a separate webbrowser

All these issues are in other (dynamic) website building solutions such as CMS (content management systems) addressed by a WYSIWYG-approach (what you see, is what you __get__). Unfortunately these CMS' are known to have deficiencies in other areas such as the final website's stability, security and performance.

## How Hugiki helps Hugo ?
Hugiki brings the element of faster reaction to changes of a Hugo website by allowing the user to edit the markdown-text, react to changes in it and show the resulting Hugo-generated webpage within the same browser-view of a Hugiki-webpage.

## What does "Hugiki" mean ?
The word Hugiki derives from
* Hugo
* Wiki - the Hawaiian word for quick (see [Wikipedia](https://en.wikipedia.org/wiki/Wiki))

# The current state of Hugiki

Hugiki is currently in a explorative development phase. and cna therefore be seen still as experimental

# Feature ideas
A list of features ideas that are technical viable to be implemented in a shorter term

1. Hugo project files exploration
1. Online configuration settings
1. Basic git integration for managing
1. Editing other files than the Hugo content files

# Using Hugiki

## Install directly

```text
go install github.com/sascha-dibbern/Hugiki@latest
```

## Install directly

```text
go install github.com/sascha-dibbern/Hugiki@latest
```

## Clone and build yourself

```text
git clone git@github.com:sascha-dibbern/Hugiki.git
cd Hugiki
go build
```

# Running Hugiki

## Create a minimal configuration file

Create a configuration file `myproject.yml` in Yaml format for a Hugo project under `/home/user/projects/demosite`:

```text
backendBaseUrl: http://localhost:1313/
hugoProject: /home/user/projects/demosite
```


## Running Hugiki

### Start Hugo
First start Hugo's development server to view the enable site.
```text
cd /home/user/projects/demosite
hugo server
```

or for enabling viewing and editing draft documents

```text
cd /home/user/projects/demosite
hugo server -D
```
(See also more at [Hugo quick start](https://gohugo.io/getting-started/quick-start/)

Hugo's development server will provide content under http://localhost:1313/ that Hugiki will proxy to.

### Start Hugiki

Next start Hugiki
```text
Hugiki --config myproject.yml
```

Hugiki will provide a webserver under http://localhost:3000/

!!! Currently content can only be edited by entering the content path as part of the Hugiki URI such as

```text
http://localhost:3000/hugiki/page/edit/content/hello/
```

for the file `/home/user/projects/demosite/content/hello.md`
