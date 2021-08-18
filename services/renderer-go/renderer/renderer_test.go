package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render_heading(t *testing.T) {
	src := `# 見出し1
## 見出し2
### 見出し3
#### 見出し4
##### 見出し5
###### 見出し6`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<h1>見出し1</h1>
<h2>見出し2</h2>
<h3>見出し3</h3>
<h4>見出し4</h4>
<h5>見出し5</h5>
<h6>見出し6</h6>
`, html)
}

func Test_Render_list(t *testing.T) {
	src := ` - foo
 - bar
   - baz`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<ul>
<li>foo</li>
<li>bar
<ul>
<li>baz</li>
</ul>
</li>
</ul>
`, html)
}

func Test_Render_link(t *testing.T) {
	src := `foo https://google.com/ bar`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<p>foo <a href="https://google.com/">https://google.com/</a> bar</p>
`, html)
}
