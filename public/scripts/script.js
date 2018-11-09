const modal = document.getElementById('myModal')
let stopPinging = {}
let currentUUID = ""

const pollResult = async (uuid) => {
    let url ='/result/' + uuid
    currentUUID = uuid

    if (stop.uuid) {
        return
    }

    try {
        let response = await fetch(url)

        let data = await response.json()
        console.log(JSON.stringify(data))
        document.getElementById("result").innerHTML = data.result
    } catch(err) {
        window.setTimeout(() => {
            pollResult(uuid)
        }, 1000)
        console.log(err)
    }
}

const sendParameters = async () => {
    let language = document.getElementById("language").value
    let number = document.getElementById("destination").value

    if (number === "") {
        return false
    }

    modal.style.display = "block"
    document.getElementById('destination').value = ""

    let body = {
            language: language,
            destination: number
        }
    let url ='/asr'
    let param = {
        headers: {
            "Content-Type": "application/json; charset=utf-8",
        },
        method: 'POST',
        body: JSON.stringify(body)
    }
    
    let response = await fetch(url, param)
    let data = await response.json()
    console.log(JSON.stringify(data))

    if (data.uuid !== "") {
        pollResult(data.uuid)
    }

    return true
}

const keypress = e => {
    var key = e.which || e.keyCode;
    if (key === 13) { // 13 is enter
        sendParameters()
    }
}

const cleanPage = () => {
    stop.currentUUID = true
    modal.style.display = ""
    result = document.getElementById("result")
    result.innerText = ""
    img = result.appendChild(document.createElement('img'))
    img.src = "/public/Images/Loading_icon.gif"
    img.alt = "Loading"
}

const main = () => {
    let btn = document.getElementById("sendButton")
    // Get the <span> element that closes the modal
    let span = document.getElementsByClassName("close")[0]

    btn.addEventListener("click", sendParameters)

    document.getElementById('destination').addEventListener('keypress', keypress)
    document.getElementById('language').addEventListener('keypress', keypress)

    // When the user clicks on <span> (x), close the modal
    span.onclick = () => {
        cleanPage()
    }

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = (event) => {
        if (event.target == modal) {
            cleanPage()
        }
    }
}

window.addEventListener("load", main)