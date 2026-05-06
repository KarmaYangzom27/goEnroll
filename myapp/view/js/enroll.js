window.onload = function () {
    fetch('/students')
        .then(response => response.text())
        .then(data => getStudents(data));

    fetch('/courses')
        .then(response => response.text())
        .then(data => getCourses(data));

    fetch('/enrolls')
        .then(response => response.text())
        .then(data => getAllEnroll(data));
}

var selectedRow = null;

function getStudents(data) {
    const students = []
    const allStudents = JSON.parse(data)
    allStudents.forEach(stud => {
        students.push(stud.stdid)
    });

    var select = document.getElementById("sid")
    for (var i = 0; i < students.length; i++) {
        var sid = students[i];
        var option = document.createElement("option")
        option.textContent = sid;
        option.value = sid;
        select.appendChild(option);
    }
}

function getCourses(data) {
    const courses = []
    const allCourses = JSON.parse(data)
    allCourses.forEach(course => {
        courses.push(course.cid)
    });

    var option = "";
    for (var i = 0; i < courses.length; i++) {
        option += '<option value="' + courses[i] + '">' + courses[i] + '</option>'
    }
    document.getElementById("cid").innerHTML = option;
}

function addEnroll() {
    var _data = {
        stdid: parseInt(document.getElementById("sid").value),
        cid: document.getElementById("cid").value,
    }
    var sid = _data.stdid;
    var cid = _data.cid;

    if (isNaN(sid) || cid == "") {
        alert("Select valid data")
        return
    }

    fetch('/enroll', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok) {
            fetch('/enroll/' + sid + '/' + cid)
                .then(response => response.text())
                .then(data => getEnrolled(data))
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        if (e == "Error: Forbidden") {
            alert(e + ". Duplicate entry!")
        }
    });
    resetFields();
}

function resetFields() {
    document.getElementById("sid").value = "";
    document.getElementById("cid").value = "";
}

function getEnrolled(data) {
    const enrolled = JSON.parse(data)
    showTable(enrolled)
}

function showTable(enrolled) {
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.length);
    var td = []
    for (i = 0; i < table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i);
    }
    td[0].innerHTML = enrolled.stdid;
    td[1].innerHTML = enrolled.cid;
    td[2].innerHTML = enrolled.date.split("T")[0];
    td[3].innerHTML = '<input type="button" onclick="deleteEnroll(this)" value="Delete" id="button-1">';
}

function getAllEnroll(data) {
    const allenroll = JSON.parse(data)
    allenroll.forEach(enroll => {
        showTable(enroll)
    });
}

const deleteEnroll = async (r) => {
    if (confirm('Are you sure you want to DELETE this?')) {
        selectedRow = r.parentElement.parentElement;
        sid = selectedRow.cells[0].innerHTML;
        cid = selectedRow.cells[1].innerHTML;
        fetch('/enroll/' + sid + "/" + cid, {
            method: "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
        }).then(response => {
            if (response.ok) {
                var rowIndex = selectedRow.rowIndex;
                if (rowIndex > 0) {
                    document.getElementById("myTable").deleteRow(rowIndex);
                }
            }
        });
    }
}