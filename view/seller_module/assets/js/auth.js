if (document.cookie == "") {
  alert(
    "Cookie expired or Could not find cookie. Please login again to renew your session."
  );
  window.open("index.html", "_self");
} else {
  console.log("cookie set");
}

function logout() {
  fetch("/logout")
    .then((response) => {
      if (response.ok) {
        window.open("/login.html", "_self");
      } else {
        throw new Error(response.statusText);
      }
    })
    .catch((e) => {
      alert(e);
    });
}
