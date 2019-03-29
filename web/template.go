package web

import "html/template"

// Template
var Template = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

<!-- Bootstrap CSS -->
<link rel="stylesheet"
		href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css"
		integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm"
		crossorigin="anonymous">
<!-- Font Awesome -->
<link rel="stylesheet"
		href="https://use.fontawesome.com/releases/v5.8.1/css/all.css"
		integrity="sha384-50oBUHEmvpQ+1lW4y57PTFmhCaXp0ML5d60M1M7uH2+nqUivzIebhndOJK28anvf"
		crossorigin="anonymous">

<script>  
window.addEventListener("load", function(evt) {

	var output = document.getElementById("output");
	var input = document.getElementById("input");
	var ws;

	$(document).ready(function(){
		// your code
		ws = new WebSocket("{{.}}");
		ws.onopen = function(evt) { print("OPEN");		}
		ws.onclose = function(evt) { print("CLOSE"); ws = null;		}
		ws.onmessage = function(evt) {
			print("RESPONSE: " + evt.data);
			obj = JSON.parse(evt.data);
			document.getElementById("artist_name").innerHTML = obj.artist;
		}
		ws.onerror = function(evt) {
				print("ERROR: " + evt.data);
		}
	});

	var print = function(message) {
		var d = document.createElement("div");
		d.innerHTML = message;
		output.appendChild(d);
	};

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		myJson={"cmd":input.value, "xtra":"haha"}
		print("SEND: " + myJson);
        ws.send(JSON.stringify(myJson));
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>

<!-- JQuery -->
<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"
		integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
		crossorigin="anonymous"></script>

<!-- Popper -->
<script
		src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"
		integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q"
		crossorigin="anonymous"></script>

<!-- Bootstrap JS -->
<script
		src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"
		integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl"
		crossorigin="anonymous"></script>

<div class="container">
  <div class="row">
	<div class="col-sm">
	  <div class="card">
		<div class="card-body">
			<h5 class="card-title" id="artist_name">&nbsp;</h5>
		</div>
	  </div>
	</div>
  </div>
  <div class="row">
	<div class="col-sm">
		<p>
		<button id="backward" class="btn btn-secondary"><i class="fas fa-backward"></i></button>
		<button id="play" class="btn btn-primary"><i class="fas fa-play"></i></button>
		<button id="pause" class="btn btn-warning"><i class="fas fa-pause"></i></button>
		<button id="forward" class="btn btn-secondary"><i class="fas fa-forward"></i></button>
		<button id="stop" class="btn btn-danger"><i class="fas fa-stop"></i></button>
		</p>
	</div>
  </div>
  <div class="row">
	<div class="col-sm">
		<div class="progress">
  			<div class="progress-bar" role="progressbar" style="width: 15%" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100"></div>
		</div>
	</div>
  </div>
  <div class="row">
	<div class="col-sm">
	<form>
		<button id="open">Open</button>
		<button id="close">Close</button>
		<p><input id="input" type="text" value="Hello world!">
		<button id="send">Send</button>
	</form>
	</div>
  </div>
  <div class="row">
	<div class="col-sm">
	  <div id="output"></div>
	</div>
  </div>
</div>
</body>
</html>
`))
