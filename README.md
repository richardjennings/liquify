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

### PHP Transpiler

#### For

A `{% for %}` block will compile to a PHP loop. @todo add support for all For constructs in liquid.

Example:
```
{% for a in page.thing %}some value{% endfor %}
```

will transpile to:

```
<?php for ($i, $a in $page["thing"]){ ?>some value<?php } ?>
```

#### Include

`{% include %}` tags are converted to a `render()` function call with the template and an array of
arguments.

Example:

```
{% include components/something.html title=page.title
background='some string value'
something=page.something
another="string" %}
```

will transpile to:

```
<?php $values = ["title"=>$page["title"],"background"=>"some string value","something"=>$page["something"],"another"=>"string",]; ?><?php render("components/something.html", $values);?>
```

#### Assign

```
{% assign things = site.posts | where:"some","thing" %}
```

```
<?php $things = $site["posts"] /* filter where "some""thing" */;?>
```

#### Plugin

```
{% j page.title %}
```

```
<?php echo j($page["title"]);?>
```

#### Capture
```
{% capture background %}some value{% endcapture %}
```

```
<?php $background = "some value";?>
```

#### Comment
```
{% comment %}some value{% endcomment %}
```

```
/* some value */
```

#### Filter

Filters are printed commented out.