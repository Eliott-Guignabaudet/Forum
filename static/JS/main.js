let globalJsonPost;
let postid = 0
    // document.getElementById("testPost").value = "yay"
    // fetch("")
    // .then((res) => {
    //     return res.json
    // })
    // .then((data) => {
    //     console.log(data)
    //     document.getElementById("testPost").value = "yay"
    // }) 
const searchBar = document.getElementsByClassName("recherche")[0];
searchBar.addEventListener(onkeydown, function(){searchPosts(searchBar.value)});

const displayPosts = (posts) =>{
    const postsDisplayed = document.querySelectorAll(".post .posts");
    console.log("displayed:")
    for (const post of postsDisplayed) {
        console.log(post)
        post.remove();
    }


    posts.forEach(element => {
        const newPost = document.createElement("div");
        const title = document.createElement("h1");
        const content = document.createElement("p");
        const divReactions = document.createElement("div");
        const divReactionStyle = document.createElement("div");
        const likeButton = document.createElement("button");
        const hrefpopup = document.createElement("a")
        hrefpopup.setAttribute("href","#popupCom")
        const comentary = document.createElement("button");
        comentary.setAttribute("onclick",`RecupererIdComms(${element.Id})`)

        hrefpopup.append(comentary)
        newPost.setAttribute("id",String(element.Id))
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
        divReactionStyle.appendChild(hrefpopup)
        divReactions.appendChild(divReactionStyle);
        newPost.appendChild(title);
        newPost.appendChild(content);
        newPost.appendChild(divReactions);

        document.getElementsByClassName("post")[0].children[0].appendChild(newPost);
    });
}

const displayComms = (comments) =>{
    const commsDisplayed = document.querySelectorAll(".comment .comments");
    console.log("displayed:")
    for (const comment of commsDisplayed) {
        console.log(post)
        comment.remove();
    }


    comments.forEach(element => {
        const newComms = document.createElement("div");
        const content = document.createElement("p");
        const divReactions = document.createElement("div");
        const divReactionStyle = document.createElement("div");
        newComms.setAttribute("id",String(element.Id))
        newComms.className = "comment";
        content.textContent = element.Content;
        divReactions.className = "reaction";
        divReactionStyle.style.backgroundColor = "beige";
        divReactionStyle.style.borderRadius = "10px";
        divReactionStyle.style.display = "flex";
        divReactionStyle.style.justifyContent = "space-around"

        divReactions.appendChild(divReactionStyle);
        newComms.appendChild(content);
        newComms.appendChild(divReactions);

        document.getElementsByClassName("1postComms")[0].appendChild(newComms);
    });
}

const searchPosts = (searchValue) => {
    console.log("search!")
    const postFiltered = globalJsonPost.filter(post => {
        if (post.Content.includes(searchValue)){
            return true
        }else if (post.Title.includes(searchValue)){
            return true
        }
        return false
    })
    console.log(postFiltered)
    displayPosts(postFiltered)
}

fetch("/GetPosts")
.then((res) => res.json())
.then((json) => {
    console.log("response", json)
    displayPosts(json);
    globalJsonPost = json;

})

let input = document.getElementById("motdepasse"); 
if (input.type === "password")
{ 
input.type = "text"; 
} 
else
{ 
input.type = "password"; 
} 


   
   
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
        return res.json()
    })
    .then((data) => {
        console.log("data:",data);

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


function ChangeBackgroundcolorV(){
    document.body.style.backgroundColor = "green"
}

function ChangeBackgroundcolorS(){
    document.body.style.backgroundColor = "brown"
}

function ChangeBackgroundcolorT(){
    document.body.style.backgroundColor = "gray"
}

function ChangeBackgroundcolorJ(){
    document.body.style.backgroundColor = "blue"
}
function ChangeBackgroundcolorR(){
    document.body.style.backgroundColor = "blueviolet"
}
function ChangeBackgroundcolorH(){
    document.body.style.backgroundColor = "violet"
}

function ChangeBackgroundcolorM(){
    document.body.style.backgroundColor = "purple"
}
function ChangeBackgroundcolorA(){
    document.body.style.backgroundColor = "whitesmoke"
}

function RecupererIdComms(id){
    postid = id
    console.log("id :",id)
    console.log("hey post id :",postid)
    idActualUser = 1;
    fetch("/GetComms", {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            PostId: id
        })

    })
    .then((res) => res.json())
    .then((json) => {
    console.log("response", json)
    displayComms(json)

})
    //afficher les comms dans la div selon l'id avec une route
}

function OnclickCreateComm(){
    idActualUser = 1;
    console.log("cliquer")
    if (idActualUser != 0){
        console.log("OnclickCreateComms : OK")

    fetch("/CreateComms", {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            userId : idActualUser,
            PostId : postid,
            content: document.getElementById("contentComms").value,
        })
    })
    .then((res) => {
        return res.json()
    })
    .then((data) => {
        console.log("data:",data);

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

// function filter(){
//     document.getElementsByClassName()
// }
