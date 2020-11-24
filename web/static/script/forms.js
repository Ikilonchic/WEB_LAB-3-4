function submitForm(event) {
    event.preventDefault();
    
    let formData = new FormData(event.target);
    
    let obj = {};
    formData.forEach((value, key) => {
        if(key === "dob") {
            obj[key] = new Date(value);
        } else {
            obj[key] = value;
        }
    });
    
    let request = new Request(event.target.action, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
        redirect: 'follow',
        body: JSON.stringify(obj)
    });
    
    fetch(request).then((response) => {
        if(response.redirected) {
            window.location.href = response.url;
        }

        return response.json();
    }).then((data) => {
        result = confirm(data.error)
    }).catch((error) => {
        console.log(error)
    });
}

document.getElementById('postform').addEventListener('submit', submitForm);