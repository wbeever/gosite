function changePage(element) {
    document.getElementById("titleOut").innerHTML = titleChange(element.id);
    document.getElementById("linksOut").innerHTML = linkChange(element.id);
    document.getElementById("paragOut").innerHTML = paragChange(element.id);
}

var tArray = ["Home", "Gallery", "About", "Contact"];
var lArray = ["Home Links", "Gallery Links", "About Links", "Contact Links"];
var pArray = ["Home Paragraph", "Gallery Paragraph", "About Paragraph", "Contact Paragraph"];

function titleChange(d) {
    g = tArray[d];
    return g;
}

function linkChange(d) {
    h = lArray[d];
    return h;

}

function paragChange(d) {
    i = pArray[d];
    return i;
}