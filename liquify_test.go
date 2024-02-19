package liquify

import (
	"bytes"
	"fmt"
	"github.com/richardjennings/liquify/expr"
	"github.com/richardjennings/liquify/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TC struct {
	content string
	assert  func(t *testing.T, l *Liquified)
	config  parser.Config
}

func testDefaultConfig() parser.Config {
	return parser.Config{
		Delims:  []string{"{{", "}}", "{%", "%}"},
		Grammar: parser.Grammer{},
	}
}

func TestLiquify(t *testing.T) {
	tcs := []TC{
		{
			content: `{% assign description = "test assignment statement" %}`,
			assert: func(t *testing.T, l *Liquified) {
				assert.Equal(t, "assign description = \"test assignment statement\"", l.Ast.String())
				assert.Equal(t, 0, len(l.FrontMatter))
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				assert.Equal(t, "<?php $description = \"test assignment statement\";?>", string(v))
			},
			config: testDefaultConfig(),
		},
		{
			content: `---
some: value
another: value
this:
  - is:
    a: nested value
---
{% assign description = "front matter" %}`,
			assert: func(t *testing.T, l *Liquified) {
				expected := `assign description = "front matter"`
				assert.Equal(t, expected, l.Ast.String())
				assert.Equal(t, 3, len(l.FrontMatter))
			},
			config: testDefaultConfig(),
		},
		{
			content: `---
some: value
another: value
this:
  - is:
    a: nested value
---
{% assign description = "if statement consuming front matter" %}

{% if page.some == "value" %}
hello
{% else %}
goodbye
{% endif %}`,
			assert: func(t *testing.T, l *Liquified) {
				assert.Equal(t, 3, len(l.FrontMatter))
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				expected := `<?php $description = "if statement consuming front matter";?>

<?php if ($page["some"] == "value") { ?>
hello
<?php } else { ?>
goodbye
<?php } ?>`
				assert.Equal(t, expected, string(v))
			},
			config: testDefaultConfig(),
		},
		{
			// filter support
			content: `{% assign things = site.posts | where:"some","thing" %}`,
			assert: func(t *testing.T, l *Liquified) {
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				expected := `<?php $things = $site["posts"] /* filter where "some""thing" */;?>`
				assert.Equal(t, expected, string(v))
			},
			config: testDefaultConfig(),
		},
		{
			// plugin
			content: `{% j page.title %}`,
			assert: func(t *testing.T, l *Liquified) {
				p := PHP{
					TagParsers: map[string]func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error{
						"j": func(b *bytes.Buffer, t *parser.ASTTag, p PHP) error {
							v, err := expr.Parse(t.Args)
							if err != nil {
								panic(err)
							}
							b.Write([]byte(fmt.Sprintf(`<?php echo %s(%s);?>`, t.Name, p.Stmt(v))))
							return nil
						},
					},
				}
				v, err := p.Transpile(l)
				assert.Nil(t, err)
				expected := `<?php echo j($page["title"]);?>`
				assert.Equal(t, expected, string(v))
			},
			config: testDefaultConfig(),
		},
		{
			// capture
			config:  testDefaultConfig(),
			content: `{% capture background %}some value{% endcapture %}`,
			assert: func(t *testing.T, l *Liquified) {
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				assert.Equal(t, `<?php $background = "some value";?>`, string(v))
			},
		},
		{
			// comment
			config:  testDefaultConfig(),
			content: `{% comment %}some value{% endcomment %}`,
			assert: func(t *testing.T, l *Liquified) {
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				assert.Equal(t, `/* some value */`, string(v))
			},
		},
		{
			// for
			config:  testDefaultConfig(),
			content: `{% for a in page.thing %}some value{% endfor %}`,
			assert: func(t *testing.T, l *Liquified) {
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				assert.Equal(t, `<?php for ($i, $a in $page["thing"]){ ?>some value<?php } ?>`, string(v))
			},
		},
	}
	for i, tc := range tcs {
		l, err := Liquify([]byte(tc.content), tc.config)
		if err != nil {
			t.Errorf("tc %d, error: %s", i, err)
		}
		tc.assert(t, l)
	}
}
