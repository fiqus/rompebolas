(function(){
    let exitButton = document.getElementById("exitButton");
    let saveButton = document.getElementById("saveButton");
    let textArea = document.getElementById("text");
    let overlayNo = document.getElementById("overlay_no");
    let overlayOk = document.getElementById("overlay_ok");
    let overlayCount = 0;

    function invoke(type, value){
        window.external.invoke(JSON.stringify({type, value}))
    }

    function showOverlayNo(){
        overlayNo.className="";
        window.setTimeout(function(){overlayNo.className="hidden"}, 2000);
    }

    function exit(){
        overlayOk.className="";
        window.setTimeout(function(){invoke("exit");}, 500);
    }

    exitButton.onclick = function(){
        if (overlayCount < 2){
            showOverlayNo();
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
                    showOverlayNo();
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