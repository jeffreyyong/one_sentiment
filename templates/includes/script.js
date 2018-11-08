const pollResult = (uuid) => {
    let url = window.location.hostname + '/result/' + uuid
    fetch(url)
    .then(data=>{
        document.getElementById('result').value = data.result
    })
    .catch(err=>{ console.log(err) })
}

const sendParameters = () => {
    let language = document.getElementById("language").value
    let number = document.getElementById("destination").value

    let url ='/asr'
    let param = {
        method: 'POST',
        body:{
            language: language,
            number: number
        }
    }
    console.log('backend url ' + url)
    fetch(url, param)
    .then(data=>{
        console.log('response from /asr' + data)
        pollResult(data.UUID)
    })
    .catch(err=>{ console.log('/asr error: ' + err) })
}

const main = () => {
    document.getElementById('sendButton').addEventListener("click", sendParameters)
}

window.addEventListener("load", main)