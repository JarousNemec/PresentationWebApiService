function remove() {
    var div;
    var id = "etarget_ifranme_";
    var ad_id = "";
    for (var i = 0; i < 1000; i++) {
        ad_id = id.concat(i.toString());
        div = document.getElementById(ad_id);
        try {
            div.parentNode.removeChild(div);
        } catch (e) {
        }
    }

    var div2 = document.getElementById("etBotBarWrap");
    try {
        div2.parentNode.removeChild(div2);
    } catch (e) {
    }
    getUNamedDivs();
}

var divs = [];

function getUNamedDivs() {
    $(document).ready(function () {
        $('div').each(function () {

            var i = $(this);
            if (i.attr("id")) {
                console.log(i);
            } else {
                divs.push(i);
            }
        });
        mzF();
    });
}

function mzF() {
    var idF = parseInt(divs.length);
    var idFm1 = idF - 1;
    var idFm2 = idF - 2;
    var idFm3 = idF - 3;
    console.log(divs[idFm1]);
    divs[idFm1].remove();
    divs[idFm2].remove();
    divs[idFm3].remove();
removingIframes();
}

function removingIframes() {
    var i = 0;
    setInterval(function () {
        try {
            $(document).ready(function () {
                $('iframe').each(function () {

                    var i = $(this);
                    if (i.attr("id")) {

                    } else {
                        i.remove();
                    }

                });
            });
        }catch (e) {
            console.log(e);
        }
        i++;
    }, 5000);
}

window.onload = remove();




