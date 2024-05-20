const editProfile = document.querySelector(".edit_btn");
const closeEditModel = document.querySelector("#model_close");
const editModel = document.querySelector(".profile_edit_model");

editProfile.addEventListener("click", () => {
  editModel.showModal();
  getInitialUserInfo();
});

closeEditModel.addEventListener("click", () => {
  editModel.close();
});

// getting user details
function getInitialUserInfo() {
  let name = document.querySelector("#name").textContent.split(" ");
  let firstName = name[0];
  let lastName = name[1];
  let emailValue = document.querySelector("#email_value").textContent;
  let contactNumberValue = document.querySelector(
    "#contact_number_value"
  ).textContent;
  let genderValue = document.querySelector("#gender").textContent;

  let firstNameEl = document.getElementById("first_name");
  let lastNameEl = document.getElementById("last_name");
  let emailEl = document.getElementById("email");
  let contactNumberEl = document.getElementById("contact_number");
  let genders = document.getElementsByName("gender");

  firstNameEl.value = firstName;
  lastNameEl.value = lastName;
  emailEl.value = emailValue;
  contactNumberEl.value = contactNumberValue;
  genders.forEach((gender) => {
    if (gender.value === genderValue) {
      gender.checked = true;
    }
  });
}

//  post delete confirmation model
const deleteBtn = document.querySelectorAll(".delete");
const cancelDelete = document.querySelector(".cancel");
const confirmationModal = document.querySelector(".confirmation_modal");

deleteBtn.forEach((btn) => {
  btn.addEventListener("click", () => {
    confirmationModal.showModal();
  });
});
cancelDelete.addEventListener("click", () => {
  confirmationModal.close();
});

// upload profile
// const profile= document.querySelector(".profile_img");
// const profileInput= document.querySelector("#profile");
// profileInput.onchange= function(){
//     profile.src= URL.createObjectURL(profileInput.files[0]);
// };

// Greeting Seller
function getCookie(name) {
  let cookies = document.cookie.split(";");
  for (let i = 0; i < cookies.length; i++) {
    let cookie = cookies[i].trim();
    // Check if this cookie string begins with the name we want
    if (cookie.indexOf(name + "=") == 0) {
      return cookie.substring((name + "=").length, cookie.length);
    }
  }
  return null;
}

// Use the function to get the contactnumber cookie
let contactNumber = getCookie("contactnumber");

fetch("/seller/" + contactNumber)
  .then((response) => response.text())
  .then((seller) => greet(seller))
  .catch((error) => console.error("Error fetching seller data:", error));

const sellerName = document.getElementById("user_name");
function greet(seller) {
  const data = JSON.parse(seller);
  sellerName.textContent = data.fname;
}

// Seller profile
fetch("/seller/" + contactNumber)
  .then((response) => response.text())
  .then((seller) => showSeller(seller))
  .catch((error) => console.error("Error fetching seller data:", error));

const sName = document.getElementById("name");
const email = document.getElementById("email_value");
const contactNum = document.getElementById("contact_number_value");
const gender = document.getElementById("gender");
function showSeller(seller) {
  const data = JSON.parse(seller);
  console.log(data);
  sName.textContent = `${data.fname} ${data.LastName}`;
  email.textContent = data.email;
  contactNum.textContent = data.cnumber;
  gender.textContent = data.gender;
}

// Get profile
const profile = document.querySelector(".profile_img");
fetch("/seller/" + contactNumber)
  .then((response) => response.text())
  .then((seller) => showSellerProfile(seller))
  .catch((error) => console.error("Error fetching seller data:", error));

function showSellerProfile(seller) {
  const data = JSON.parse(seller);
  profile.src = `data:image/jpeg;base64,${data.profilepicture}`;
}

// update profile
document.addEventListener('DOMContentLoaded', function () {
    const profileInput = document.getElementById("profile");
    const profileImg = document.querySelector(".profile_img");
    
    profileInput.addEventListener('change', () => {
        const file = profileInput.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onloadend = () => {
                const base64Profile = reader.result;
                const profileObj = {
                    profilepicture: base64Profile,
                };
                
                fetch('/profile/' + contactNumber, {
                    method: "PUT",
                    body: JSON.stringify(profileObj),
                    headers: {
                        "Content-Type": "application/json; charset=UTF-8"
                    },
                })
                .then((res) => {
                    if (res.ok) {
                        profileImg.src = URL.createObjectURL(file);
                    } else {
                        alert("Server: update request error");
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
            };
            reader.readAsDataURL(file);
        }
    });
});

