
    // document.getElementById("testPost").value = "yay"
    // fetch("")
    // .then((res) => {
    //     return res.json
    // })
    // .then((data) => {
    //     console.log(data)
    //     document.getElementById("testPost").value = "yay"
    // }) 


fetch("/GetPosts")
.then((res) => res.json())
.then((json) => {
    console.log("response", json)
    json.forEach(element => {
        const newPost = document.createElement("div");
        const title = document.createElement("h1");
        const content = document.createElement("p");
        const divReactions = document.createElement("div");
        const divReactionStyle = document.createElement("div");
        const likeButton = document.createElement("button");
        const comentary = document.createElement("button");

        console.log(element)
        newPost.className = "posts";
        title.textContent = element.Title;
        content.textContent = element.Content;
        divReactions.className = "reaction";
        divReactionStyle.style.backgroundColor = "beige";
        divReactionStyle.style.borderRadius = "10px";
        divReactionStyle.style.display = "flex";
        divReactionStyle.style.justifyContent = "space-around"
        likeButton.textContent = "❤️";
        comentary.textContent = "Commentaire";

        divReactionStyle.appendChild(likeButton);
        divReactionStyle.appendChild(comentary);
        divReactions.appendChild(divReactionStyle);
        newPost.appendChild(title);
        newPost.appendChild(content);
        newPost.appendChild(divReactions);

        document.getElementsByClassName("post")[0].children[0].appendChild(newPost);



    });


})
function Afficher()
{ 
let input = document.getElementById("motdepasse"); 
if (input.type === "password")
{ 
input.type = "text"; 
} 
else
{ 
input.type = "password"; 
} 
} 

let idActualUser = 0 
   
   
   fetch("")
    .then((res) => {
        return res.json
    })
    .then((data) => {
        console.log(data)
    }) 


  
function OnclickCreatePost(){
    idActualUser = 1;
    console.log("cliquer")
    if (idActualUser != 0){
        console.log("OnclickCreatepost : OK")

    fetch("/CreatePost", {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            userId : idActualUser,
            title: document.getElementById("postTitle").value,
            content: document.getElementById("contentPost").value,
            category: document.getElementById("theme-select").value
        })
    })
    .then((res) => {
        res.json()
    })
    .then((data) => {
        // console.log(data);

        // if (!!data.error){
        //     document.getElementById("errorPost").innerText = data.error
        //     return
        // }
    })
    .catch((resp) => console.log(resp))
    }else{
        console.log("nop")
    }
    
}