function login(){
    var data={
       cnumber: parseInt(document.getElementById("phone_number").value),
       password: document.getElementById("password").value
    };
    fetch("/login",{
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-Type":"application/json; charset=UTF-8"},
    }).then((response)=>{
        if(response.ok){
            window.open("seller_module/home.html","_self");
        }else{
            throw new Error(response.statusText);
        }
    })
    .catch((e)=>{
        if(e=="Error: Unauthorized"){
            alert(e + ". User not found!");
            return
        }
    })
}

const form= document.querySelector("form");
form.addEventListener("submit", (e)=>{
    e.preventDefault();
    login()
});