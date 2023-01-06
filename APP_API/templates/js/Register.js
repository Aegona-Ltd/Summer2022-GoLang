function load_ajax() {
    $.ajax({
        type: "POST",
        url: "http://localhost:5000/user",
        method: "POST",
        dataType: "json",
        data: JSON.stringify(
            {
                fullname : $('#fullname').val(),
                companyname :  $('#companyname').val(),
                businessphone :  $('#businessphone').val(),
                email :  $('#email').val(),
                message :  $('#message').val()
            }),
            
        })
};

$(document).on('click', '#btn2', function () {
var response = grecaptcha.getResponse();
// alert(response);
if (response == 0) {
    alert("Please check reCaptcha failed");
    return false;
}

});