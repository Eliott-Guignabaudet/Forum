
    document.getElementById("testPost").value = "yay"
    fetch("")
    .then((res) => {
        return res.json
    })
    .then((data) => {
        console.log(data)
        document.getElementById("testPost").value = "yay"
    }) 
  
function OnclickCreatePost(){
    console.log("OnclickCreatepost : OK")
    fetch("/CreatePost", {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            postTitle: document.getElementById("postTitle").value,
            postContent: document.getElementById("contentPost").value,
            postTheme: document.getElementById("theme-select").value
        })
    })
}