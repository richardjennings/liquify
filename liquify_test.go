package liquify

import (
	"github.com/richardjennings/liquify/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TC struct {
	path   string
	assert func(t *testing.T, l *Liquified)
	config parser.Config
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
			path: "testdata/_posts/1.md",
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
			path: "testdata/_posts/2.md",
			assert: func(t *testing.T, l *Liquified) {
				assert.Equal(t, "assign description = \"front matter\"", l.Ast.String())
				assert.Equal(t, 3, len(l.FrontMatter))
			},
			config: testDefaultConfig(),
		},
		{
			path: "testdata/_posts/3.md",
			assert: func(t *testing.T, l *Liquified) {
				expected := "assign description = \"if statement consuming front matter\"\n\nif page.some == \"value\"\nhello\nelse \ngoodbye\nendif "
				assert.Equal(t, expected, l.Ast.String())
				assert.Equal(t, 3, len(l.FrontMatter))
				v, err := PHP{}.Transpile(l)
				assert.Nil(t, err)
				expected = `<?php $description = "if statement consuming front matter";?>

<?php if ($page["some"] == "value") { ?>
hello
<?php } else { ?>
goodbye
<?php } ?>`
				assert.Equal(t, expected, string(v))
			},
			config: testDefaultConfig(),
		},
	}
	for i, tc := range tcs {
		l, err := Liquify(tc.path, tc.config)
		if err != nil {
			t.Errorf("tc %d, error: %s", i, err)
		}
		tc.assert(t, l)
	}
}
