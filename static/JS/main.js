
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
  
function OnclickCreatePost(){
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
}