function generateTableHead(data, table) {
    let thead = table.createTHead();
    let row = thead.insertRow();
    generateHeaderCell("Player", row);
    generateHeaderCell("Play time", row);
    generateHeaderCell("Mine count", row);
    generateHeaderCell("Field size", row);
}

function generateTable(data, table) {
    for (let dataRow of data) {
        let row = table.insertRow();
        //console.log(dataRow['fieldSize']);
        generateCellFromKey(dataRow, row, 'player');
        generateCellFromKey(dataRow, row, 'playTime');
        generateCellFromKey(dataRow, row, 'mineCount');
        generateCellFromKey(dataRow, row, 'fieldSize');
    }
}

function generateHeaderCell(title, row) {
    let th = document.createElement("th");
    let text = document.createTextNode(title);
    th.appendChild(text);
    row.appendChild(th);
}

function generateCellFromKey(dataRow, row, key) {
    let cell = row.insertCell();
    let text = document.createTextNode(dataRow[key]);
    //console.log(dataRow[key]);
    cell.appendChild(text);
}

function getDataFromApi() {
    const url = "/allresults";
    var headers = {}
    fetch(url, {method: 'GET', headers: headers})
        .then(response => response.json())
        .then(data => {
            generateLeaderBoard(data);
        });
}

function generateLeaderBoard(data) {
    let table = document.querySelector("table");
    //console.log("all received data: ",data);
    // console.log("first object from received data: ",Object.keys(data[0]));
    // console.log("field size of first object from received data: ",data[0][2]);
    generateTableHead(Object.keys(data[0]), table);
    generateTable(data, table);
}

window.onload = getDataFromApi();
