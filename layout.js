window.addEventListener("load", function (evt) {
  var ws;

  var btnPlay = document.getElementById("play");
  var btnPause = document.getElementById("pause");
  var btnResume = document.getElementById("resume");
  var btnStop = document.getElementById("stop");
  var btnNext = document.getElementById("next");
  var btnPrevious = document.getElementById("previous");
  var pause = function () {
    btnPlay.style.display = "none";
    btnPause.style.display = "none";
    btnResume.style.display = "";
    btnStop.style.display = ""; btnStop.disabled = "";
    btnNext.style.display = ""; btnNext.disabled = "";
    btnPrevious.style.display = ""; btnPrevious.disabled = "";
  };

  var play = function () {
    btnPlay.style.display = "none";
    btnPause.style.display = "";
    btnResume.style.display = "none";
    btnStop.style.display = ""; btnStop.disabled = "";
    btnNext.style.display = ""; btnNext.disabled = "";
    btnPrevious.style.display = ""; btnPrevious.disabled = "";
  };

  var stop = function () {
    btnPlay.style.display = "";
    btnPause.style.display = "none";
    btnResume.style.display = "none";
    btnStop.style.display = ""; btnStop.disabled = "disabled";
    btnNext.style.display = ""; btnNext.disabled = "disabled";
    btnPrevious.style.display = ""; btnPrevious.disabled = "disabled";
  };

  var csTitle = document.getElementById("csTitle");
  var csArtist = document.getElementById("csArtist");
  var csAlbumArtist = document.getElementById("csAlbumArtist");
  var csAlbum = document.getElementById("csAlbum");
  var csElapsed = document.getElementById("csElapsed");
  var csDuration = document.getElementById("csDuration");

  var currentPlaylist = document.getElementById("playlist");

  var duration = 1.0;
  var elapsed = 0.0;
  var state = "pause";
  var readableSeconds = function (value) {
    var min = parseInt(value / 60);
    var sec = parseInt(value % 60);
    if (sec < 10) { sec = "0" + sec; }
    return min + ":" + sec;
  };
  var updateProgress = function () {
    // console.log("progress: " + duration + " / " + elapsed);
    csElapsed.innerHTML = readableSeconds(elapsed);
    csDuration.innerHTML = readableSeconds(duration);
    if ((state == "play") && (elapsed < duration)) { elapsed += 1.0; }
    setTimeout(updateProgress, 1000);
  };

  ws = new WebSocket("ws://192.168.0.111:8080/echo");
  ws.onopen = function (evt) { console.log("OPEN"); }
  ws.onclose = function (evt) { console.log("CLOSE"); ws = null; }
  ws.onmessage = function (evt) {
    console.log("RESPONSE: " + evt.data);
    obj = JSON.parse(evt.data);
    if (obj.type == 1) {
    } else if (obj.type == 2) {
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
        duration = 666.0;
        elapsed = 0.0;
      }
    } else if (obj.type == 3) {
      csArtist.innerHTML = obj.data.artist;
      csTitle.innerHTML = obj.data.title;
      if (obj.data.artist != obj.data.album_artist) {
        csAlbumArtist.style.display = "";
        csAlbumArtist.innerHTML = "[" + obj.data.album_artist + "]&nbsp;";
      } else {
        csAlbumArtist.style.display = "none";
      }
      csAlbum.innerHTML = obj.data.album;
    } else if (obj.type == 4) {
      console.log("currentPlaylist")
      currentPlaylist.innerHTML = "";
      var list = "";
      for (var i = 0; i < obj.data.Playlist.length; i++) {
        var playlistEntry = document.getElementById("playlistEntry")
        var node = playlistEntry.cloneNode(true);
        node.id = "plRow" + i;
        node.style.display = "";
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#plTitle").innerHTML = obj.data.Playlist[i].title;
        node.querySelector("#plAlbum").innerHTML = obj.data.Playlist[i].album;
        if (obj.data.Playlist[i].artist != obj.data.Playlist[i].album_artist) {
          node.querySelector("#plAlbumArtist").innerHTML = "[" + obj.data.Playlist[i].album_artist + "]&nbsp;";
        } else {
          node.querySelector("#plAlbumArtist").style.display = "none";
        }
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        {
          const j = i;
          node.querySelector("#plRemove").onclick = function (evt) {
            return command("remove" + j);
          };
        }
        console.log(node.innerHTML);
        currentPlaylist.append(node);

        //					list += '<div class="row">'; // row 1
        //					list += '<div class="col-xl p-1">'; // col-xl
        /*
        node.setAttribute
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
        */
      }
    }
  }
  ws.onerror = function (evt) {
    console.log("ERROR: " + evt.data);
  }
  window.onfocus = function (event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend out
    // progress bar setTimeout functions
    command("status");
  }
  updateProgress();

  var command = function (cmd) {
    if (!ws) { return false; }
    myJson = { "command": cmd }
    console.log("SEND: " + myJson);
    ws.send(JSON.stringify(myJson));
    return false;
  }

  // add onclick function for all controls
  var controls = ["play", "resume", "pause", "stop", "next", "previous"]
  controls.forEach(activator);
  function activator(value) {
    document.getElementById(value).onclick = function (evt) {
      return command(value);
    };
  }

});