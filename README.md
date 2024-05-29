# e-tshong an online E-commerce platform.

# Contents.
- [Overview.](#Overview)
- [Learning.](#learning)
- [Tech Stack.](#tech-stack)
- [challenges.](#challenges)
- [A little about the GO language.](#a-little-about-the-go-language)
- [code snippets.](#code-snippets)
- [Acknowledgement.](#acknowledgement)
- [Links.](#links)

## Overview.
The e-tshong project was taken on as part of my Backend web development module for my spring semester (2024). Myself and three other members worked on bringing the project to it's current state. To talk abit about it e-tshong was initailly meant for selling second hand goods but later we realized that not only second hand goods we could also sell first hand or new goods on it.

## Learning.
By developing e-tshong I got an indepth knowledge of the GO language and in general how the web works. I have learned a lot of things like:
- How the browser processes web requests and gives responds.
- How to set up connections between the frontend and the backend or server sides of a web application.
- Handling various kinds of data transfer between the frontend and backend.

## Tech stack
- Frontend:
  - HTML, CSS, Javascript
- Backend:
  - GO
- Database:
  - PostgresSQL
- Deployment:
  - Render 

## Challenges.
Although GO is known for it's simplicity and being a statically typed language compared to other backend languages like Node js but if you are new to coding GO really takes the hell out of you 😮‍💨 (Sorry GO but I mean no offence). Initially being new to the GO language I was hopeless, thought it wasn't possible to build such an application with GO but as I went on working on e-tshong, I came to realize that GO really was an interesting language to learn. That's all about my cahllenges with GO and my initial views on it. Some challenges I faced while building e-tshong:
- Image processing and storing was really challenging in GO(took me a whole day and a half to figure it out).
- Reading and transfering JSON data also seemed to be challenging initially.

## A little about the GO language.
Since this is the first project I have done with GO I want to dedicate this section to talk a little about the language. GO as many people say (I am saying it this way because I don't have much experience with other backend languages) is a powerful programming language created by Robert Griesmer, Rob Pike, and Ken Thomson at Google in 2007. Since then as mentioned by Hitesh Choudary in his GO course on youtube (link is in the link section!) it has been loved by startups and has been used by giants including Meta, Google, Twitter, Netflix and Uber. GO is known for many things like build in concurrency support (which means it can handle multiple tasks at a time), Simplicity, garbage collection and scalabillity. 

If you are just getting into programming like me at this stage it's sure that all the above mentioned benifits of GO will not make sense to you as most of these things also don't make much sense to me at the time I am writing this documentation but as you get a hang of GO things slowly start to make sense and you start to love GO.

## code snippets.
```
func Register(w http.ResponseWriter, r *http.Request) {

	var seller model.Seller
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	hashedPassword, err := httpresp.HashPassword(seller.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	seller.Password = hashedPassword

	// convert contact number to int from string
	saveErr := seller.Create()
	if saveErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpresp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Seller registered"})
}
```

```
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
```

## Acknowledgement
I am greatful to my Backend web development lecturer Mr. Choki Dorji and the authors of the all works I went through in learning GO and building a real web application like e-tshong.

## Links
- [Go Official website.](https://go.dev/)
- [Hitesh Choudary's GO tutorial.](https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&ab_channel=HiteshChoudhary)
- [live server link.](https://e-tshong-platform.onrender.com)
- [Render official page.](https://render.com/)
