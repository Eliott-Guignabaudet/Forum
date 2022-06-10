let globalJsonPost;
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
        const comentary = document.createElement("button");

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