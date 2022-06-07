
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
    if (idActualUser != 0){
        console.log("OnclickCreatepost : OK")
    fetch("/CreatePost", {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            userId : 1,
            title: document.getElementById("postTitle").value,
            content: document.getElementById("contentPost").value,
            category: document.getElementById("theme-select").value
        })
    })
    .then((res) => {
        res.json()
    })
    .then((data) => {
        if (!!data.error){
            document.getElementById("errorPost").innerText = data.error
            return
        }
    })
    }else{
        console.log("nop")
    }
    
}