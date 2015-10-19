// sending data to server
var modal = document.querySelector('#openModal');
modal.addEventListener('click', function(e){
    if ((e.target.id === 'modal-submit') && (e.target.textContent !== '')) {
        var modalTweet = document.querySelector('#modal-tweet')
        // make your ajax post
        var msg = modalTweet.value;
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/tweet");
        var json = JSON.stringify({Message: msg});
        xhr.send(json);
        // clear entry
        modalTweet.value = '';
        // receive data back from the post
        xhr.addEventListener('readystatechange', function () {
            if ((xhr.status === 200) && (xhr.readyState === 4)) {
                location.reload(false);
            }
        })
        //
    }
}, false);

var opensModal = document.querySelector('#opens-modal');
opensModal.addEventListener('click', function(e){
    window.setTimeout(function(){
        var tweet = document.querySelector('#modal-tweet');
        tweet.value = '';
        tweet.focus();
    }, 100);
}, false);

/*
checkout for the future:
 http://caniuse.com/#search=xmlhttp

 stream and post blobs
 video, etc
*/