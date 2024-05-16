const form = document.querySelector("#registration_form");
const firstNameEl = document.querySelector("#first_name");
const lastNameEl = document.querySelector("#last_name");
const phoneNumberEl = document.querySelector("#contact_number");
const passwordEl = document.querySelector("#password");
const confirmPasswordEl = document.querySelector("#confirm_password");

let isFirstNameValid = false,
  isLastNameValid= false,
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
    location.reload();
  }
});

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

confirmPasswordEl.addEventListener('input', ()=>{
    if(!(confirmPasswordEl.value===passwordEl.value)){
        showError(confirmPasswordEl, "Password doesn't match")
    }else{
        hideError(confirmPasswordEl)
        isConfirmPasswordValid=true;
    }
})
