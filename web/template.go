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
<style>
/* https://stackoverflow.com/questions/19254146/position-fixed-not-working-in-mobile-browser */
.fixed {
  -webkit-backface-visibility: hidden;
}
</style>
<script> 
window.addEventListener("load", function(evt) {
	var ws;

	var btnPlay = document.getElementById("play");
	var btnPause = document.getElementById("pause");
	var btnResume = document.getElementById("resume");
	var btnStop = document.getElementById("stop");
	var btnNext = document.getElementById("next");
	var btnPrevious = document.getElementById("previous");
	var pause = function() {
		btnPlay.style.display = "none";
		btnPause.style.display = "none";
		btnResume.style.display = "";
	  btnStop.style.display = "";		btnStop.disabled = "";
	  btnNext.style.display = "";		btnNext.disabled = "";
		btnPrevious.style.display = ""; btnPrevious.disabled = "";
	};

	var play = function() {
		btnPlay.style.display = "none";
		btnPause.style.display = "";
		btnResume.style.display = "none";
	  btnStop.style.display = "";		btnStop.disabled = "";
	  btnNext.style.display = "";		btnNext.disabled = "";
		btnPrevious.style.display = ""; btnPrevious.disabled = "";
	};

	var stop = function() {
		btnPlay.style.display = "";
		btnPause.style.display = "none";
		btnResume.style.display = "none";
	  btnStop.style.display = "";		btnStop.disabled = "disabled";
	  btnNext.style.display = "";		btnNext.disabled = "disabled";
		btnPrevious.style.display = ""; btnPrevious.disabled = "disabled";
	};

	var currentSongTitle = document.getElementById("csTitle");
	var currentSongArtist = document.getElementById("csArtist");
	var currentSongAlbumArtist = document.getElementById("csAlbumArtist");
	var currentSongAlbum = document.getElementById("csAlbum");
	var currentSongReleased = document.getElementById("csReleased");

	var currentPlaylist = document.getElementById("cpContainer");

	var progressBar = document.getElementById("progressBar");
	var progressLabel = document.getElementById("progressLabel");
	var duration = 1.0;
	var elapsed = 0.0;
	var state = "pause";
	var readableSeconds = function(value) {
		var min = parseInt(value/60);
		var sec = parseInt(value % 60);
		if (sec < 10) { sec = "0" + sec; }
		return min + ":" + sec;
	};
	var updateProgressBar = function() {
		// console.log("progress: " + duration + " / " + elapsed);
		progressBar.style.width = (elapsed/duration*100) + "%";
		progressLabel.innerHTML = readableSeconds(elapsed) + "/" + readableSeconds(duration);
		if ((state=="play") && (elapsed<duration)) { elapsed += 1.0; }
		setTimeout(updateProgressBar, 1000);
	};

	$(document).ready(function(){
		ws = new WebSocket("{{.ws}}");
		ws.onopen = function(evt) { console.log("OPEN");}
		ws.onclose = function(evt) { console.log("CLOSE"); ws = null;	}
		ws.onmessage = function(evt) {
			console.log("RESPONSE: " + evt.data);
			obj = JSON.parse(evt.data);
			if (obj.type == {{.string}}) {
			} else if (obj.type == {{.status}}) {
				if (obj.data.state == "pause") {
					pause();
					state = "pause";
					duration = obj.data.duration;
					elapsed = obj.data.elapsed;
				} else if (obj.data.state == "play") {
					play();
					state = "play";
					duration = obj.data.duration;
					elapsed = obj.data.elapsed;
				} else if (obj.data.state == "stop") {
					stop();
					state = "stop";
					duration = 1.0;
					elapsed = 0.0;
				}
			} else if (obj.type == {{.currentSong}}) {
				currentSongArtist.innerHTML = obj.data.artist;
				currentSongTitle.innerHTML = obj.data.title;
				if (obj.data.album_artist != "") {
					currentSongAlbumArtist.innerHTML = obj.data.album_artist;
				} else {
					currentSongAlbumArtist.innerHTML = obj.data.artist;
				}
				currentSongAlbum.innerHTML = obj.data.album;
				if (obj.data.released != "") {
					currentSongReleased.innerHTML = "&nbsp;(" + obj.data.released + ")";
				} else {
					currentSongReleased.innerHTML = "";
				}
			} else if (obj.type == {{.currentPlaylist}}) {
				console.log("currentPlaylist")
				var list = "";
				for (var i=0; i<obj.data.Playlist.length; i++) {
//					list += '<div class="row">'; // row 1
//					list += '<div class="col-xl p-1">'; // col-xl
					list += '<div class="d-flex p-1 bd-highlight flex-nowrap">'; // row 2
					list += '<div class="d-flex align-middle">';
					list += '<button id="plPlay' + i + '" class="btn btn-outline-primary btn-sm"><i class="fas fa-play"></i></button>';
					list += '&nbsp;';
					list += '<button id="plRemove' + i + '" class="btn btn-outline-danger btn-sm"><i class="fas fa-minus"></i></button>';
					list += '</div>';
					list += '<div class="d-flex p-1 align-middle">'; // row 2
					list += obj.data.Playlist[i].artist;
					list += '&nbsp;&ndash;&nbsp;';
					list += obj.data.Playlist[i].title;
					list += '&nbsp;&ndash;&nbsp;';
					list += obj.data.Playlist[i].album;
					list += '</div></div>';
	//				list += '</div>'; // col-xl
//					list += '</div>'; // row 1
				}
				currentPlaylist.innerHTML = list;
				for (var i=0; i<obj.data.Playlist.length; i++) {
					{
						const j=i;
						document.getElementById("plRemove" + j).onclick = function(evt) {
							return command("remove" + j);
						};
					}
				}
			}
		}
		ws.onerror = function(evt) {
				console.log("ERROR: " + evt.data);
		}
		window.onfocus = function(event) {
			// request a fresh status as some browsers (e. g. Chrome) suspend out
			// progress bar setTimeout functions
			command("status");
		}
		updateProgressBar();
	});

	var command = function(cmd) {
        if (!ws) { return false; }
		myJson={"command":cmd}
		console.log("SEND: " + myJson);
        ws.send(JSON.stringify(myJson));
        return false;
	}

	// add onclick function for all controls
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
  <div class="row sticky-top bg-light fixed">
		<div class="col-md order-2">
			<div class="container-fluid">
				<div class="row">
					<div class="col-auto p-1">
						<h6>
						<span id="csArtist"></span>&nbsp;&ndash;&nbsp;
						<span id="csTitle">&nbsp;</span>
						</h6>
					</div>
					<div class="col p-1 text-center">
						<h6 id="progressLabel" class="text-muted">&nbsp;</h6>
					</div>
					<div class="col-auto p-1">
						<h6>
							<span id="csAlbumArtist"></span>&nbsp;&ndash;&nbsp;
							<span id="csAlbum"></span>
							<span id="csReleased"></span>
						</h6>
					</div>
				</div>
				<div class="row">
					<div class="col p-1">
						<div class="progress">
							<div id="progressBar" class="progress-bar" role="progressbar" 
								style="width: 0%; transition: width 1s ease-in-out"></div>
						</div>
					</div>
				</div>
			</div>
		</div><!-- col -->

		<div class="col-sm-auto order-1 text-center p-1">
			<button id="play" class="btn btn-outline-primary">&nbsp;<i class="fas fa-play"></i>&nbsp;</button>
			<button id="resume" class="btn btn-outline-warning">&nbsp;<i class="fas fa-play"></i>&nbsp;</button>
			<button id="pause" class="btn btn-outline-warning">&nbsp;<i class="fas fa-pause"></i>&nbsp;</button>
			<button id="previous" class="btn btn-outline-secondary">&nbsp;<i class="fas fa-backward"></i>&nbsp;</button>
			<button id="next" class="btn btn-outline-secondary">&nbsp;<i class="fas fa-forward"></i>&nbsp;</button>
			<button id="stop" class="btn btn-outline-danger">&nbsp;<i class="fas fa-stop"></i>&nbsp;</button>
		</div><!-- col -->
	</div><!-- row -->
	<div id="cpContainer"></div>
</div>
</body>
</html>
`))
