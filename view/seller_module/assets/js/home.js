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

window.onload = function() {
    fetch("/products")
        .then((response) => response.text()) // parse the JSON directly
        .then((data) => showProducts(data))
        .catch((error) => console.error('Error fetching products:', error));
}

function showProducts(data){
    const products=JSON.parse(data);
    products.forEach(product=>newProduct(product))
}

const main = document.querySelector("main");

function newProduct(product) {
    // Create card container
    const container = document.createElement("div");
    container.classList.add("card_container");
    main.appendChild(container);

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

    // Create and append view more button
    const button = document.createElement("button");
    button.textContent = "View More";
    button.onclick = () => showProductDetail(product.productid);
    cardBody.appendChild(button);
}

function showProductDetail(pid){
    // Add the product id to the URL parameter
    const url = `product.html?pid=${pid}`;
    window.open(url, "_self");
}