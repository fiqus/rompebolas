(function(){
    let exitButton = document.getElementById("exitButton");
    let saveButton = document.getElementById("saveButton");
    let textArea = document.getElementById("text");
    let overlay = document.getElementById("overlay");
    let overlaySpan = document.getElementById("overlaySpan");
    let overlayCount = 0;

    function invoke(type, value){
        window.external.invoke(JSON.stringify({type, value}))
    }

    function showOverlay(){
        overlay.className="";
        window.setTimeout(function(){overlay.className="hidden"}, 2000);
    }

    function exit(){
        overlaySpan.innerText = "OK";
        overlay.className="ok";
        window.setTimeout(function(){invoke("exit");}, 1000);
    }

    exitButton.onclick = function(){
        if (overlayCount < 2){
            showOverlay();
            overlayCount++;
        }else{
            exit()
        }
    };

    saveButton.onclick = function(){
        invoke("submit", textArea.value.trim());
        exit();
    };

    textArea.oninput = function () {
        if(textArea.value.trim()){
            saveButton.removeAttribute("disabled");
        }else{
            saveButton.setAttribute("disabled", "disabled")
        }
    };

    textArea.onkeypress = function (evt) {
        if (evt.keyCode === 13 && evt.shiftKey) {
            if(textArea.value.trim()){
                invoke("submit", textArea.value.trim());
                exit();
            }else{
                if (overlayCount < 2){
                    showOverlay();
                    overlayCount++;
                }else{
                    exit()
                }
            }
            evt.preventDefault();
            return false;
        }
    }
})();