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