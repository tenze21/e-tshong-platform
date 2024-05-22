window.onload=function(){
    // Extract Prpduct ID from the URL
    const urlParams=new URLSearchParams(window.location.search);
    const productid=urlParams.get('pid');

    // Check if productId is present
    if(productid){
        fetch(`/product/${productid}`)
        .then(res=>res.text())
        .then(data=>showProductDetails(data))
        .catch(error=>console.error("Error fetching product details:", error));
    }else{
        console.error('Product ID not in the URL');
    }
}

function showProductDetails(product){
    // get HTML elements to insert product Data
    const p=JSON.parse(product);
    const pimg1=document.querySelector('#productimg1');
    const pimg2=document.querySelector('#productimg2');
    const pimg3=document.querySelector('#productimg3');
    const pimg4=document.querySelector('#productimg4');
    const ptitle=document.querySelector(".product_title");
    const aprice=document.querySelector("#actual_price");
    const sprice=document.querySelector("#new_price");
    const sname=document.querySelector('#name');
    const semail=document.querySelector('#email');
    const snumber=document.querySelector('#contact_number');
    const pdescription=document.querySelector("#product_description");

    pimg1.src=`data:image/jpeg;base64,${p.productimg1}`;
    pimg2.src=`data:image/jpeg;base64,${p.productimg2}`;
    pimg3.src=`data:image/jpeg;base64,${p.productimg3}`;
    pimg4.src=`data:image/jpeg;base64,${p.productimg4}`;
    ptitle.textContent=p.ptitle;
    aprice.textContent=p.aprice;
    sprice.textContent=p.sprice;
    sname.textContent=`${p.first_name} ${p.last_name}`;
    semail.textContent=p.email;
    snumber.textContent=p.contact_number;
    pdescription.textContent=p.pdescription;
}