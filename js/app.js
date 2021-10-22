function maxValue(str, column) {
    const headers = str.slice(0, str.indexOf("\n")).split(",");
    const rows = str.slice(str.indexOf("\n") + 1).split("\n");

    const arr = rows.map(function (row) {
        const values = row.split(",");
        return headers.reduce(function (object, header, index) {
            object[header] = values[index];
            return object;
        }, {});
    });

    let filtered = [];

    for (let i = 0; i < arr.length; i++) {
        let obj = arr[i];

        for (let key in obj) {
            if (key === column) {
                filtered.push(obj[key]);
            }
        }

    }

    return parseInt(filtered.sort((a, b) => b - a)[0])
}

console.log(typeof maxValue("id,name,age,act.,room,dep.\n1,Jack,68,T,13,8\n17,Betty,28,F,15,7", "age"))
console.log(maxValue("area,land\n3722,CN\n6612,RU\n3855,CA\n3797,USA", "area"))
console.log(maxValue("city,temp2,temp\nParis,7,-3\nDubai,4,-4\nPorto,-1,-2", "temp"))