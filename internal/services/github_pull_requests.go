package services

import "github.com/google/uuid"

const description = `
Here's a Markdown example with about 10k characters. It covers various Markdown elements like headers, paragraphs, lists, code, blockquotes, and more.

# Comprehensive Markdown Guide

Welcome to this comprehensive Markdown guide! Markdown is a lightweight markup language with plain-text formatting syntax. Its design allows it to be converted to many output formats, but it's most commonly used to generate HTML. Below is a detailed demonstration of various Markdown features.

## Table of Contents

1. [Headers](#headers)
2. [Emphasis](#emphasis)
3. [Lists](#lists)
4. [Links and Images](#links-and-images)
5. [Code](#code)
6. [Tables](#tables)
7. [Blockquotes](#blockquotes)
8. [Horizontal Rules](#horizontal-rules)
9. [Line Breaks](#line-breaks)
10. [Escaping Characters](#escaping-characters)

## Headers

Headers are created using the # symbol before your header text. The number of # symbols used will determine the size of the header.

# Header 1
## Header 2
### Header 3
#### Header 4
##### Header 5
###### Header 6

## Emphasis

You can emphasize text with bold or italic styles.

*Italic text with asterisks*

_Italic text with underscores_

**Bold text with asterisks**

__Bold text with underscores__

## Lists

Markdown supports ordered (numbered) and unordered (bulleted) lists.

### Ordered List

1. First item
2. Second item
3. Third item

### Unordered List

- Bulleted item
- Another bulleted item
- Yet another item

## Links and Images

Links and images are a vital part of any document.

### Links

[OpenAI](https://www.openai.com)

### Images

![Alt text for the image](https://example.com/image.jpg)

## Code

You can include inline code, like var example = true;, or code blocks.

function exampleFunction() {
  console.log("Hello, world!");
}

## Tables

Create tables by separating text with pipes  |  and dashes  - .

| Header 1 | Header 2 | Header 3 |
|----------|----------|----------|
| Row 1    | Data     | Data     |
| Row 2    | Data     | Data     |

## Blockquotes

Blockquotes are useful for quoting blocks of text.

> This is a blockquote.
> 
> This is the second paragraph in the blockquote.

## Horizontal Rules

A horizontal rule is a line that goes across the middle of the page.

---

Hyphens

***

Asterisks

___

Underscores

## Line Breaks

To create a line break, end a line with two or more spaces before hitting Return.

This line will  
break here.

## Escaping Characters

Use the backslash to escape Markdown characters.

\*This text is not italicized\*

---

*End of the Markdown guide*

This template should be close to 10k characters, incorporating a variety of Markdown features. Adjust the content if you need more precise control over the character count.		
`

