// compile MD to HTML-viewable content
const readme = document.getElementById('readme');
if (readme) {
    const converter = new showdown.Converter();
    readme.innerHTML = converter.makeHtml(readme.innerText);
}
