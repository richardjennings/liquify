# Liquify

Borrows heavily from MIT Licenced https://github.com/osteele/liquid Where `osteele/liquid` is
designed to render Liquid templating (and is used by GoJekyll) - this usage of their code makes 
the AST more the primary focus, removing rendering (this is the AST transpilers job). In some 
places the modifications are substantial and in other places a copy/paste. @todo document what
is borrowed.



## About

Parse Jekyll front matter and Liquid templating into a Liquified struct.

The Liquified Struct contains FrontMatter and an AST describing the page content / templating.

A Transpiler can then walk the AST converting the templating to another language (e.g. PHP).
