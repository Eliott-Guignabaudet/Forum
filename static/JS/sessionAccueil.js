

const connected = () => {


    document.getElementById("registerButton").remove()
    document.getElementById("loginButton").remove()
    const newButton = document.createElement("button")
    newButton.setAttribute("class","inscr")
    newButton.textContent = "Déconnexion"
    newButton.id = "deconnexion"
    newButton.addEventListener("click", logOut, false)
    document.getElementsByClassName("inscriptionConnexion")[0].appendChild(newButton)
}


const logOut = ()=> {
    fetch("/Logout")
    .then(res => res.json())
    .then(JSON => {
        document.location.reload()
    })
}

fetch("/GetSession")
.then(res => res.json())
.then(JSON =>{
    if (!JSON.resp){
        connected()
        return
    }
    document.getElementById("popuplink").remove()
})

