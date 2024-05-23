const form = document.querySelector("#registration_form");
const firstNameEl = document.querySelector("#first_name");
const lastNameEl = document.querySelector("#last_name");
const phoneNumberEl = document.querySelector("#contact_number");
const passwordEl = document.querySelector("#password");
const confirmPasswordEl = document.querySelector("#confirm_password");
const emailEl = document.querySelector("#email");
const genderEl = document.getElementsByName("gender");
let selectedGender;
genderEl.forEach((radio) => {
  radio.addEventListener("click", () => {
    selectedGender = radio.value;
  });
});
let isFirstNameValid = false,
  isLastNameValid = false,
  isphoneNumberValid = false,
  isPasswordSecure = false,
  isConfirmPasswordValid = false;
form.addEventListener("submit", (e) => {
  e.preventDefault();

  let isFormValid =
    isFirstNameValid &&
    isLastNameValid &&
    isphoneNumberValid &&
    isPasswordSecure &&
    isConfirmPasswordValid;
    if (isFormValid) {
      let data = {
          fname: firstNameEl.value,
          lname: lastNameEl.value,
          cnumber: parseInt(phoneNumberEl.value),
          email: emailEl.value,
          gender: selectedGender,
          password: passwordEl.value,
      };
      console.log(JSON.stringify(data));

      // Register the user first
      fetch("/register", {
          method: "POST",
          headers: {
              "Content-Type": "application/json; charset=UTF-8",
          },
          body: JSON.stringify(data),
      })
      .then(response => {
          if (response.status === 201) {
              alert("Registration Successful. Welcome to e-tshong.");
              
              fetch("/sellerdetail/" + data.cnumber)
              .then(response => response.text())
              .then(data => {
                  addDefaultProfile(data);
              })
              .catch(error => {
                  console.error("Error:", error);
              });
          } 
      })
  }
});

function addDefaultProfile(seller){
  const data=JSON.parse(seller);
  fetch('/assets/images/default_profile.jpg')
  .then(response=>response.blob())
  .then(blob=>{
    const formData=new FormData();
    formData.append('cnumber', data.cnumber)
    formData.append('profile_picture', blob, 'default_profile.jpg');

    fetch("/profile", {
      method: 'POST',
      body: formData,
    })
    .then(response=>{
      if(response.status==201){
        console.log("default profile set successfully.");
        window.open("login.html", "_self");
      }else{
        console.error("failed to set default profile.")
      }
    })
    .catch(error=>{
      console.error('error:', error)
    })
  })
}

function showError(input, message) {
  const formField = input.parentElement;
  const error = formField.querySelector("small");
  error.textContent = message;
}

function hideError(input) {
  const formField = input.parentElement;
  const error = formField.querySelector("small");
  error.textContent = "";
}

firstNameEl.addEventListener("input", () => {
  const re = /^[a-zA-Z\s]+(?:\s+[a-zA-Z\s]+)?$/;
  if (!re.test(firstNameEl.value)) {
    showError(firstNameEl, "Name Cannot have numbers and special characters.");
  } else {
    hideError(firstNameEl);
    isFirstNameValid = true;
  }
});
lastNameEl.addEventListener("input", () => {
  const re = /^[a-zA-Z\s]+(?:\s+[a-zA-Z\s]+)?$/;
  if (!re.test(lastNameEl.value)) {
    showError(lastNameEl, "Name Cannot have numbers and special characters.");
  } else {
    hideError(lastNameEl);
    isLastNameValid = true;
  }
});

phoneNumberEl.addEventListener("input", () => {
  const re = /^(17|77)\d{6}$/;
  if (!re.test(phoneNumberEl.value)) {
    showError(phoneNumberEl, "Invalid phone number.");
  } else {
    hideError(phoneNumberEl);
    isphoneNumberValid = true;
  }
});
passwordEl.addEventListener("input", () => {
  const re = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/;
  if (!re.test(passwordEl.value)) {
    showError(
      passwordEl,
      "Password must have atleast 8 characters including at least 1 lowercase letter, 1 uppercase letter, 1 digit and 1 special character."
    );
  } else {
    hideError(passwordEl);
    isPasswordSecure = true;
  }
});

confirmPasswordEl.addEventListener("input", () => {
  if (!(confirmPasswordEl.value === passwordEl.value)) {
    showError(confirmPasswordEl, "Password doesn't match");
  } else {
    hideError(confirmPasswordEl);
    isConfirmPasswordValid = true;
  }
});
