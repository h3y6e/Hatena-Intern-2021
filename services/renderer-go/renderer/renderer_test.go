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

func Test_Render_mathjax(t *testing.T) {
	src := `$$
\mathbb{E}(X) = \int x d F(x) = \left\{ \begin{aligned} \sum_x x f(x) \; & \text{ if } X \text{ is discrete} 
\\ \int x f(x) dx \; & \text{ if } X \text{ is continuous }
\end{aligned} \right.
$$
	
Inline math $\frac{1}{2}$`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<p><span class="math display">\[\mathbb{E}(X) = \int x d F(x) = \left\{ \begin{aligned} \sum_x x f(x) \; & \text{ if } X \text{ is discrete} 
\\ \int x f(x) dx \; & \text{ if } X \text{ is continuous }
\end{aligned} \right.
\]</span></p>
<p>Inline math <span class="math inline">\(\frac{1}{2}\)</span></p>
`, html)
}
