var btn3 = document.getElementById("btn3");
var btn2 = document.getElementById("btn2");
var btn1 = document.getElementById("btn1");
var btn4 = document.getElementById("btn4");
var menuIsOpen = false;
var width = screen.width;

function setMenuStyleAfterWindowChange() {
    setMenuStyle();
}

function setMenuStyle() {
    if (width < 555) {
        console.log("nobig");
        btn3.style.display = "none";
        btn2.style.display = "none";
        btn1.style.display = "none";
        btn4.style.display = "block";
    } else if (width > 555) {
        console.log("big");
        btn4.style.display = "none";

    }
}

function openMenu() {
    if (menuIsOpen) {
        btn3.style.display = "none";

        btn2.style.display = "none";
        btn1.style.display = "none";
        btn4.style.display = "block";
        btn4.innerHTML = "<i class=\"fa fa-flag\"></i>Menu";
        btn4.style.backgroundColor = "rgb(9,19,46)";
        menuIsOpen = false;
    } else {
        btn3.style.display = "block";
        btn3.style.backgroundColor = "rgb(7,14,35)";
        btn2.style.display = "block";
        btn2.style.backgroundColor = "rgb(7,14,35)";
        btn1.style.display = "block";
        btn1.style.backgroundColor = "rgb(7,14,35)";
        btn4.style.display = "block";
        btn4.style.backgroundColor = "rgb(7,14,35)";
        btn4.innerHTML = "Close";
        menuIsOpen = true;
    }
}

window.onload = setMenuStyle();
window.onchange = setMenuStyleAfterWindowChange();