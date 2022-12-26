function Validate() {
    var fullname = document.forms.myform.full_name.value;
    var companyname = document.forms.myform.company_name.value;
    var businessphone = document.forms.myform.business_phone.value;
    var email = document.forms.myform.email.value;
    var message = document.forms.myform.message.value;
    
    var regFullname = /^([a-zA-Z0-9]+|[a-zA-Z0-9]+\s{1}[a-zA-Z0-9]{1,}|[a-zA-Z0-9]+\s{1}[a-zA-Z0-9]{3,}\s{1}[a-zA-Z0-9]{1,})+$/;
    var regCompanyname = /^([a-zA-Z0-9]+|[a-zA-Z0-9]+\s{1}[a-zA-Z0-9]{1,}|[a-zA-Z0-9]+\s{1}[a-zA-Z0-9]{3,}\s{1}[a-zA-Z0-9]{1,})+$/;
    var regBusinessphone=/^\d{10}$/;   
    var regEmail=/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/; 

    //Javascript reGex for Email Validation.
    if (fullname == "" || !regFullname.test(fullname)){
        alert("Please enter a valid full_name.");
        fullname.focus();
        return false;
    }		
    if (companyname == "" || !regCompanyname.test(companyname)){
        alert("Please enter a valid company_name.");
        companyname.focus();
        return false;
    }	
    if (businessphone == "" || !regBusinessphone.test(businessphone)){
        alert("Please enter a valid business_phone.");
        businessphone.focus();
        return false;
    }								 							 						 // Javascript reGex for Phone Number validation.						 // Javascript reGex for Name validation
    if (email == "" || !regEmail.test(email)) {
        alert("Please enter a valid e-mail address.");
        email.focus();
        return false;
    }
    
    if (message == "") {
        alert("Please enter your message");
        message.focus();
        return false;
    }
    return true;
}
function onclick(){
    var x = Validate();
    if (x == 0){
        document.getElementById("myform").reset();
    }
    
}