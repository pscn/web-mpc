package web

import "html/template"

// Template the one and only Template
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
	var ws;

	var pause = function() {
		document.getElementById("play").style.display = "none";
		document.getElementById("pause").style.display = "none";
		document.getElementById("resume").style.display = "";
		document.getElementById("stop").style.display = "";
		document.getElementById("next").style.display = "";
		document.getElementById("previous").style.display = "";
	};

	var play = function() {
		document.getElementById("play").style.display = "none";
		document.getElementById("pause").style.display = "";
		document.getElementById("resume").style.display = "none";
		document.getElementById("stop").style.display = "";
		document.getElementById("next").style.display = "";
		document.getElementById("previous").style.display = "";
	};

	var stop = function() {
		document.getElementById("play").style.display = "";
		document.getElementById("pause").style.display = "none";
		document.getElementById("resume").style.display = "none";
		document.getElementById("stop").style.display = "none";
		document.getElementById("next").style.display = "none";
		document.getElementById("previous").style.display = "none";
	};
	var currentSongTitle = document.getElementById("cs-title");
	var currentSongArtist = document.getElementById("cs-artist");
	var currentSongAlbum = document.getElementById("cs-album");

	$(document).ready(function(){
		ws = new WebSocket("{{.ws}}");
		ws.onopen = function(evt) { console.log("OPEN");}
		ws.onclose = function(evt) { console.log("CLOSE"); ws = null;	}
		ws.onmessage = function(evt) {
			console.log("RESPONSE: " + evt.data);
			obj = JSON.parse(evt.data);
			if (obj.type == {{.string}}) {
			} else if (obj.type == {{.status}}) {
				if (obj.data.state == "pause") { pause(); }
				else if (obj.data.state == "play") { play(); }
				else if (obj.data.state == "stop") { stop(); }
			} else if (obj.type == {{.currentSong}}) {
				currentSongTitle.innerHTML = obj.data.title;
				currentSongArtist.innerHTML = obj.data.artist;
				currentSongAlbum.innerHTML = obj.data.album;
			}
		}
		ws.onerror = function(evt) {
				console.log("ERROR: " + evt.data);
		}
	});

	var command = function(cmd) {
        if (!ws) { return false; }
		myJson={"command":cmd}
		console.log("SEND: " + myJson);
        ws.send(JSON.stringify(myJson));
        return false;
	}

	// FIXME: loop over array with play, pause etc...
	var controls = ["play", "resume", "pause", "stop", "next", "previous"]
	controls.forEach(activator);
	function activator(value) {
			document.getElementById(value).onclick = function(evt) {
			return command(value);
		};
	}

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

<div class="container-fluid">
  <div class="row">
	<div class="col-sm">
		<p>
		<button id="play" class="btn btn-primary">&nbsp;<i class="fas fa-play"></i>&nbsp;</button>
		<button id="resume" class="btn btn-warning">&nbsp;<i class="fas fa-play"></i>&nbsp;</button>
		<button id="pause" class="btn btn-warning">&nbsp;<i class="fas fa-pause"></i>&nbsp;</button>
		<button id="previous" class="btn btn-secondary">&nbsp;<i class="fas fa-backward"></i>&nbsp;</button>
		<button id="next" class="btn btn-secondary">&nbsp;<i class="fas fa-forward"></i>&nbsp;</button>
		<button id="stop" class="btn btn-danger">&nbsp;<i class="fas fa-stop"></i>&nbsp;</button>
		</p>
	</div>
  </div>

  <div class="row">
	<div class="col-sm">
	  <div class="card">
		<div class="card-body">
			<h5 class="card-title" id="cs-title">&nbsp;</h5>
			<p class="card-text">
				<span id="cs-artist"></span>&nbsp;&ndash;&nbsp;
				<span id="cs-album"></span>
			</p>
		</div>
	  </div>
	</div>
  </div>
  <div class="row">
	<div class="col-sm">
		<div class="progress">
  			<div class="progress-bar" role="progressbar" style="width: 15%" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100"></div>
		</div>
	</div>
  </div>
</div>
</body>
</html>
`))
