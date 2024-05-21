# Spruce

> Before you continue, please know that i am not well versed in writing parsers/languages. This is my first attempt at writing a parser. I am just a beginner. I am writing this parser for my own learning purposes. If you are looking for a production ready parser, this is not it.

Spruce is a simple html transpiler made for creating simple quick static websites.
The purpose of this project is to be able to make a templating language that is simple to use and understand while having features to easily create pages without having to change other files to make changes.

# Goal

The goal of this project is to make a simple to understand (with basic html/css knowledge) and easy/short to write templating language in order to make static sites.
In order to reach these goals there are 3 things i want to fix in HTML:
- No opening/closing tags, instead just use a scope.
- Templates/Slots, dont repeat yourself anymore and keep code/files shorter.
- Add better ways to add classes, stop having a single line that has over 50 characters because you have to add all the classes.

Ideally the language would look something like this:
```
@import "layout.spr"
@import "classes.spr"

@meta {
  url = "/subjects"
  tags = []
  title = "All Subjects"
}

url = {
  a.link-class {
    href=$href

    $@
  }
}

layout {
  h1#title {
    class=classes.title + "color-blue"
    
    meta.title
  }

  main {
    ul {
      for page in @pages.tags(["subject"]) {
        li {
          url href=page.url {
            page.title
          }
        }
      }
    }
  }
}
```

# Features/Roadmap
- [x] HTML transpiling
- [ ] Variables
- [x] IDs, Classes and Attributes
- [ ] Loops
- [ ] Templates
- [ ] Conditionals
- [ ] Operators
- [ ] Keyword safety
- [ ] Logging
- [ ] Comments
- [ ] Script/JS support
- [ ] Style support
- [ ] CSS support
- [ ] SCSS support
- [ ] Media compression
- [ ] Automatic SEO
- [ ] Automatic SiteMap
- [ ] Automatic RSS
- [ ] Automatic Accessibility Features
- [ ] Language support (Vim, VSCode, etc.)
- [ ] Documentation
- [ ] Integrations (Docker, Github Actions, etc.)
