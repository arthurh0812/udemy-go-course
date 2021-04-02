const loginForm = document.querySelector(".login-form");

if (loginForm) {
  loginForm.addEventListener("submit", function () {
    const email = document.querySelector("#email").value;
    const password = document.querySelector("#password").value;

    console.log(email, password);
    // fetch()
  })
}