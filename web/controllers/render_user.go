package controllers

import "github.com/gofiber/fiber/v2"

var sampleMD = `# test

**I would like to** _test_ this ~~where~~ I see how powerful is 

<details>
<summary>Click to expand/collapse</summary>

this markdown compiler

</details>

## how much stuff I can do

| Name    | Age |
| ------- | --- |
| Alice   | 25  |
| Bob     | 30  |

---
### and how far I can go with this

The quadratic formula is given by:

$$
x = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a}
$$

where a, b, and c are coefficients of a quadratic equation.


![something](/hal9000.png)

where I [link](https://google.com) stuff
`

// handler to render homepage of a user
func RenderUser(ctx *fiber.Ctx) error {
	var homeView string
	tab := ctx.Query("tab")
	if tab == "repositories" {
		homeView = "views/user/list_repos"
	} else {
		homeView = "views/user/user"
	}

	return ctx.Render(homeView, fiber.Map{
		"sample": sampleMD,
	}, "main")
}
