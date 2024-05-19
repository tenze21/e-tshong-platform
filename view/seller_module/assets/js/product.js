function getCookie(name) {
    let cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i].trim();
        // Check if this cookie string begins with the name we want
        if (cookie.indexOf(name + '=') == 0) {
            return cookie.substring((name + '=').length, cookie.length);
        }
    }
    return null;
}

// Use the function to get the contactnumber cookie
let contactNumber = getCookie('contactnumber');

fetch("/seller/" + contactNumber)
.then((response)=>response.text())
.then((seller)=>greet(seller))
.catch((error) => console.error('Error fetching seller data:', error));

const sellerName= document.getElementById("user_name");
function greet(seller){
    const data= JSON.parse(seller)
    sellerName.textContent=data.fname
}
