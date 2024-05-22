// product image upload
const img1= document.querySelector(".img1");
const img2= document.querySelector(".img2");
const img3= document.querySelector(".img3");
const img4= document.querySelector(".img4");

const img1Input= document.querySelector("#img1");
const img2Input= document.querySelector("#img2");
const img3Input= document.querySelector("#img3");
const img4Input= document.querySelector("#img4");

img1Input.onchange=function (){
    img1.src= URL.createObjectURL(img1Input.files[0]);
};
img2Input.onchange=function (){
    img2.src= URL.createObjectURL(img2Input.files[0]);
};
img3Input.onchange=function (){
    img3.src= URL.createObjectURL(img3Input.files[0]);
};
img4Input.onchange=function (){
    img4.src= URL.createObjectURL(img4Input.files[0]);
};


// Greeting seller
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


document.getElementById('submit_btn').addEventListener('click', function (e) {
    e.preventDefault(); // Prevent the default form submission
  
    const form = document.querySelector('form');
    const formData = new FormData(form);
  
    fetch('/product/' + contactNumber, {
      method: 'POST',
      body: formData,
    })
    .then(response => {
      if (response.ok) {
        alert('Product uploaded successfully.');
        location.reload();
      } else {
        alert('Failed to upload product.');
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });
  });
  