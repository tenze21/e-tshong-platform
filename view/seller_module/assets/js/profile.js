window.onload = function () {
  greetSeller();
  fetch("/products/" + contactNumber)
    .then((res) => res.text())
    .then((data) => {
      if (data) {
        showProducts(data);
      }
    })
    .catch((error) => console.error("Error fetching products:", error));
};

function showProducts(data) {
  const noProducts = document.querySelector(".no_post");
  noProducts.style.display = "none";
  const products = JSON.parse(data);
  products.forEach((product) => newProduct(product));
}

function newProduct(product) {
    const productsWrapper = document.querySelector(".d_grid");
    // Create card container
    const container = document.createElement("div");
    container.classList.add("card_container");
    productsWrapper.appendChild(container);

    // Create product image container
    const productImage = document.createElement("div");
    productImage.classList.add("product_img");
    container.appendChild(productImage);

    // Create and append product image
    const img = document.createElement("img");
    img.src = `data:image/jpeg;base64,${product.productimg1}`;
    img.alt = "product image";
    productImage.appendChild(img);

    // Create card body
    const cardBody = document.createElement("div");
    cardBody.classList.add("card_body");
    container.appendChild(cardBody);

    // Create and append product title
    const productTitle = document.createElement("h3");
    productTitle.classList.add("product_title");
    productTitle.textContent = product.ptitle;
    cardBody.appendChild(productTitle);

    // Create price container
    const price = document.createElement("div");
    price.classList.add("price");
    cardBody.appendChild(price);

    // Create and append actual price
    const actualPrice = document.createElement("p");
    actualPrice.classList.add("actual_price");
    actualPrice.innerHTML = `Nu. <span id="actual_price">${product.aprice}</span>`;
    if(product.aprice<=0){
      actualPrice.style.display="none";
    }
    price.appendChild(actualPrice);

    // Create and append new price
    const newPrice = document.createElement("p");
    newPrice.classList.add("new_price");
    newPrice.innerHTML = `Nu. <span id="new_price">${product.sprice}</span>`;
    price.appendChild(newPrice);

    // Create and append seller contact
    const contact = document.createElement("p");
    contact.classList.add("contact");
    contact.innerHTML = `Seller Contact: <span id="seller_contact">${product.contact_number}</span>`;
    cardBody.appendChild(contact);

    // Create and append card buttons
    const buttonContainer= document.createElement("div");
    buttonContainer.classList.add("card_btns");
    cardBody.appendChild(buttonContainer);

    const viewMoreButton = document.createElement("button");
    viewMoreButton.classList.add("view_more");
    viewMoreButton.textContent = "View More";
    viewMoreButton.onclick = () => showProductDetail(product.productid);
    buttonContainer.appendChild(viewMoreButton)

    const deleteButton= document.createElement("button");
    deleteButton.classList.add("delete");
    deleteButton.textContent="Delete";
    buttonContainer.appendChild(deleteButton);
}

function showProductDetail(pid){
  // Add the product id to the URL parameter
  const url = `product.html?pid=${pid}`;
  window.open(url, "_self");
}

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

// form elements
let form = document.querySelector("#profile_editing_form");
let firstNameEl = document.getElementById("first_name");
let lastNameEl = document.getElementById("last_name");
let emailEl = document.getElementById("email");
let contactNumberEl = document.getElementById("contact_number");
let genders = document.getElementsByName("gender");

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
  console.log(genderValue);
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

// getting form data
function getFormData() {
  let newData = {
    fname: firstNameEl.value,
    lname: lastNameEl.value,
    cnumber: parseInt(contactNumberEl.value),
    email: emailEl.value,
    gender: getGender(),
  };
  return newData;
}

function getGender() {
  for (i = 0; i < genders.length; i++) {
    if (genders[i].checked) {
      return genders[i].value;
    }
  }
}

// form validation and detail updation
let isFirstNameValid = true,
  isLastNameValid = true,
  isphoneNumberValid = true;
form.addEventListener("submit", (e) => {
  e.preventDefault();

  let isFormValid = isFirstNameValid && isLastNameValid && isphoneNumberValid;
  if (isFormValid) {
    let newSellerDetail = getFormData();
    fetch("/seller/" + contactNumber, {
      method: "PUT",
      body: JSON.stringify(newSellerDetail),
      headers: { "Content-Type": "application/json; charset=UTF-8" },
    }).then((res) => {
      if (res.ok) {
        showNewSellerData(newSellerDetail);
        editModel.close();
        greetSeller();
      } else {
        alert("server: update request error.");
      }
    });
  } else {
    alert("invalid form");
  }
});

// Show form input error
function showError(input, message) {
  const formField = input.parentElement;
  const error = formField.querySelector("small");
  error.textContent = message;
}
// Hide form input error
function hideError(input) {
  const formField = input.parentElement;
  const error = formField.querySelector("small");
  error.textContent = "";
}
// Verify first name
firstNameEl.addEventListener("input", () => {
  const re = /^[a-zA-Z\s]+(?:\s+[a-zA-Z\s]+)?$/;
  if (!re.test(firstNameEl.value)) {
    showError(firstNameEl, "Name Cannot have numbers and special characters.");
    isFirstNameValid = false;
  } else {
    hideError(firstNameEl);
    isFirstNameValid = true;
  }
});
// verify last name
lastNameEl.addEventListener("input", () => {
  const re = /^[a-zA-Z\s]+(?:\s+[a-zA-Z\s]+)?$/;
  if (!re.test(lastNameEl.value)) {
    showError(lastNameEl, "Name Cannot have numbers and special characters.");
    isLastNameValid = false;
  } else {
    hideError(lastNameEl);
    isLastNameValid = true;
  }
});
// vaerify contact number
contactNumberEl.addEventListener("input", () => {
  const re = /^(17|77)\d{6}$/;
  if (!re.test(contactNumberEl.value)) {
    showError(contactNumberEl, "Invalid phone number.");
    isphoneNumberValid = false;
  } else {
    hideError(contactNumberEl);
    isphoneNumberValid = true;
  }
});

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
function greetSeller() {
  fetch("/seller/" + contactNumber)
    .then((response) => response.text())
    .then((seller) => greet(seller))
    .catch((error) => console.error("Error fetching seller data:", error));
}
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

// Display seller details.
function showSeller(seller) {
  const data = JSON.parse(seller);
  sName.textContent = `${data.fname} ${data.LastName}`;
  email.textContent = data.email;
  contactNum.textContent = data.cnumber;
  gender.textContent = data.gender;
}
function showNewSellerData(seller) {
  sName.textContent = `${seller.fname} ${seller.lname}`;
  email.textContent = seller.email;
  contactNum.textContent = seller.cnumber;
  gender.textContent = seller.gender;
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
document.addEventListener("DOMContentLoaded", function () {
  const profileInput = document.getElementById("profile");
  const profileImg = document.querySelector(".profile_img");

  profileInput.addEventListener("change", () => {
    const file = profileInput.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64Profile = reader.result;
        const profileObj = {
          profilepicture: base64Profile,
        };

        fetch("/profile/" + contactNumber, {
          method: "PUT",
          body: JSON.stringify(profileObj),
          headers: {
            "Content-Type": "application/json; charset=UTF-8",
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
            console.error("Error:", error);
          });
      };
      reader.readAsDataURL(file);
    }
  });
});
