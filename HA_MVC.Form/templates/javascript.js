function load_ajax() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form",
    method: "GET",
    dataType: "json",
    success: function (data) {
      // function createTable(tableData, TableRow) {

      //   var table = document.getElementById("myTable");

      //   var tableBody = document.createElement("tbody");
      //   tableData.forEach(function (rowData) {
      //     var row = document.createElement("tr");
      //     rowData.forEach(function (cellData) {
      //       var cell = document.createElement("th");
      //       cell.appendChild(document.createTextNode(cellData));
      //       row.appendChild(cell);
      //     });
      //     tableBody.appendChild(row);
      //   });
      //   TableRow.forEach(function (rowData) {
      //     let row = document.createElement("tr");
      //     for (let property in rowData) {
      //       let cell = document.createElement("td");
      //       cell.appendChild(document.createTextNode(rowData[property]));
      //       row.appendChild(cell);
      //     }
      //     tableBody.appendChild(row);
      //   });
      //   table.appendChild(tableBody);
      //   document.body.appendChild(table);
      // }

      // function getdata(item) {
      //         return [item.id, item.fullname, item.companyname, item.businessphone, item.email, item.message, item.createdtime, item.updatedtime];
      // };

      // var arr = data.map(getdata)
      // console.log(arr);
      // document.getElementById("row").innerHTML = arr;
      // createTable(
      //   [["Id", "Fullname", "Company name", "Business phone", "Email", "Message", "CreatedTime", "UpdatedTime"]],
      //   arr
      // );
      id: $("#id").html(data.map(getid));
      fullname: $("#fullname").html(data.map(getfullname));
      companyname: $("#companyname").html(data.map(getcompany));
      businessphone: $("#businessphone").html(data.map(getbusinessphone));
      email: $("#email").html(data.map(getemail));
      message: $("#message").html(data.map(getmessage));
      createdtime: $("#createtime").html(data.map(getcreatetime));
      updatedtime: $("#updatetime").html(data.map(getupdatetime));
      console.log(data.map(getupdatetime));

      function getid(item) {
        return item.id + "<br>";
      }
      function getfullname(item) {
        return item.fullname + "<br>";
      }
      function getcompany(item) {
        return item.companyname + "<br>";
      }
      function getbusinessphone(item) {
        return item.businessphone + "<br>";
      }
      function getemail(item) {
        return item.email + "<br>";
      }
      function getmessage(item) {
        return item.message + "<br>";
      }
      function getcreatetime(item) {
        return item.createtime +"<br>";
      }
      function getupdatetime(item) {
          return item.updatetime +"<br>";
      }
    },
  });
}
function GetEmail() {
  stringemail = document.getElementById("param").value;
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/" + stringemail,
    method: "GET",
    dataType: "json",
    success: function (data) {
      id: $("#id").html(data.map(getid));
      fullname: $("#fullname").html(data.map(getfullname));
      companyname: $("#companyname").html(data.map(getcompany));
      businessphone: $("#businessphone").html(data.map(getbusinessphone));
      email: $("#email").html(data.map(getemail));
      message: $("#message").html(data.map(getmessage));
      createdtime: $("#createtime").html(data.map(getcreatetime));
      updatedtime: $("#updatetime").html(data.map(getupdatetime));

      function getid(item) {
        return item.id + "<br>";
      }
      function getfullname(item) {
        return item.fullname + "<br>";
      }
      function getcompany(item) {
        return item.companyname + "<br>";
      }
      function getbusinessphone(item) {
        return item.businessphone + "<br>";
      }
      function getemail(item) {
        return item.email + "<br>";
      }
      function getmessage(item) {
        return item.message + "<br>";
      }
      function getcreatetime(item) {
        return item.createtime + "<br>";
      }
      function getupdatetime(item) {
        return item.updatetime + "<br>";
      }
    },
  });
}
function GetName() {
  stringname = document.getElementById("param").value;
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/name/" + stringname,
    method: "GET",
    dataType: "json",
    success: function (data) {
      id: $("#id").html(data.map(getid));
      fullname: $("#fullname").html(data.map(getfullname));
      companyname: $("#companyname").html(data.map(getcompany));
      businessphone: $("#businessphone").html(data.map(getbusinessphone));
      email: $("#email").html(data.map(getemail));
      message: $("#message").html(data.map(getmessage));
      createdtime: $("#createtime").html(data.map(getcreatetime));
      updatedtime: $("#updatetime").html(data.map(getupdatetime));
      function getid(item) {
        return item.id + "<br>";
      }
      function getfullname(item) {
        return item.fullname + "<br>";
      }
      function getcompany(item) {
        return item.companyname + "<br>";
      }
      function getbusinessphone(item) {
        return item.businessphone + "<br>";
      }
      function getemail(item) {
        return item.email + "<br>";
      }
      function getmessage(item) {
        return item.message + "<br>";
      }
      function getcreatetime(item) {
        return item.createtime + "<br>";
      }
      function getupdatetime(item) {
        return item.updatetime + "<br>";
      }
    },
  });
}
function GetOder() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/email",
    method: "GET",
    dataType: "json",
    success: function (data) {
      id: $("#id").html(data.map(getid));
      fullname: $("#fullname").html(data.map(getfullname));
      companyname: $("#companyname").html(data.map(getcompany));
      businessphone: $("#businessphone").html(data.map(getbusinessphone));
      email: $("#email").html(data.map(getemail));
      message: $("#message").html(data.map(getmessage));
      createdtime: $("#createtime").html(data.map(getcreatetime));
      updatedtime: $("#updatetime").html(data.map(getupdatetime));
      console.log(data.map(getid));
      function getid(item) {
        return item.id + "<br>";
      }
      function getfullname(item) {
        return item.fullname + "<br>";
      }
      function getcompany(item) {
        return item.companyname + "<br>";
      }
      function getbusinessphone(item) {
        return item.businessphone + "<br>";
      }
      function getemail(item) {
        return item.email + "<br>";
      }
      function getmessage(item) {
        return item.message + "<br>";
      }
      function getcreatetime(item) {
        return item.createtime + "<br>";
      }
      function getupdatetime(item) {
        return item.updatetime + "<br>";
      }
    },
  });
}
function GetOderByCreateTime() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/date",
    method: "GET",
    dataType: "json",
    success: function (data) {
      id: $("#id").html(data.map(getid));
      fullname: $("#fullname").html(data.map(getfullname));
      companyname: $("#companyname").html(data.map(getcompany));
      businessphone: $("#businessphone").html(data.map(getbusinessphone));
      email: $("#email").html(data.map(getemail));
      message: $("#message").html(data.map(getmessage));
      createdtime: $("#createtime").html(data.map(getcreatetime));
      updatedtime: $("#updatetime").html(data.map(getupdatetime));
      console.log(data.map(getid));
      function getid(item) {
        return item.id + "<br>";
      }
      function getfullname(item) {
        return item.fullname + "<br>";
      }
      function getcompany(item) {
        return item.companyname + "<br>";
      }
      function getbusinessphone(item) {
        return item.businessphone + "<br>";
      }
      function getemail(item) {
        return item.email + "<br>";
      }
      function getmessage(item) {
        return item.message + "<br>";
      }
      function getcreatetime(item) {
        return item.createtime + "<br>";
      }
      function getupdatetime(item) {
        return item.updatetime + "<br>";
      }
    },
  });
}