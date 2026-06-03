# fhtml
Flux HTML templating language -- an HTML-aided templating language, inspired by Vue.js and Go's text/template.

## How it looks like

```html
<article>
  <section>
    <h3>{{ name }}</h3>

    <ul>
      <li
        @for="item on items" <!-- Output a "li" tag for each value from "items" array -->
        :class:append="checked @ .checked" <!-- Append to the "class" attribute when ".checked" is true -->
        :class:append="color-red @ .first" <!-- Append to the "class" attribute when ".first" is true -->
        :aria-selected=".checked" <!-- Include the "aria-selected" attribute if ".checked" is true -->
      >{{ item.name }}</li>
    <ul>

    <ul>
      <li
        @for="items" <!-- Output a "li" tag for each value from "items" array -->
        :class:append="checked @ .checked" <!-- Append to the "class" attribute when ".checked" is true -->
        :class:append="color-red @ .first" <!-- Append to the "class" attribute when ".first" is true -->
      >{{ .name }}</li>
    <ul>

    <p @if="hasDetails">{{ globalDetails }}</p> <!-- Include "p" tag if "hasDetails" is true -->
  </section>
</article>
```
