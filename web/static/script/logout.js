function logout(event) {
    let request = new Request({
        method: 'GET',
        credentials: 'same-origin',
        redirect: 'follow',
    });
    
    fetch("/logout", request).then((response) => {
        if(response.redirected) {
            window.location.href = response.url;
        }
    }).catch((error) => {
        console.log(error)
    });
}