const diff = `diff --git a/go.mod b/go.mod
index 4532416..780b81e 100644
--- a/go.mod
+++ b/go.mod
@@ -5,27 +5,37 @@ go 1.22
 require (
 	github.com/charmbracelet/bubbles v0.18.0
 	github.com/charmbracelet/bubbletea v0.25.0
+	github.com/charmbracelet/glamour v0.6.0
 	github.com/charmbracelet/lipgloss v0.10.0
 	github.com/google/uuid v1.6.0
+	github.com/muesli/termenv v0.15.2
 	github.com/spf13/cobra v1.8.0
 )
 
 require (
+	github.com/alecthomas/chroma v0.10.0 // indirect
 	github.com/atotto/clipboard v0.1.4 // indirect
 	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
+	github.com/aymerick/douceur v0.2.0 // indirect
 	github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81 // indirect
+	github.com/dlclark/regexp2 v1.4.0 // indirect
+	github.com/gorilla/css v1.0.0 // indirect
 	github.com/inconshreveable/mousetrap v1.1.0 // indirect
 	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
 	github.com/mattn/go-isatty v0.0.18 // indirect
 	github.com/mattn/go-localereader v0.0.1 // indirect
 	github.com/mattn/go-runewidth v0.0.15 // indirect
+	github.com/microcosm-cc/bluemonday v1.0.21 // indirect
 	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
 	github.com/muesli/cancelreader v0.2.2 // indirect
 	github.com/muesli/reflow v0.3.0 // indirect
-	github.com/muesli/termenv v0.15.2 // indirect
+	github.com/olekukonko/tablewriter v0.0.5 // indirect
 	github.com/rivo/uniseg v0.4.7 // indirect
 	github.com/sahilm/fuzzy v0.1.1-0.20230530133925-c48e322e2a8f // indirect
 	github.com/spf13/pflag v1.0.5 // indirect
+	github.com/yuin/goldmark v1.5.2 // indirect
+	github.com/yuin/goldmark-emoji v1.0.1 // indirect
+	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
 	golang.org/x/sync v0.1.0 // indirect
 	golang.org/x/sys v0.12.0 // indirect
 	golang.org/x/term v0.6.0 // indirect
diff --git a/go.sum b/go.sum
index 0a395ea..3dcb51d 100644
--- a/go.sum
+++ b/go.sum
@@ -1,39 +1,63 @@
+github.com/alecthomas/chroma v0.10.0 h1:7XDcGkCQopCNKjZHfYrNLraA+M7e0fMiJ/Mfikbfjek=
+github.com/alecthomas/chroma v0.10.0/go.mod h1:jtJATyUxlIORhUOFNA9NZDWGAQ8wpxQQqNSB4rjA/1s=
 github.com/atotto/clipboard v0.1.4 h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=
 github.com/atotto/clipboard v0.1.4/go.mod h1:ZY9tmq7sm5xIbd9bOK4onWV4S6X0u6GY7Vn0Yu86PYI=
+github.com/aymanbagabas/go-osc52 v1.0.3/go.mod h1:zT8H+Rk4VSabYN90pWyugflM3ZhpTZNC7cASDfUCdT4=
 github.com/aymanbagabas/go-osc52/v2 v2.0.1 h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=
 github.com/aymanbagabas/go-osc52/v2 v2.0.1/go.mod h1:uYgXzlJ7ZpABp8OJ+exZzJJhRNQ2ASbcXHWsFqH8hp8=
+github.com/aymerick/douceur v0.2.0 h1:Mv+mAeH1Q+n9Fr+oyamOlAkUNPWPlA8PPGR0QAaYuPk=
+github.com/aymerick/douceur v0.2.0/go.mod h1:wlT5vV2O3h55X9m7iVYN0TBM0NH/MmbLnd30/FjWUq4=
 github.com/charmbracelet/bubbles v0.18.0 h1:PYv1A036luoBGroX6VWjQIE9Syf2Wby2oOl/39KLfy0=
 github.com/charmbracelet/bubbles v0.18.0/go.mod h1:08qhZhtIwzgrtBjAcJnij1t1H0ZRjwHyGsy6AL11PSw=
 github.com/charmbracelet/bubbletea v0.25.0 h1:bAfwk7jRz7FKFl9RzlIULPkStffg5k6pNt5dywy4TcM=
 github.com/charmbracelet/bubbletea v0.25.0/go.mod h1:EN3QDR1T5ZdWmdfDzYcqOCAps45+QIJbLOBxmVNWNNg=
+github.com/charmbracelet/glamour v0.6.0 h1:wi8fse3Y7nfcabbbDuwolqTqMQPMnVPeZhDM273bISc=
+github.com/charmbracelet/glamour v0.6.0/go.mod h1:taqWV4swIMMbWALc0m7AfE9JkPSU8om2538k9ITBxOc=
 github.com/charmbracelet/lipgloss v0.10.0 h1:KWeXFSexGcfahHX+54URiZGkBFazf70JNMtwg/AFW3s=
 github.com/charmbracelet/lipgloss v0.10.0/go.mod h1:Wig9DSfvANsxqkRsqj6x87irdy123SR4dOXlKa91ciE=
 github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81 h1:q2hJAaP1k2wIvVRd/hEHD7lacgqrCPS+k8g1MndzfWY=
 github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81/go.mod h1:YynlIjWYF8myEu6sdkwKIvGQq+cOckRm6So2avqoYAk=
 github.com/cpuguy83/go-md2man/v2 v2.0.3/go.mod h1:tgQtvFlXSQOSOSIRvRPT7W67SCa46tRHOmNcaadrF8o=
+github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
+github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
+github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
+github.com/dlclark/regexp2 v1.4.0 h1:F1rxgk7p4uKjwIQxBs9oAXe5CqrXlCduYEJvrF4u93E=
+github.com/dlclark/regexp2 v1.4.0/go.mod h1:2pZnwuY/m+8K6iRw6wQdMtk+rH5tNGR1i55kozfMjCc=
 github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
 github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
+github.com/gorilla/css v1.0.0 h1:BQqNyPTi50JCFMTw/b67hByjMVXZRwGha6wxVGkeihY=
+github.com/gorilla/css v1.0.0/go.mod h1:Dn721qIggHpt4+EFCcTLTU/vk5ySda2ReITrtgBl60c=
 github.com/inconshreveable/mousetrap v1.1.0 h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=
 github.com/inconshreveable/mousetrap v1.1.0/go.mod h1:vpF70FUmC8bwa3OWnCshd2FqLfsEA9PFc4w1p2J65bw=
 github.com/kylelemons/godebug v1.1.0 h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=
 github.com/kylelemons/godebug v1.1.0/go.mod h1:9/0rRGxNHcop5bhtWyNeEfOS8JIWk580+fNqagV/RAw=
 github.com/lucasb-eyer/go-colorful v1.2.0 h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=
 github.com/lucasb-eyer/go-colorful v1.2.0/go.mod h1:R4dSotOR9KMtayYi1e77YzuveK+i7ruzyGqttikkLy0=
+github.com/mattn/go-isatty v0.0.16/go.mod h1:kYGgaQfpe5nmfYZH+SKPsOc2e4SrIfOl2e/yFXSvRLM=
 github.com/mattn/go-isatty v0.0.18 h1:DOKFKCQ7FNG2L1rbrmstDN4QVRdS89Nkh85u68Uwp98=
 github.com/mattn/go-isatty v0.0.18/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=
 github.com/mattn/go-localereader v0.0.1 h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=
 github.com/mattn/go-localereader v0.0.1/go.mod h1:8fBrzywKY7BI3czFoHkuzRoWE9C+EiG4R1k4Cjx5p88=
+github.com/mattn/go-runewidth v0.0.9/go.mod h1:H031xJmbD/WCDINGzjvQ9THkh0rPKHF+m2gUSrubnMI=
 github.com/mattn/go-runewidth v0.0.12/go.mod h1:RAqKPSqVFrSLVXbA8x7dzmKdmGzieGRCM46jaSJTDAk=
+github.com/mattn/go-runewidth v0.0.14/go.mod h1:Jdepj2loyihRzMpdS35Xk/zdY8IAYHsh153qUoGf23w=
 github.com/mattn/go-runewidth v0.0.15 h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=
 github.com/mattn/go-runewidth v0.0.15/go.mod h1:Jdepj2loyihRzMpdS35Xk/zdY8IAYHsh153qUoGf23w=
+github.com/microcosm-cc/bluemonday v1.0.21 h1:dNH3e4PSyE4vNX+KlRGHT5KrSvjeUkoNPwEORjffHJg=
+github.com/microcosm-cc/bluemonday v1.0.21/go.mod h1:ytNkv4RrDrLJ2pqlsSI46O6IVXmZOBBD4SaJyDwwTkM=
 github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b h1:1XF24mVaiu7u+CFywTdcDo2ie1pzzhwjt6RHqzpMU34=
 github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b/go.mod h1:fQuZ0gauxyBcmsdE3ZT4NasjaRdxmbCS0jRHsrWu3Ho=
 github.com/muesli/cancelreader v0.2.2 h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=
 github.com/muesli/cancelreader v0.2.2/go.mod h1:3XuTXfFS2VjM+HTLZY9Ak0l6eUKfijIfMUZ4EgX0QYo=
 github.com/muesli/reflow v0.3.0 h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=
 github.com/muesli/reflow v0.3.0/go.mod h1:pbwTDkVPibjO2kyvBQRBxTWEEGDGq0FlB1BIKtnHY/8=
+github.com/muesli/termenv v0.13.0/go.mod h1:sP1+uffeLaEYpyOTb8pLCUctGcGLnoFjSn4YJK5e2bc=
 github.com/muesli/termenv v0.15.2 h1:GohcuySI0QmI3wN8Ok9PtKGkgkFIk7y6Vpb5PvrY+Wo=
 github.com/muesli/termenv v0.15.2/go.mod h1:Epx+iuz8sNs7mNKhxzH4fWXGNpZwUaJKRS1noLXviQ8=
+github.com/olekukonko/tablewriter v0.0.5 h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=
+github.com/olekukonko/tablewriter v0.0.5/go.mod h1:hPp6KlRPjbx+hW8ykQs1w3UBbZlj6HuIJcUGPhkA7kY=
+github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
+github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
 github.com/rivo/uniseg v0.1.0/go.mod h1:J6wj4VEh+S6ZtnVlnTBMWIodfgj8LQOQFoIToxlJtxc=
 github.com/rivo/uniseg v0.2.0/go.mod h1:J6wj4VEh+S6ZtnVlnTBMWIodfgj8LQOQFoIToxlJtxc=
 github.com/rivo/uniseg v0.4.7 h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=
@@ -45,15 +69,33 @@ github.com/spf13/cobra v1.8.0 h1:7aJaZx1B85qltLMc546zn58BxxfZdR/W22ej9CFoEf0=
 github.com/spf13/cobra v1.8.0/go.mod h1:WXLWApfZ71AjXPya3WOlMsY9yMs7YeiHhFVlvLyhcho=
 github.com/spf13/pflag v1.0.5 h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=
 github.com/spf13/pflag v1.0.5/go.mod h1:McXfInJRrz4CZXVZOBLb0bTZqETkiAhM9Iw0y3An2Bg=
+github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
+github.com/stretchr/testify v1.7.0 h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=
+github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
+github.com/yuin/goldmark v1.2.1/go.mod h1:3hX8gzYuyVAZsxl0MRgGTJEmQBFcNTphYh9decYSb74=
+github.com/yuin/goldmark v1.5.2 h1:ALmeCk/px5FSm1MAcFBAsVKZjDuMVj8Tm7FFIlMJnqU=
+github.com/yuin/goldmark v1.5.2/go.mod h1:6yULJ656Px+3vBD8DxQVa3kxgyrAnzto9xy5taEt/CY=
+github.com/yuin/goldmark-emoji v1.0.1 h1:ctuWEyzGBwiucEqxzwe0SOYDXPAucOrE9NQC18Wa1os=
+github.com/yuin/goldmark-emoji v1.0.1/go.mod h1:2w1E6FEWLcDQkoTE+7HU6QF1F6SLlNGjRIBbIZQFqkQ=
+golang.org/x/net v0.0.0-20221002022538-bcab6841153b h1:6e93nYa3hNqAvLr0pD4PN1fFS+gKzp2zAXqrnTCstqU=
+golang.org/x/net v0.0.0-20221002022538-bcab6841153b/go.mod h1:YDH+HFinaLZZlnHAfSS6ZXJJ9M9t4Dl22yv3iI2vPwk=
 golang.org/x/sync v0.1.0 h1:wsuoTGHzEhffawBOhz5CYhcrV4IdKZbEyZjBMuTp12o=
 golang.org/x/sync v0.1.0/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
+golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
+golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
+golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
 golang.org/x/sys v0.1.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
 golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
 golang.org/x/sys v0.12.0 h1:CM0HF96J0hcLAwsHPJZjfdNzs0gftsLfgKt57wWHJ0o=
 golang.org/x/sys v0.12.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
+golang.org/x/term v0.0.0-20210927222741-03fcf44c2211/go.mod h1:jbD1KX2456YbFQfuXm/mYQcufACuNUgVhRMnK/tPxf8=
 golang.org/x/term v0.6.0 h1:clScbb1cHjoCkyRbWwBEUZ5H/tIFu5TAXIqaZD0Gcjw=
 golang.org/x/term v0.6.0/go.mod h1:m6U89DPEgQRMq3DNkDClhWw02AUbt2daBVO4cn4Hv9U=
+golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
 golang.org/x/text v0.3.8 h1:nAL+RVCQ9uMn3vJZbV+MRnydTJFPf8qqY42YiA6MrqY=
 golang.org/x/text v0.3.8/go.mod h1:E6s5w1FMmriuDzIBO73fBruAKo1PCIq6d2Q6DHfQ8WQ=
+golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
 gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
+gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
+gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
 gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
diff --git a/internal/pages/pull_requests_review.go b/internal/pages/pull_requests_review.go
index 1993507..d91ad65 100644
--- a/internal/pages/pull_requests_review.go
+++ b/internal/pages/pull_requests_review.go
@@ -7,6 +7,7 @@ import (
 	"github.com/charmbracelet/bubbles/help"
 	"github.com/charmbracelet/bubbles/key"
 	tea "github.com/charmbracelet/bubbletea"
+	"github.com/charmbracelet/glamour"
 	"github.com/charmbracelet/lipgloss"
 )
 
@@ -108,7 +109,22 @@ func (p *PullRequestReview) View() string {
 	if p.currentPr != nil {
 		pr := p.currentPr
 		title := pr.Title
-		description := pr.Description
+
+		style := glamour.DefaultStyles["dracula"]
+		style.Document.Margin = func() *uint {
+			var zero uint = 0
+			return &zero
+		}()
+		renderer, err := glamour.NewTermRenderer(glamour.WithStyles(*style), glamour.WithWordWrap(p.width/2-2))
+		if err != nil {
+			panic(err)
+		}
+
+		description, err := renderer.Render(pr.Description)
+		if err != nil {
+			panic(err)
+		}
+
 		comments := strings.Join(pr.Comments, "\n\n")
 		statusChecks := strings.Join(pr.StatusChecks, "\n\n")
 		diff := pr.Diff
diff --git a/internal/services/github_pull_requests.go b/internal/services/github_pull_requests.go
index cafec7b..2c4bc4b 100644
--- a/internal/services/github_pull_requests.go
+++ b/internal/services/github_pull_requests.go
@@ -15,7 +15,7 @@ func newBogusPr() GitHubPullRequest {
 
 	return GitHubPullRequest{
 		Title:       "some pr" + uuid,
-		Description: "some long text" + uuid,
+		Description: "# some long text\n- [ ] something\n - [x] something else\n\n<div>Bogus</div>\n\n" + uuid,
 		Comments: []string{
 			"some comment" + uuid,
 			"some comment" + uuid,
`

type GitHubPullRequest struct {
	Title        string
	Description  string
	Comments     []string
	StatusChecks []string
	Diff         string
}

func newBogusPr() GitHubPullRequest {
	uuid := uuid.NewString()

	return GitHubPullRequest{
		Title:       "some pr" + uuid,
		Description: description,
		Comments: []string{
			"some comment" + uuid,
			"some comment" + uuid,
		},
		StatusChecks: []string{
			"some status check" + uuid,
			"some status check" + uuid,
		},
		Diff: diff,
	}
}

func newBogusPrs(amount int) []GitHubPullRequest {
	prs := make([]GitHubPullRequest, 0, amount)

	for range amount {
		prs = append(prs, newBogusPr())
	}

	return prs
}

type GitHubPullRequestService struct {
	prs []GitHubPullRequest
}

func NewGitHubPullRequestService() *GitHubPullRequestService {
	return &GitHubPullRequestService{
		prs: newBogusPrs(50),
	}
}

func (g *GitHubPullRequestService) GetNext() (pr *GitHubPullRequest, ok bool) {
	if len(g.prs) > 0 {
		pr := g.prs[0]
		g.prs = g.prs[1:]
		return &pr, true
	}

	return nil, false
}
