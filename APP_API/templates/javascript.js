function load_ajax() {
  $.ajax({
    type: "GET",
    url: "http://localhost:1974/user",
    method: "GET",
    dataType: "json",
    success: function (data) {
      let placeholder = document.querySelector("#tbody");
      let out = "";
      for(let product of data){
         out += `
            <tr>
               <td> ${product.userid}</td>
               <td> ${product.fullname}</td>
               <td> ${product.companyname}</td>
               <td> ${product.businessphone}</td>
               <td> ${product.email}</td>
               <td> ${product.message}</td>
               <td> ${product.createtime}</td>
               <td> ${product.updatetime}</td> 
            </tr>
         `;
      }
      placeholder.innerHTML = out;
    }
  });
}
function GetEmail() {
  stringemail = document.getElementById("param").value;
  $.ajax({
    type: "GET",
    url: "http://localhost:5000/user/" + stringemail,
    method: "GET",
    dataType: "json",
    success: function (data) {
      let placeholder = document.querySelector("#tbody");
      let out = "";
      for(let product of data){
         out += `
            <tr>
               <td> ${product.id}</td>
               <td> ${product.fullname}</td>
               <td> ${product.companyname}</td>
               <td> ${product.businessphone}</td>
               <td> ${product.email}</td>
               <td> ${product.message}</td>
               <td> ${product.createtime}</td>
               <td> ${product.updatetime}</td>
            </tr>
         `;
      }
      placeholder.innerHTML = out;
    }
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
      let placeholder = document.querySelector("#tbody");
      let out = "";
      for(let product of data){
         out += `
            <tr>
               <td> ${product.id}</td>
               <td> ${product.fullname}</td>
               <td> ${product.companyname}</td>
               <td> ${product.businessphone}</td>
               <td> ${product.email}</td>
               <td> ${product.message}</td>
               <td> ${product.createtime}</td>
               <td> ${product.updatetime}</td>
            </tr>
         `;
      }
      placeholder.innerHTML = out;
    }
  });
}
function GetOder() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/email",
    method: "GET",
    dataType: "json",
    success: function (data) {
      let placeholder = document.querySelector("#tbody");
      let out = "";
      for(let product of data){
         out += `
            <tr>
               <td> ${product.id}</td>
               <td> ${product.fullname}</td>
               <td> ${product.companyname}</td>
               <td> ${product.businessphone}</td>
               <td> ${product.email}</td>
               <td> ${product.message}</td>
               <td> ${product.createtime}</td>
               <td> ${product.updatetime}</td>
            </tr>
         `;
      }
      placeholder.innerHTML = out;
    }
  });
}
function GetOderByCreateTime() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/form/date",
    method: "GET",
    dataType: "json",
    success: function (data) {
      let placeholder = document.querySelector("#tbody");
      let out = "";
      for(let product of data){
         out += `
            <tr>
               <td> ${product.id}</td>
               <td> ${product.fullname}</td>
               <td> ${product.companyname}</td>
               <td> ${product.businessphone}</td>
               <td> ${product.email}</td>
               <td> ${product.message}</td>
               <td> ${product.createtime}</td>
               <td> ${product.updatetime}</td>
            </tr>
         `;
      }
      placeholder.innerHTML = out;
    }
  });
    $(document).ready( function () {
      $('#myTable').bdt();
    });

}



