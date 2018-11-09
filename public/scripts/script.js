const pollResult = async (uuid) => {
    let url ='/result/' + uuid

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
    console.log('backend url ' + url)

    let response = await fetch(url, param)
    let data = await response.json()
    console.log(JSON.stringify(data))
    pollResult(data.uuid)
}

const main = () => {
    document.getElementById('sendButton').addEventListener("click", sendParameters)

    window.addEventListener('keypress', function (e) {
        var key = e.which || e.keyCode;
        if (key === 13) { // 13 is enter
            sendParameters
        }
    });
}

window.addEventListener("load", main)