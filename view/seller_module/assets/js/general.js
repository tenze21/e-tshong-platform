const greetingText= document.getElementById("greeting");
const date=new Date;
const hour=date.getHours()
if(hour>=12 && hour<=17){
    greetingText.textContent="Good Afternoon"
}else if(hour>=18 && hour<=24){
    greetingText.textContent="Good Evening"
}else{
    greetingText.textContent="Good Morning"
}