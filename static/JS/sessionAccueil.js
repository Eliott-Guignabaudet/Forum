const connected = () => {
    document.getElementsByClassName("envoyerpost")[0].remove()

    document.getElementById("registerButton").remove()
    document.getElementById("loginButton").remove()
    newButton = document.createElement("button")
    newButton.class = "inscr"
    newButton.textContent = "Déconnexion"
    document.getElementsByClassName("inscriptionConnexion")[0].appendChild(newButton)
}


fetch("/GetSession")
.then(res => res.json())
.then(JSON =>{
    if (!JSON.resp){
        connected()
    }
    console.log(JSON)
})