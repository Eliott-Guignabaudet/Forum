fetch("/GetSession")
.then(res => res.json())
.then(JSON =>{

    if (!JSON.resp){
        connected()
    }
    console.log(JSON)
})