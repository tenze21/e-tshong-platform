const editProfile= document.querySelector(".edit_btn");
const closeEditModel= document.querySelector("#model_close");
const editModel= document.querySelector(".profile_edit_model");

editProfile.addEventListener("click", ()=>{
    editModel.showModal();
    getInitialUserInfo();
})

closeEditModel.addEventListener("click",()=>{
    editModel.close();
})

// getting user details
function getInitialUserInfo(){
    let name= document.querySelector("#name").textContent.split(" ");
    let firstName= name[0];
    let lastName= name[1];
    let emailValue= document.querySelector("#email_value").textContent;
    let contactNumberValue=document.querySelector("#contact_number_value").textContent;
    let genderValue= document.querySelector("#gender").textContent;

    let firstNameEl= document.getElementById("first_name");
    let lastNameEl= document.getElementById("last_name");
    let emailEl= document.getElementById("email");
    let contactNumberEl=document.getElementById("contact_number");
    let genders= document.getElementsByName("gender");

    firstNameEl.value=firstName;
    lastNameEl.value=lastName;
    emailEl.value=emailValue;
    contactNumberEl.value=contactNumberValue;
    genders.forEach(gender=>{
        if(gender.value===genderValue){
            gender.checked=true;
        }
    })
}

//  post delete confirmation model
const deleteBtn=document.querySelectorAll(".delete");
const cancelDelete= document.querySelector(".cancel");
const confirmationModal= document.querySelector(".confirmation_modal");

deleteBtn.forEach(btn=>{
    btn.addEventListener("click", ()=>{
        confirmationModal.showModal();
    })
})
cancelDelete.addEventListener("click", ()=>{
    confirmationModal.close();
})

// upload profile
const profile= document.querySelector(".profile_img");
const profileInput= document.querySelector("#profile");
profileInput.onchange= function(){
    profile.src= URL.createObjectURL(profileInput.files[0]);
};
